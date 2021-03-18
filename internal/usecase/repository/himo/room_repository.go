package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/go-gorp/gorp"
)

// RoomRepository はインターフェース
type RoomRepository interface {
	Create(db *gorp.DbMap, room model.Room) error
	FetchThemesByChannelName(db *gorp.DbMap, channelName string) ([]model.Theme, error)
	FetchByChannelName(db *gorp.DbMap, channelName string) (model.Room, error)
}
