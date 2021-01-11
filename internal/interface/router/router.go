package router

import (
	"net/http"

	"github.com/bamboooo-dev/himo-ingame/internal/interface/handler"
)

// NewServer は http パッケージのマルチプレクササーバーを返す
func NewServer() *http.ServeMux {
	mux := http.NewServeMux()

	indexHandler := handler.NewIndexHandler()
	wsHandler := handler.NewWebSocketHandler()

	mux.Handle("/", indexHandler)
	mux.Handle("/ws", wsHandler)

	return mux
}
