package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/bamboooo-dev/himo-ingame/internal/interface/router"
	"github.com/bamboooo-dev/himo-ingame/pkg/env"
)

// from LDFLAGS
var revision = "undefined"

// from flag
var templateFilePath string

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic '%v' captured\n", err)
		}
	}()

	fmt.Printf("Version is %s\n", revision)
	flag.StringVar(&templateFilePath, "f", "config/application.yml.tpl", "path of config template")
	flag.Parse()

	cfg, err := env.LoadConfigFromTemplate(templateFilePath)
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
