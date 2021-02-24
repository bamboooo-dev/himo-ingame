package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bamboooo-dev/himo-ingame/internal/interface/router"
	"github.com/bamboooo-dev/himo-ingame/pkg/env"
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

	s := router.NewServer()

	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatal("error starting http server::", err)
		return
	}
}
