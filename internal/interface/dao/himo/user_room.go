package himo

// UserRoom は部屋とお題を結びつける DAO
type UserRoom struct {
	UserID int `db:"user_id"`
	RoomID int `db:"room_id"`
}
