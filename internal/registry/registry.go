package registry

import (
	himo_mysql "github.com/bamboooo-dev/himo-ingame/internal/interface/mysql/himo"
	himo_repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/bamboooo-dev/himo-ingame/pkg/env"
	"go.uber.org/zap"
)

// Registry は DI コンテナ
type Registry interface {
	Config() *env.Config
	NewRoomRepository() himo_repo.RoomRepository
	NewRoomThemeRepository() himo_repo.RoomThemeRepository
	NewThemeRepository() himo_repo.ThemeRepository
}

type registry struct {
	config *env.Config
	l      *zap.SugaredLogger
}

// NewRegistry is Registry constructor.
func NewRegistry(cfg *env.Config, l *zap.SugaredLogger) Registry {
	return &registry{cfg, l}
}

func (r *registry) NewRoomRepository() himo_repo.RoomRepository {
	return himo_mysql.NewRoomRepositoryMysql(r.l)
}

func (r *registry) NewRoomThemeRepository() himo_repo.RoomThemeRepository {
	return himo_mysql.NewRoomThemeRepositoryMysql(r.l)
}

func (r *registry) NewThemeRepository() himo_repo.ThemeRepository {
	return himo_mysql.NewThemeRepositoryMysql(r.l)
}

func (r *registry) Config() *env.Config {
	if r.config == nil {
		return &env.Config{}
	}
	return r.config
}
