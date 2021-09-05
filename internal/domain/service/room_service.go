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
	userRepo      repo.UserRepository
	userRoomRepo  repo.UserRoomRepository
}

// NewRoomService は RoomService のコンストラクタ
func NewRoomService(r registry.Registry) *RoomService {
	return &RoomService{
		roomRepo:      r.NewRoomRepository(),
		roomThemeRepo: r.NewRoomThemeRepository(),
		themeRepo:     r.NewThemeRepository(),
		userRepo:      r.NewUserRepository(),
		userRoomRepo:  r.NewUserRoomRepository(),
	}
}

// Create は部屋を作成する
func (r *RoomService) Create(db *gorp.DbMap, max int, channelName string, themeIDs []int, userID int) (model.Room, error) {
	user, err := r.userRepo.FetchByID(db, userID)
	if err != nil {
		return model.Room{}, err
	}

	themes, err := r.themeRepo.FetchByIDs(db, themeIDs)
	if err != nil {
		return model.Room{}, err
	}

	room := model.Room{
		MaxUserNum:  max,
		ChannelName: channelName,
		Themes:      themes,
	}

	room, err = r.roomRepo.Create(db, room)
	if err != nil {
		return model.Room{}, err
	}

	err = r.roomThemeRepo.BulkCreate(db, room, themes)
	if err != nil {
		return model.Room{}, err
	}

	err = r.userRoomRepo.BulkCreate(db, user, room)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}

// Enter は部屋に入る
func (r *RoomService) Enter(db *gorp.DbMap, channelName string, userID int) (model.Room, error) {
	user, err := r.userRepo.FetchByID(db, userID)
	if err != nil {
		return model.Room{}, err
	}

	themes, err := r.roomRepo.FetchThemesByChannelName(db, channelName)
	if err != nil {
		return model.Room{}, err
	}

	room, err := r.roomRepo.FetchByChannelName(db, channelName)
	if err != nil {
		return model.Room{}, err
	}
	room.Themes = themes

	err = r.userRoomRepo.BulkCreate(db, user, room)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}

// Start はゲーム開始
func (r *RoomService) Start(db *gorp.DbMap, channelName string) (model.Room, error) {
	users, err := r.userRoomRepo.FetchUsersByChannelName(db, channelName)
	if err != nil {
		return model.Room{}, err
	}

	room, err := r.roomRepo.FetchByChannelName(db, channelName)
	if err != nil {
		return model.Room{}, err
	}

	room.Members = users

	return room, nil
}

// Update はもう一回ゲーム開始
func (r *RoomService) Update(db *gorp.DbMap, channelName string, themeIDs []int, userID int) (model.Room, error) {
	user, err := r.userRepo.FetchByID(db, userID)
	if err != nil {
		return model.Room{}, err
	}

	room, err := r.roomRepo.FetchByChannelName(db, channelName)
	if err != nil {
		return model.Room{}, err
	}

	// 元々の room に紐づく user を一旦全て削除
	err = r.userRoomRepo.BulkDelete(db, room)
	if err != nil {
		return model.Room{}, err
	}
	oldThemes, err := r.roomRepo.FetchThemesByChannelName(db, channelName)
	if err != nil {
		return model.Room{}, err
	}
	room.Themes = oldThemes

	newThemes, err := r.themeRepo.FetchByIDs(db, themeIDs)
	if err != nil {
		return model.Room{}, err
	}

	// 元々の room に紐づくお題を更新
	err = r.roomThemeRepo.BulkUpdate(db, room, newThemes)
	if err != nil {
		return model.Room{}, err
	}
	room.Themes = newThemes

	// 対象の room に紐づく user に自分を追加
	err = r.userRoomRepo.BulkCreate(db, user, room)
	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}
