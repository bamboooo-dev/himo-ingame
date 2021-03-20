package himo

// User はユーザーの DAO
type User struct {
	ID       int    `db:"id, primarykey, autoincrement"`
	Nickname string `db:"nickname"`
}
