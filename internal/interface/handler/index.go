package handler

import "net/http"

// IndexHandler は / 以下のエンドポイントを管理する構造体です。
type IndexHandler struct {
}

// NewIndexHandler は IndexHandler のポインタを生成する関数です。
func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

// / にアクセスしたら index.html を読み込む
func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
