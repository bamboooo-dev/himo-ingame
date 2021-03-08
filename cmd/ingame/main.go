package main

import (
	"context"
	"fmt"

	"github.com/bamboooo-dev/himo-ingame/internal/interface/handler"
	"github.com/bamboooo-dev/himo-ingame/internal/interface/mysql"
	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	"github.com/bamboooo-dev/himo-ingame/pkg/env"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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
	router.POST("/room", func(c *gin.Context) { roomHandler.Create(c) })
	router.POST("/enter", func(c *gin.Context) { roomHandler.Enter(c) })
	router.Run(":8080")
}
