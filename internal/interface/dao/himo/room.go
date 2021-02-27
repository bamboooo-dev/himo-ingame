package himo

// Room は部屋の DAO
type Room struct {
	ID          int64  `db:"id, primarykey, autoincrement"`
	MaxUserNum  int64  `db:"max_user_num"`
	ChannelName string `db:"channel_name"`
}
