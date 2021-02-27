package handler

import (
	"math/rand"

	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	"github.com/bamboooo-dev/himo-ingame/internal/usecase/interactor"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// RoomHandler は / 以下のエンドポイントを管理する構造体です。
type RoomHandler struct {
	logger  *zap.SugaredLogger
	creator *interactor.CreateRoomInteractor
	db      *gorp.DbMap
}

// NewRoomHandler は IndexHandler のポインタを生成する関数です。
func NewRoomHandler(l *zap.SugaredLogger, r registry.Registry, db *gorp.DbMap) *RoomHandler {
	return &RoomHandler{
		logger:  l,
		creator: interactor.NewCreateRoomInteractor(r),
		db:      db,
	}
}

// Create creates room
func (r RoomHandler) Create(c Context) {
	max := c.PostForm("max_num").(int64)

	// channelName に使うランダム文字列を生成
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 15)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	randomString := string(b)

	room, err := r.creator.Call(r.db, max, randomString)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(200, room.ChannelName)
}
