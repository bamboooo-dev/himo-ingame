package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/go-gorp/gorp"
)

// UserRoomRepository はインターフェース
type UserRoomRepository interface {
	BulkCreate(db *gorp.DbMap, user model.User, room model.Room) error
}
