package handler

import (
	"math/rand"

	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	"github.com/bamboooo-dev/himo-ingame/internal/usecase/interactor"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// RoomHandler は / 以下のエンドポイントを管理する構造体です。
type RoomHandler struct {
	logger  *zap.SugaredLogger
	creator *interactor.CreateRoomInteractor
	enteror *interactor.EnterRoomInteractor
	db      *gorp.DbMap
}

type CreateRoomRequest struct {
	FieldMaxNum   int   `json:"max_num"`
	FieldThemeIds []int `json:"theme_ids"`
}

type EnterRoomRequest struct {
	FieldChannelName string `json:"channel_name"`
}

// NewRoomHandler は IndexHandler のポインタを生成する関数です。
func NewRoomHandler(l *zap.SugaredLogger, r registry.Registry, db *gorp.DbMap) *RoomHandler {
	return &RoomHandler{
		logger:  l,
		creator: interactor.NewCreateRoomInteractor(r),
		enteror: interactor.NewEnterRoomInteractor(r),
		db:      db,
	}
}

// Create creates room
func (r *RoomHandler) Create(c *gin.Context) {
	// request の中身を取得
	var json CreateRoomRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	max := json.FieldMaxNum
	themeIDs := json.FieldThemeIds

	// channelName に使うランダム文字列を生成
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 15)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	randomString := string(b)

	room, err := r.creator.Call(r.db, max, randomString, themeIDs)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(200, gin.H{
		"message":      "Room successfully created",
		"channel_name": room.ChannelName,
		"max_num":      room.MaxUserNum,
		"theme_ids":    themeIDs,
	})
}

// Enter find room and return themes
func (r *RoomHandler) Enter(c *gin.Context) {
	// request の中身を取得
	var json EnterRoomRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	channelName := json.FieldChannelName

	themes, err := r.enteror.Call(r.db, channelName)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}
	c.JSON(200, gin.H{
		"message": "Successfully entered room",
		"themes":  themes,
	})
}
