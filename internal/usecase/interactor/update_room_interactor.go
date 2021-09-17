package interactor

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/bamboooo-dev/himo-ingame/internal/domain/service"
	"github.com/bamboooo-dev/himo-ingame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
)

// UpdateRoomInteractor はもう一度同じメンバーでゲームをする際に部屋を更新するユースケース(お題の変更など)を司る構造体
type UpdateRoomInteractor struct {
	roomThemeRepo himo_repo.RoomThemeRepository
	roomRepo      himo_repo.RoomRepository
	roomService   *service.RoomService
}

// NewUpdateRoomInteractor は UpdateRoomInteractor のコンストラクタ
func NewUpdateRoomInteractor(r registry.Registry) *UpdateRoomInteractor {
	return &UpdateRoomInteractor{
		roomThemeRepo: r.NewRoomThemeRepository(),
		roomRepo:      r.NewRoomRepository(),
		roomService:   service.NewRoomService(r),
	}
}

// Call は部屋を更新する関数
func (c *UpdateRoomInteractor) Call(db *gorp.DbMap, channelName string, themeIDs []int, userID int) (model.Room, error) {
	updatedRoom, err := c.roomService.Update(db, channelName, themeIDs, userID)
	if err != nil {
		return model.Room{}, err
	}

	return updatedRoom, nil
}
