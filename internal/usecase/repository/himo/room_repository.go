package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/go-gorp/gorp"
)

// RoomRepository はインターフェース
type RoomRepository interface {
	Create(db *gorp.DbMap, max int, channelName string) (model.Room, error)
	FetchThemesByChannelName(db *gorp.DbMap, channelName string) ([]model.Theme, error)
}
