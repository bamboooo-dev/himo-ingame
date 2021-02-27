package service

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// RoomService は部屋に関するドメインサービスの構造体
type RoomService struct {
	RoomRepo repo.RoomRepository
}

// NewRoomService は RoomService のコンストラクタ
func NewRoomService(r registry.Registry) *RoomService {
	return &RoomService{
		RoomRepo: r.NewRoomRepository(),
	}
}

// Create は部屋を作成する
func (r *RoomService) Create(db *gorp.DbMap, max int, channelName string) (model.Room, error) {
	Room, err := r.RoomRepo.Create(db, max, channelName)
	if err != nil {
		return model.Room{}, err
	}
	return Room, nil
}
