package interactor

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/bamboooo-dev/himo-ingame/internal/domain/service"
	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// StartRoomInteractor は部屋を作るユースケースを司る構造体
type StartRoomInteractor struct {
	roomRepo     himo_repo.RoomRepository
	userRepo     himo_repo.UserRepository
	userRoomRepo himo_repo.UserRoomRepository
	roomService  *service.RoomService
}

// NewStartRoomInteractor は StartRoomInteractor のコンストラクタ
func NewStartRoomInteractor(r registry.Registry) *StartRoomInteractor {
	return &StartRoomInteractor{
		roomRepo:     r.NewRoomRepository(),
		userRepo:     r.NewUserRepository(),
		userRoomRepo: r.NewUserRoomRepository(),
		roomService:  service.NewRoomService(r),
	}
}

// Call は部屋を作る関数
func (s *StartRoomInteractor) Call(db *gorp.DbMap, channelName string) (model.Room, error) {
	room, err := s.roomService.Start(db, channelName)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}
