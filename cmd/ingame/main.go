package main

import (
	"log"
	"net/http"

	"github.com/bamboooo-dev/himo-ingame/internal/interface/router"
)

func main() {

	s := router.NewServer()

	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatal("error starting http server::", err)
		return
	}
}
