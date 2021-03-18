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
	themeRepo     repo.ThemeRepository
}

// NewRoomService は RoomService のコンストラクタ
func NewRoomService(r registry.Registry) *RoomService {
	return &RoomService{
		roomRepo:      r.NewRoomRepository(),
		roomThemeRepo: r.NewRoomThemeRepository(),
		themeRepo:     r.NewThemeRepository(),
	}
}

// Create は部屋を作成する
func (r *RoomService) Create(db *gorp.DbMap, max int, channelName string, themeIDs []int) (model.Room, error) {
	themes, err := r.themeRepo.FetchByIDs(db, themeIDs)
	if err != nil {
		return model.Room{}, err
	}

	room := model.Room{
		MaxUserNum:  max,
		ChannelName: channelName,
		Themes:      themes,
	}

	err = r.roomRepo.Create(db, room)
	if err != nil {
		return model.Room{}, err
	}

	err = r.roomThemeRepo.BulkCreate(db, room, themes)
	if err != nil {
		return model.Room{}, err
	}
	return room, nil
}

// Enter は部屋に入る
func (r *RoomService) Enter(db *gorp.DbMap, channelName string) (model.Room, error) {
	themes, err := r.roomRepo.FetchThemesByChannelName(db, channelName)
	if err != nil {
		return model.Room{}, err
	}

	room, err := r.roomRepo.FetchByChannelName(db, channelName)
	if err != nil {
		return model.Room{}, err
	}

	room = model.Room{
		MaxUserNum:  room.MaxUserNum,
		ChannelName: room.ChannelName,
	}

	room.Themes = themes

	return room, nil
}
