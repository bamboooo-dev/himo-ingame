package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

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
	starter *interactor.StartRoomInteractor
	db      *gorp.DbMap
}

type CreateRoomRequest struct {
	FieldMaxNum   int   `json:"max_num"`
	FieldThemeIds []int `json:"theme_ids"`
}

type EnterRoomRequest struct {
	FieldChannelName string `json:"channel_name"`
}

type StartRoomRequest struct {
	FieldChannelName string `json:"channel_name"`
}
type StartRoomMessage struct {
	FieldType    string   `json:"type"`
	FieldNumbers []int    `json:"nubmbers"`
	FieldNames   []string `json:"names"`
	FieldMessage string   `json:"message"`
}

// NewRoomHandler は IndexHandler のポインタを生成する関数です。
func NewRoomHandler(l *zap.SugaredLogger, r registry.Registry, db *gorp.DbMap) *RoomHandler {
	return &RoomHandler{
		logger:  l,
		creator: interactor.NewCreateRoomInteractor(r),
		enteror: interactor.NewEnterRoomInteractor(r),
		starter: interactor.NewStartRoomInteractor(r),
		db:      db,
	}
}

// Create creates room
func (r *RoomHandler) Create(c *gin.Context) {
	// request の中身を取得
	var reqJson CreateRoomRequest

	// log for debug
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	b := string(buf[0:n])

	// body が Read で空になったので再度入れ込む処理
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(b)))

	fmt.Printf("post /room request header:\n %v\n", c.Request.Header)
	fmt.Printf("post /room request body:\n %v\n", b)

	if err := c.ShouldBindJSON(&reqJson); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	max := reqJson.FieldMaxNum
	themeIDs := reqJson.FieldThemeIds

	userID, _ := c.Get("AuthorizedUser")

	room, err := r.creator.Call(r.db, max, themeIDs, userID.(int))
	if err == sql.ErrNoRows {
		c.JSON(404, "Not Found")
		return
	}
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}

	// log for debug
	fmt.Printf("post /room response:\n %v\n", gin.H{
		"message":      "Room successfully created",
		"channel_name": room.ChannelName,
		"max_num":      room.MaxUserNum,
		"themes":       room.Themes,
	})

	c.JSON(200, gin.H{
		"message":      "Room successfully created",
		"channel_name": room.ChannelName,
		"max_num":      room.MaxUserNum,
		"themes":       room.Themes,
	})
}

// Enter find room and return themes
func (r *RoomHandler) Enter(c *gin.Context) {
	// request の中身を取得
	var reqJson EnterRoomRequest

	// log for debug
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	b := string(buf[0:n])

	// body が Read で空になったので再度入れ込む処理
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(b)))

	fmt.Printf("post /enter request header:\n %v\n", c.Request.Header)
	fmt.Printf("post /enter request body:\n %v\n", b)

	if err := c.ShouldBindJSON(&reqJson); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	channelName := reqJson.FieldChannelName

	userID, _ := c.Get("AuthorizedUser")

	room, err := r.enteror.Call(r.db, channelName, userID.(int))
	if err == sql.ErrNoRows {
		c.JSON(404, "Not Found")
		return
	}
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}

	// log for debug
	fmt.Printf("post /room response: %v", gin.H{
		"message": "Successfully entered room",
		"themes":  room.Themes,
		"max_num": room.MaxUserNum,
	})

	c.JSON(200, gin.H{
		"message": "Successfully entered room",
		"themes":  room.Themes,
		"max_num": room.MaxUserNum,
	})
}

func (r *RoomHandler) Start(c *gin.Context) {
	// request の中身を取得
	var reqJson StartRoomRequest

	// log for debug
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	b := string(buf[0:n])

	// body が Read で空になったので再度入れ込む処理
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(b)))

	fmt.Printf("post /start request header:\n %v\n", c.Request.Header)
	fmt.Printf("post /start request body:\n %v\n", b)

	if err := c.ShouldBindJSON(&reqJson); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	channelName := reqJson.FieldChannelName

	room, err := r.starter.Call(r.db, channelName)
	if err == sql.ErrNoRows {
		c.JSON(404, "Not Found")
		return
	}
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}

	// publish する Message を構成
	users := room.Members
	var userNames []string
	for _, user := range users {
		userNames = append(userNames, user.Nickname)
	}

	numbers := pickup(1, 100, len(userNames))

	pubMessage := StartRoomMessage{
		FieldType:    "answer",
		FieldNumbers: numbers,
		FieldNames:   userNames,
		FieldMessage: "Successfully entered room",
	}

	// 構成したメッセージを json で POST して publish
	pubMessageJson, _ := json.Marshal(pubMessage)
	fmt.Printf("[+] %s\n", string(pubMessageJson))

	url := "http://nchan/pub?channel_id=" + channelName

	response, err := http.Post(url, "application/json", bytes.NewBuffer(pubMessageJson))
	defer response.Body.Close()
	if err != nil {
		fmt.Println("[!] " + err.Error())
	} else {
		fmt.Println("[*] " + response.Status)
	}

	c.JSON(200, gin.H{
		"type":    "answer",
		"numbers": numbers,
		"names":   userNames,
		"message": "Successfully entered room",
	})
}

// 1~100 からランダムに数字を取ってくるための関数たち
func allKeys(m map[int]bool) []int {
	i := 0
	result := make([]int, len(m))
	for key, _ := range m {
		result[i] = key
		i++
	}
	return result
}

func pickup(min int, max int, num int) []int {
	numRange := max - min

	selected := make(map[int]bool)
	rand.Seed(time.Now().UnixNano())
	for counter := 0; counter < num; {
		n := rand.Intn(numRange) + min
		if selected[n] == false {
			selected[n] = true
			counter++
		}
	}
	keys := allKeys(selected)
	return keys
}
