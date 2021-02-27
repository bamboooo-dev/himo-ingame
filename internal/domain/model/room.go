package model

// Room は部屋
type Room struct {
	ID          int64
	MaxUserNum  int64
	ChannelName string
	// ThemeIds     []Theme
}
