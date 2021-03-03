package interactor

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/bamboooo-dev/himo-ingame/internal/domain/service"
	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// CreateRoomInteractor は部屋を作るユースケースを司る構造体
type CreateRoomInteractor struct {
	roomThemeRepo himo_repo.RoomThemeRepository
	roomRepo      himo_repo.RoomRepository
	roomService   *service.RoomService
}

// NewCreateRoomInteractor は CreateRoomInteractor のコンストラクタ
func NewCreateRoomInteractor(r registry.Registry) *CreateRoomInteractor {
	return &CreateRoomInteractor{
		roomThemeRepo: r.NewRoomThemeRepository(),
		roomRepo:      r.NewRoomRepository(),
		roomService:   service.NewRoomService(r),
	}
}

// Call は部屋を作る関数
func (c *CreateRoomInteractor) Call(db *gorp.DbMap, max int, channelName string, themeIDs []int) (model.Room, error) {
	room, err := c.roomService.Create(db, max, channelName)
	if err != nil {
		return model.Room{}, err
	}

	_, err = c.roomThemeRepo.Create(db, room.ID, themeIDs)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}
