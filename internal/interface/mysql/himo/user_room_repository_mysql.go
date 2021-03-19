package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	dao "github.com/bamboooo-dev/himo-ingame/internal/interface/dao/himo"
	repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// UserRoomRepositoryMysql は UserRoomRepository の MySQL 実装
type UserRoomRepositoryMysql struct {
	logger *zap.SugaredLogger
}

// NewUserRoomRepositoryMysql は UserRoomRepositoryMysql のコンストラクタ
func NewUserRoomRepositoryMysql(l *zap.SugaredLogger) repo.UserRoomRepository {
	return UserRoomRepositoryMysql{logger: l}
}

// Create new UserRoom
func (u UserRoomRepositoryMysql) BulkCreate(db *gorp.DbMap, user model.User, room model.Room) error {
	userRoomDAO := &dao.UserRoom{
		UserID: user.ID,
		RoomID: room.ID,
	}
	err := db.Insert(userRoomDAO)
	if err != nil {
		return err
	}

	return nil
}
