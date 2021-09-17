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
func (r RoomThemeRepositoryMysql) BulkCreate(db *gorp.DbMap, room model.Room, themes []model.Theme) error {
	for _, theme := range themes {
		roomThemeDAO := &dao.RoomTheme{
			RoomID:  room.ID,
			ThemeID: theme.ID,
		}

		// TODO: N+1 なってるからバルクインサートするようにしたい
		err := db.Insert(roomThemeDAO)
		if err != nil {
			return err
		}
	}

	return nil
}

// Update RoomTheme
func (r RoomThemeRepositoryMysql) BulkUpdate(db *gorp.DbMap, room model.Room, themes []model.Theme) error {
	// バルクアップデートしている
	_, err := db.Exec(
		"update room_themes set theme_id = elt(field(theme_id, ?, ?, ?), ?, ?, ?) where room_id = ?",
		room.Themes[0].ID, room.Themes[1].ID, room.Themes[2].ID,
		themes[0].ID, themes[1].ID, themes[2].ID,
		room.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
