package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bamboooo-dev/himo-ingame/internal/interface/mysql"
	"github.com/bamboooo-dev/himo-ingame/internal/interface/router"
	"github.com/bamboooo-dev/himo-ingame/pkg/env"
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

	s := router.NewServer()

	errs := http.ListenAndServe(":8080", s)
	if errs != nil {
		log.Fatal("error starting http server::", errs)
		return
	}
}
