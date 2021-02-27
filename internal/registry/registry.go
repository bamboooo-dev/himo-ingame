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

func (r *registry) Config() *env.Config {
	if r.config == nil {
		return &env.Config{}
	}
	return r.config
}
