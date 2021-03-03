package himo

// RoomTheme は部屋とお題を結びつける DAO
type RoomTheme struct {
	RoomID  int `db:"room_id"`
	ThemeID int `db:"theme_id"`
}
