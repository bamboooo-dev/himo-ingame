package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	dao "github.com/bamboooo-dev/himo-ingame/internal/interface/dao/himo"
	repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// RoomRepositoryMysql は RoomRepository の MySQL 実装
type RoomRepositoryMysql struct {
	logger *zap.SugaredLogger
}

// NewRoomRepositoryMysql は RoomRepositoryMysql のコンストラクタ
func NewRoomRepositoryMysql(l *zap.SugaredLogger) repo.RoomRepository {
	return RoomRepositoryMysql{logger: l}
}

// Create new Room
func (r RoomRepositoryMysql) Create(db *gorp.DbMap, max int, channelName string) (model.Room, error) {
	roomDAO := &dao.Room{
		MaxUserNum:  max,
		ChannelName: channelName,
	}

	err := db.Insert(roomDAO)
	if err != nil {
		return model.Room{}, err
	}

	room := model.Room{
		ID:          roomDAO.ID,
		MaxUserNum:  roomDAO.MaxUserNum,
		ChannelName: roomDAO.ChannelName,
	}
	return room, nil
}
