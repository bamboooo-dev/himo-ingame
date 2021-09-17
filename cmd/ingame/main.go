package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/bamboooo-dev/himo-ingame/internal/interface/handler"
	"github.com/bamboooo-dev/himo-ingame/internal/interface/mysql"
	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	"github.com/bamboooo-dev/himo-ingame/pkg/env"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// jwt 用
type User struct {
	ID          string
	Nickname    string
	AccessToken string
}

// from LDFLAGS
var revision = "undefined"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic '%v' captured\n", err)
		}
	}()

	fmt.Printf("Version is %s\n", revision)

	cfg, err := env.LoadConfigFromTemplate()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	himoDB, err := mysql.NewDB(cfg.HimoMySQL)
	if err != nil {
		sugar.Error(ctx, err)
		return
	}
	defer func() {
		if err := himoDB.Db.Close(); err != nil {
			sugar.Error(ctx, err)
			return
		}
	}()

	registry := registry.NewRegistry(cfg, sugar)

	router := gin.Default()
	roomHandler := handler.NewRoomHandler(sugar, registry, himoDB)

	var identityKey = "user_id"

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(os.Getenv("JWT_SECRET")),
		IdentityKey: identityKey,
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				ID: claims[identityKey].(string),
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok {
				userID, _ := strconv.Atoi(v.ID)
				c.Set("AuthorizedUser", userID) // gin の Context に得られた userID を入れておく
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			// log for debug
			buf := make([]byte, 2048)
			n, _ := c.Request.Body.Read(buf)
			b := string(buf[0:n])

			// body が Read で空になったので再度入れ込む処理
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(b)))

			fmt.Printf("auth middleware request header:\n %v\n", c.Request.Header)
			fmt.Printf("auth middleware request body:\n %v\n", b)

			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	router.Use(authMiddleware.MiddlewareFunc())
	{
		router.POST("/room", func(c *gin.Context) { roomHandler.Create(c) })
		router.POST("/enter", func(c *gin.Context) { roomHandler.Enter(c) })
		router.POST("/start", func(c *gin.Context) { roomHandler.Start(c) })
		router.POST("/update", func(c *gin.Context) { roomHandler.Update(c) })
	}
	router.Run(":8080")
}
