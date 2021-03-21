package interactor

import (
	"fmt"

	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/bamboooo-dev/himo-ingame/internal/domain/service"
	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
	"github.com/google/uuid"
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
func (c *CreateRoomInteractor) Call(db *gorp.DbMap, max int, themeIDs []int, userID int) (model.Room, error) {
	// channelName に使うランダム文字列を生成(uuid から部分的に抽出)
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return model.Room{}, err
	}
	uu := u.String()
	channelName := uu[1:8] // 抽出の場所は適当

	room, err := c.roomService.Create(db, max, channelName, themeIDs, userID)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}
