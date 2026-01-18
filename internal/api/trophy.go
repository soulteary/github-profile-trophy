package api

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/github-profile-trophy/internal/cards"
	"github.com/soulteary/github-profile-trophy/internal/common"
	"github.com/soulteary/github-profile-trophy/internal/fetchers"
	"github.com/soulteary/github-profile-trophy/internal/themes"
)

// TrophyHandler handles trophy card requests
func TrophyHandler(c *gin.Context) {
	// Recover from panics
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic in TrophyHandler: %v", r)
			renderErrorPage(c, http.StatusInternalServerError, fmt.Sprintf("Internal server error: %v", r))
		}
	}()

	// Parse query parameters
	query := c.Request.URL.RawQuery
	params := common.NewParams(query)

	username := params.GetStringValue("username", "")
	if username == "" {
		renderErrorPage(c, http.StatusBadRequest, "username is a required query parameter")
		return
	}

	// Get theme
	themeParam := params.GetStringValue("theme", "default")
	theme := themes.GetTheme(themeParam)

	// Get other parameters
	row := params.GetNumberValue("row", common.DefaultMaxRow)
	column := params.GetNumberValue("column", common.DefaultMaxColumn)
	marginWidth := params.GetNumberValue("margin-w", common.DefaultMarginW)
	marginHeight := params.GetNumberValue("margin-h", common.DefaultMarginH)
	noBackground := params.GetBooleanValue("no-bg", common.DefaultNoBackground)
	noFrame := params.GetBooleanValue("no-frame", common.DefaultNoFrame)

	titles := params.GetAll("title")
	ranks := params.GetAll("rank")

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
			log.Printf("Error fetching user info for %s: %v", username, err)
			if customErr, ok := err.(*common.CustomError); ok {
				if customErr.Type == common.ErrorTypeNotFound {
					renderErrorPage(c, http.StatusNotFound, fmt.Sprintf("User '%s' not found", html.EscapeString(username)))
					return
				}
			}
			renderErrorPage(c, http.StatusInternalServerError, fmt.Sprintf("Failed to fetch user data: %v", err))
			return
		}
		userInfo = &info
		common.SetUserInfoCache(username, userInfo)
	}

	// Create and render card
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic in card rendering for user %s: %v", username, r)
		}
	}()

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
		log.Printf("Error: Empty SVG generated for user %s", username)
		renderErrorPage(c, http.StatusInternalServerError, "Failed to generate trophy card")
		return
	}

	// Set cache headers
	cacheControl := fmt.Sprintf("public, max-age=%d, s-maxage=%d, stale-while-revalidate=%d",
		common.CacheMaxAge, common.CDNCacheMaxAge, common.StaleWhileRevalidate)

	c.Header("Content-Type", "image/svg+xml")
	c.Header("Cache-Control", cacheControl)
	c.String(http.StatusOK, svg)
}

// renderErrorPage renders an error page
func renderErrorPage(c *gin.Context, statusCode int, message string) {
	baseURL := c.Request.URL.Scheme + "://" + c.Request.Host + c.Request.URL.Path
	messageEscaped := html.EscapeString(message)
	baseURLEscaped := html.EscapeString(baseURL)

	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Error</title>
  <style>
    body {
      font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif;
      max-width: 600px;
      margin: 50px auto;
      padding: 20px;
    }
    h2 { color: #d73a49; }
    code { background: #f6f8fa; padding: 2px 6px; border-radius: 3px; }
    form { margin-top: 20px; }
    input, button { padding: 8px; margin: 5px; }
  </style>
</head>
<body>
  <section>
    <div>
      <h2>%s</h2>
      <p>The URL should look like:</p>
      <div>
        <p id="base-show">%s?username=USERNAME</p>
        <button onclick="copyBaseUrl()">Copy Base Url</button>
        <span id="temporary-span"></span>
      </div>
      <p>where <code>USERNAME</code> is <em>your GitHub username.</em></p>
    </div>
    <div>
      <h2>You can use this form:</h2>
      <p>Enter your username and click get trophies</p>
      <form action="%s" method="get">
        <label for="username">GitHub Username</label>
        <input type="text" name="username" id="username" placeholder="Ex. ryo-ma" required>
        <label for="theme">Theme (Optional)</label>
        <input type="text" name="theme" id="theme" placeholder="Ex. onedark" value="default">
        <p>See all the available themes <a href="https://github.com/ryo-ma/github-profile-trophy?tab=readme-ov-file#apply-theme" target="_blank">here</a></p>
        <br>
        <button type="submit">Get Trophy's</button>
      </form>
    </div>
  </section>
  <script>
    function copyBaseUrl() {
      const text = document.querySelector("#base-show").textContent;
      navigator.clipboard.writeText(text);
      document.querySelector("#temporary-span").textContent = "Copied!";
      setTimeout(() => {
        document.querySelector("#temporary-span").textContent = "";
      }, 1500);
    }
  </script>
</body>
</html>`, messageEscaped, baseURLEscaped, baseURLEscaped)

	c.Header("Content-Type", "text/html")
	c.String(statusCode, html)
}
