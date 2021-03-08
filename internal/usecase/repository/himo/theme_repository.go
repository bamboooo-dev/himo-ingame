package himo

import (
	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	"github.com/go-gorp/gorp"
)

// ThemeRepository はインターフェース
type ThemeRepository interface {
	FetchByIDs(db *gorp.DbMap, themeIDs []int) ([]model.Theme, error)
}
