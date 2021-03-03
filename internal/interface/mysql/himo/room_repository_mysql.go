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

// FetchThemesByChannelName fetch themes by channelName
func (r RoomRepositoryMysql) FetchThemesByChannelName(db *gorp.DbMap, channelName string) ([]model.Theme, error) {
	var daoThemes []dao.Theme

	_, err := db.Select(&daoThemes, "SELECT t.id, t.sentence FROM room_themes AS rt INNER JOIN rooms AS r ON rt.room_id = r.id INNER JOIN themes AS t ON rt.theme_id = t.id WHERE r.channel_name = ?", channelName)
	if err != nil {
		return []model.Theme{}, err
	}

	themes := []model.Theme{}
	for _, daoTheme := range daoThemes {
		theme := model.Theme{
			ID:       daoTheme.ID,
			Sentence: daoTheme.Sentence,
		}
		themes = append(themes, theme)
	}
	return themes, nil
}
