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

// FetchUsersByChannelName fetch users by channelName
func (u UserRoomRepositoryMysql) FetchUsersByChannelName(db *gorp.DbMap, channelName string) ([]model.User, error) {
	var daoUsers []dao.User

	_, err := db.Select(&daoUsers, "SELECT u.id, u.nickname FROM user_rooms AS ur INNER JOIN users AS u ON ur.user_id = u.id INNER JOIN rooms AS r ON ur.room_id = r.id WHERE r.channel_name = ?", channelName)
	if err != nil {
		return []model.User{}, err
	}

	users := []model.User{}
	for _, daoUser := range daoUsers {
		user := model.User{
			ID:       daoUser.ID,
			Nickname: daoUser.Nickname,
		}
		users = append(users, user)
	}
	return users, nil
}

// Delete UserRooms by roomID
func (u UserRoomRepositoryMysql) BulkDelete(db *gorp.DbMap, room model.Room) error {
	_, err := db.Exec("delete from user_rooms where room_id = ?", room.ID)
	if err != nil {
		return err
	}
	return nil
}
