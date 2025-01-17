package themeLoader

import (
	"SeKai/internal/config"
	"SeKai/internal/model"
)

var backStageThemeMap = make(map[string]themeConfig)
var frontStageThemeMap = make(map[string]themeConfig)

var BackStageTheme model.Theme
var FrontStageTheme model.Theme

func InitThemeLoader() {
	ThemeBasicScan("./themes/backStage", backStageThemeMap)
	ThemeBasicScan("./themes/frontStage", frontStageThemeMap)

	BackStageTheme = SingleThemeScan(
		"./themes/backStage",
		config.ApplicationConfig.SiteConfig.SiteBackStageTheme,
		backStageThemeMap,
	)
	FrontStageTheme = SingleThemeScan(
		"./themes/frontStage",
		config.ApplicationConfig.SiteConfig.SiteFrontStageTheme,
		frontStageThemeMap,
	)
}
