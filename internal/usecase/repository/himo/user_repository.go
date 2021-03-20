package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/go-gorp/gorp"
)

// UserRepository はインターフェース
type UserRepository interface {
	FetchByID(db *gorp.DbMap, userID int) (model.User, error)
}
