package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/go-gorp/gorp"
)

// RoomThemeRepository はインターフェース
type RoomThemeRepository interface {
	Create(db *gorp.DbMap, room model.Room) ([]model.RoomTheme, error)
}
