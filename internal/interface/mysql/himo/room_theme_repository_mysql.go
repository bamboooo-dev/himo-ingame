package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	dao "github.com/bamboooo-dev/himo-ingame/internal/interface/dao/himo"
	repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// RoomThemeRepositoryMysql は RoomThemeRepository の MySQL 実装
type RoomThemeRepositoryMysql struct {
	logger *zap.SugaredLogger
}

// NewRoomThemeRepositoryMysql は RoomThemeRepositoryMysql のコンストラクタ
func NewRoomThemeRepositoryMysql(l *zap.SugaredLogger) repo.RoomThemeRepository {
	return RoomThemeRepositoryMysql{logger: l}
}

// Create new RoomTheme
func (r RoomThemeRepositoryMysql) Create(db *gorp.DbMap, room model.Room) ([]model.RoomTheme, error) {

	var roomThemes []model.RoomTheme

	themes := room.Themes

	for _, theme := range themes {
		roomThemeDAO := &dao.RoomTheme{
			RoomID:  room.ID,
			ThemeID: theme.ID,
		}

		// TODO: N+1 なってるからバルクインサートするようにしたい
		err := db.Insert(roomThemeDAO)
		if err != nil {
			return []model.RoomTheme{}, err
		}

		roomTheme := model.RoomTheme{
			RoomID:  room.ID,
			ThemeID: theme.ID,
		}
		roomThemes = append(roomThemes, roomTheme)
	}

	return roomThemes, nil
}
