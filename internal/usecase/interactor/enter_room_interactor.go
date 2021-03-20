package interactor

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/bamboooo-dev/himo-ingame/internal/domain/service"
	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// EnterRoomInteractor は部屋に入るユースケースを司る構造体
type EnterRoomInteractor struct {
	roomRepo    himo_repo.RoomRepository
	roomService *service.RoomService
}

// NewEnterRoomInteractor は EnterRoomInteractor のコンストラクタ
func NewEnterRoomInteractor(r registry.Registry) *EnterRoomInteractor {
	return &EnterRoomInteractor{
		roomRepo:    r.NewRoomRepository(),
		roomService: service.NewRoomService(r),
	}
}

// Call は部屋に入る関数
func (e *EnterRoomInteractor) Call(db *gorp.DbMap, channelName string, userID int) (model.Room, error) {
	room, err := e.roomService.Enter(db, channelName, userID)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}
