package service

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// RoomService は部屋に関するドメインサービスの構造体
type RoomService struct {
	roomRepo      repo.RoomRepository
	roomThemeRepo repo.RoomThemeRepository
}

// NewRoomService は RoomService のコンストラクタ
func NewRoomService(r registry.Registry) *RoomService {
	return &RoomService{
		roomRepo:      r.NewRoomRepository(),
		roomThemeRepo: r.NewRoomThemeRepository(),
	}
}

// Create は部屋を作成する
func (r *RoomService) Create(db *gorp.DbMap, max int, channelName string, themeIDs []int) (model.Room, error) {
	room, err := r.roomRepo.Create(db, max, channelName, themeIDs)
	if err != nil {
		return model.Room{}, err
	}

	_, err = r.roomThemeRepo.Create(db, room)
	if err != nil {
		return model.Room{}, err
	}
	return room, nil
}
