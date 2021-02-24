package himo

// User はユーザーの DAO
type User struct {
	ID       int64  `db:"id, primarykey, autoincrement"`
	Nickname string `db:"nickname"`
}
