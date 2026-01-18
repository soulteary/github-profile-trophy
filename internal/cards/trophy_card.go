package cards

import (
	"fmt"
	"strings"

	"github.com/soulteary/github-profile-trophy/internal/fetchers"
	"github.com/soulteary/github-profile-trophy/internal/themes"
	"github.com/soulteary/github-profile-trophy/internal/trophies"
)

// Card represents a trophy card
type Card struct {
	width        int
	height       int
	titles       []string
	ranks        []string
	maxColumn    int
	maxRow       int
	panelSize    int
	marginWidth  int
	marginHeight int
	noBackground bool
	noFrame      bool
}

// NewCard creates a new trophy card
func NewCard(
	titles []string,
	ranks []string,
	maxColumn int,
	maxRow int,
	panelSize int,
	marginWidth int,
	marginHeight int,
	noBackground bool,
	noFrame bool,
) *Card {
	width := panelSize*maxColumn + marginWidth*(maxColumn-1)
	if maxColumn == -1 {
		width = 0 // Will be calculated later
	}
	return &Card{
		width:        width,
		titles:       titles,
		ranks:        ranks,
		maxColumn:    maxColumn,
		maxRow:       maxRow,
		panelSize:    panelSize,
		marginWidth:  marginWidth,
		marginHeight: marginHeight,
		noBackground: noBackground,
		noFrame:      noFrame,
	}
}

// Render generates the SVG for the trophy card
func (c *Card) Render(userInfo fetchers.UserInfo, theme themes.Theme) string {
	trophyList := trophies.NewTrophyList(userInfo)

	trophyList.FilterByHidden()

	if len(c.titles) > 0 {
		includeTitles := []string{}
		for _, title := range c.titles {
			if !strings.HasPrefix(title, "-") {
				includeTitles = append(includeTitles, title)
			}
		}
		if len(includeTitles) > 0 {
			trophyList.FilterByTitles(includeTitles)
		}
		trophyList.FilterByExclusionTitles(c.titles)
	}

	if len(c.ranks) > 0 {
		trophyList.FilterByRanks(c.ranks)
	}

	trophyList.SortByRank()

	// Handle adaptive column
	if c.maxColumn == -1 {
		c.maxColumn = trophyList.Length()
		if c.maxColumn == 0 {
			c.maxColumn = 1
		}
		c.width = c.panelSize*c.maxColumn + c.marginWidth*(c.maxColumn-1)
	}

	// Ensure maxColumn is at least 1
	if c.maxColumn <= 0 {
		c.maxColumn = 1
		c.width = c.panelSize
	}

	row := c.getRow(trophyList)
	c.height = c.getHeight(row)

	trophySVGs := c.renderTrophies(trophyList, theme)

	return fmt.Sprintf(`
    <svg
      width="%d"
      height="%d"
      viewBox="0 0 %d %d"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      %s
    </svg>`, c.width, c.height, c.width, c.height, trophySVGs)
}

// getRow calculates the number of rows needed
func (c *Card) getRow(trophyList *trophies.TrophyList) int {
	length := trophyList.Length()
	if length == 0 {
		return 1
	}
	row := (length-1)/c.maxColumn + 1
	if row > c.maxRow {
		row = c.maxRow
	}
	if row <= 0 {
		row = 1
	}
	return row
}

// getHeight calculates the total height of the card
func (c *Card) getHeight(row int) int {
	return c.panelSize*row + c.marginHeight*(row-1)
}

// renderTrophies renders all trophies in the list
func (c *Card) renderTrophies(trophyList *trophies.TrophyList, theme themes.Theme) string {
	trophyArray := trophyList.GetArray()
	var svgParts []string

	for i, trophy := range trophyArray {
		currentColumn := i % c.maxColumn
		currentRow := i / c.maxColumn
		x := c.panelSize*currentColumn + c.marginWidth*currentColumn
		y := c.panelSize*currentRow + c.marginHeight*currentRow

		svgParts = append(svgParts, trophy.Render(theme, x, y, c.panelSize, c.noBackground, c.noFrame))
	}

	return strings.Join(svgParts, "")
}
