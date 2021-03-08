package himo

import (
	"fmt"
	"strings"

	"github.com/bamboooo-dev/himo-ingame/internal/domain/model"
	dao "github.com/bamboooo-dev/himo-ingame/internal/interface/dao/himo"
	repo "github.com/bamboooo-dev/himo-ingame/internal/usecase/repository/himo"
	"github.com/go-gorp/gorp"
	"go.uber.org/zap"
)

// ThemeRepositoryMysql は ThemeRepository の MySQL 実装
type ThemeRepositoryMysql struct {
	logger *zap.SugaredLogger
}

// NewThemeRepositoryMysql は ThemeRepositoryMysql のコンストラクタ
func NewThemeRepositoryMysql(l *zap.SugaredLogger) repo.ThemeRepository {
	return ThemeRepositoryMysql{logger: l}
}

// FetchByIDs fetch themes by ids
func (r ThemeRepositoryMysql) FetchByIDs(db *gorp.DbMap, themeIDs []int) ([]model.Theme, error) {
	var daoThemes []dao.Theme
	args := make([]interface{}, len(themeIDs))
	quarks := make([]string, len(themeIDs))
	for i, themeID := range themeIDs {
		args[i] = themeID
		quarks[i] = "?"
	}

	_, err := db.Select(&daoThemes, fmt.Sprintf("SELECT * FROM themes WHERE id IN (%s)", strings.Join(quarks, ", ")), args...)
	if err != nil {
		return []model.Theme{}, err
	}

	themes := []model.Theme{}
	for _, daoTheme := range daoThemes {
		theme := model.Theme{
			ID:       daoTheme.ID,
			Sentence: daoTheme.Sentence,
		}
		themes = append(themes, theme)
	}

	return themes, nil
}
