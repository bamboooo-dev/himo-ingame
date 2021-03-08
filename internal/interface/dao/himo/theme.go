package himo

// Theme はお題の DAO
type Theme struct {
	ID       int    `db:"id, primarykey, autoincrement"`
	Sentence string `db:"sentence"`
	UserID   int    `db:"user_id"`
}
