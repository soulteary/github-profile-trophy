package api

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/soulteary/github-profile-trophy/internal/cards"
	"github.com/soulteary/github-profile-trophy/internal/common"
	"github.com/soulteary/github-profile-trophy/internal/fetchers"
	"github.com/soulteary/github-profile-trophy/internal/themes"
)

// writeOutput writes SVG content to file or stdout
func writeOutput(content, outputPath string) error {
	if outputPath == "" {
		_, err := os.Stdout.WriteString(content)
		return err
	}
	return os.WriteFile(outputPath, []byte(content), 0644)
}

// getParam gets a parameter from map with default value
func getParam(params map[string]string, key, defaultValue string) string {
	if val, ok := params[key]; ok && val != "" {
		return val
	}
	return defaultValue
}

// GenerateTrophyCard generates trophy card from parameters
func GenerateTrophyCard(params map[string]string, outputPath string) error {
	username := getParam(params, "username", "")
	if username == "" {
		return fmt.Errorf("username is a required parameter")
	}

	// Get theme
	themeParam := getParam(params, "theme", "default")
	theme := themes.GetTheme(themeParam)

	// Get other parameters
	rowStr := getParam(params, "row", "")
	row := common.DefaultMaxRow
	if rowStr != "" {
		queryParams := common.NewParams("row=" + rowStr)
		row = queryParams.GetNumberValue("row", common.DefaultMaxRow)
	}

	columnStr := getParam(params, "column", "")
	column := common.DefaultMaxColumn
	if columnStr != "" {
		queryParams := common.NewParams("column=" + columnStr)
		column = queryParams.GetNumberValue("column", common.DefaultMaxColumn)
	}

	marginWStr := getParam(params, "margin-w", "")
	marginWidth := common.DefaultMarginW
	if marginWStr != "" {
		queryParams := common.NewParams("margin-w=" + marginWStr)
		marginWidth = queryParams.GetNumberValue("margin-w", common.DefaultMarginW)
	}

	marginHStr := getParam(params, "margin-h", "")
	marginHeight := common.DefaultMarginH
	if marginHStr != "" {
		queryParams := common.NewParams("margin-h=" + marginHStr)
		marginHeight = queryParams.GetNumberValue("margin-h", common.DefaultMarginH)
	}

	noBackground := common.DefaultNoBackground
	if noBgStr := getParam(params, "no-bg", ""); noBgStr != "" {
		if parsed := common.ParseBoolean(noBgStr); parsed != nil {
			noBackground = *parsed
		}
	}

	noFrame := common.DefaultNoFrame
	if noFrameStr := getParam(params, "no-frame", ""); noFrameStr != "" {
		if parsed := common.ParseBoolean(noFrameStr); parsed != nil {
			noFrame = *parsed
		}
	}

	// Parse title and rank parameters
	titleStr := getParam(params, "title", "")
	var titles []string
	if titleStr != "" {
		// Use common.NewParams to parse comma-separated values
		queryParams := common.NewParams("title=" + titleStr)
		titles = queryParams.GetAll("title")
	}

	rankStr := getParam(params, "rank", "")
	var ranks []string
	if rankStr != "" {
		// Use common.NewParams to parse comma-separated values
		queryParams := common.NewParams("rank=" + rankStr)
		ranks = queryParams.GetAll("rank")
	}

	// Try to get from cache
	var userInfo *fetchers.UserInfo
	if cachedData, found := common.GetUserInfoCache(username); found {
		// Unmarshal cached data
		var cachedInfo fetchers.UserInfo
		if err := json.Unmarshal(cachedData, &cachedInfo); err == nil {
			userInfo = &cachedInfo
		}
	}

	if userInfo == nil {
		// Fetch from GitHub API
		client := fetchers.NewGitHubClient()
		info, err := client.RequestUserInfo(username)
		if err != nil {
			if customErr, ok := err.(*common.CustomError); ok {
				if customErr.Type == common.ErrorTypeNotFound {
					return fmt.Errorf("User '%s' not found", username)
				}
			}
			return fmt.Errorf("Failed to fetch user data: %v", err)
		}
		userInfo = &info
		common.SetUserInfoCache(username, userInfo)
	}

	// Create and render card
	card := cards.NewCard(
		titles,
		ranks,
		column,
		row,
		common.DefaultPanelSize,
		marginWidth,
		marginHeight,
		noBackground,
		noFrame,
	)

	svg := card.Render(*userInfo, theme)
	if svg == "" {
		return fmt.Errorf("Failed to generate trophy card for user %s", username)
	}

	return writeOutput(svg, outputPath)
}
