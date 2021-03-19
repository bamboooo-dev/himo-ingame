package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	dao "github.com/bamboooo-dev/himo-ingame/internal/interface/dao/himo"
	repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// UserRepositoryMysql は UserRepository の MySQL 実装
type UserRepositoryMysql struct {
	logger *zap.SugaredLogger
}

// NewUserRepositoryMysql は UserRepositoryMysql のコンストラクタ
func NewUserRepositoryMysql(l *zap.SugaredLogger) repo.UserRepository {
	return UserRepositoryMysql{logger: l}
}

// Create new User
func (u UserRepositoryMysql) FetchByID(db *gorp.DbMap, userID int) (model.User, error) {

	var daoUser dao.User

	err := db.SelectOne(&daoUser, "SELECT * FROM users WHERE id = ?", userID)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		ID:       daoUser.ID,
		Nickname: daoUser.Nickname,
	}

	return user, nil
}
