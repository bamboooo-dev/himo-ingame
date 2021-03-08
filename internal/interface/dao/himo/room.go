package himo

// Room は部屋の DAO
type Room struct {
	ID          int    `db:"id, primarykey, autoincrement"`
	MaxUserNum  int    `db:"max_user_num"`
	ChannelName string `db:"channel_name"`
}
