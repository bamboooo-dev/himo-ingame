package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/go-gorp/gorp"
)

// RoomThemeRepository はインターフェース
type RoomThemeRepository interface {
	BulkCreate(db *gorp.DbMap, room model.Room, themes []model.Theme) error
}
