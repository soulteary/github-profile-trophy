package trophies

import (
	"fmt"
	"sort"

	"github.com/soulteary/github-profile-trophy/internal/common"
	"github.com/soulteary/github-profile-trophy/internal/themes"
)

// GetNextRankBar is a wrapper for common.GetNextRankBar
var GetNextRankBar = common.GetNextRankBar

// GetTrophyIcon is a wrapper for common.GetTrophyIcon
var GetTrophyIcon = common.GetTrophyIcon

// RankCondition represents a condition for achieving a rank
type RankCondition struct {
	Rank          common.Rank
	Message       string
	RequiredScore int
}

// Trophy represents a trophy with rank and score
type Trophy struct {
	RankCondition  *RankCondition
	Rank           common.Rank
	TopMessage     string
	BottomMessage  string
	Title          string
	FilterTitles   []string
	Hidden         bool
	Score          int
	RankConditions []RankCondition
}

// NewTrophy creates a new trophy
func NewTrophy(score int, rankConditions []RankCondition) *Trophy {
	t := &Trophy{
		Score:          score,
		RankConditions: rankConditions,
		Rank:           common.RankUnknown,
		TopMessage:     "Unknown",
		BottomMessage:  common.AbridgeScore(score),
	}
	t.SetRank()
	return t
}

// SetRank determines the rank based on score and conditions
func (t *Trophy) SetRank() {
	// Sort rank conditions by rank order (highest first)
	sortedConditions := make([]RankCondition, len(t.RankConditions))
	copy(sortedConditions, t.RankConditions)
	sort.Slice(sortedConditions, func(i, j int) bool {
		return common.GetRankIndex(sortedConditions[i].Rank) < common.GetRankIndex(sortedConditions[j].Rank)
	})

	// Find the first condition that matches
	for _, condition := range sortedConditions {
		if t.Score >= condition.RequiredScore {
			t.Rank = condition.Rank
			t.RankCondition = &condition
			t.TopMessage = condition.Message
			return
		}
	}
}

// CalculateNextRankPercentage calculates progress to next rank
func (t *Trophy) CalculateNextRankPercentage() float64 {
	if t.Rank == common.RankUnknown {
		return 0
	}

	currentIndex := common.GetRankIndex(t.Rank)
	nextIndex := currentIndex - 1

	// If already at max rank
	if nextIndex < 0 || t.Rank == common.RankSSS {
		return 1
	}

	nextRank := common.RankOrder[nextIndex]
	var nextCondition *RankCondition
	for i := range t.RankConditions {
		if t.RankConditions[i].Rank == nextRank {
			nextCondition = &t.RankConditions[i]
			break
		}
	}

	if nextCondition == nil || t.RankCondition == nil {
		return 0
	}

	distance := float64(nextCondition.RequiredScore - t.RankCondition.RequiredScore)
	if distance <= 0 {
		return 1
	}

	progress := float64(t.Score - t.RankCondition.RequiredScore)
	result := progress / distance
	if result > 1 {
		return 1
	}
	if result < 0 {
		return 0
	}
	return result
}

// Render generates SVG for the trophy
func (t *Trophy) Render(theme themes.Theme, x, y, panelSize int, noBackground, noFrame bool) string {
	nextRankBar := GetNextRankBar(t.Title, t.CalculateNextRankPercentage(), theme.NextRankBar)
	trophyIcon := GetTrophyIcon(theme, t.Rank)

	strokeOpacity := "1"
	if noFrame {
		strokeOpacity = "0"
	}
	fillOpacity := "1"
	if noBackground {
		fillOpacity = "0"
	}

	return fmt.Sprintf(`
        <svg
          x="%d"
          y="%d"
          width="%d"
          height="%d"
          viewBox="0 0 %d %d"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <rect
            x="0.5"
            y="0.5"
            rx="4.5"
            width="%d"
            height="%d"
            stroke="#e1e4e8"
            fill="%s"
            stroke-opacity="%s"
            fill-opacity="%s"
          />
          %s
          <text x="50%%" y="18" text-anchor="middle" font-family="Segoe UI,Helvetica,Arial,sans-serif,Apple Color Emoji,Segoe UI Emoji" font-weight="bold" font-size="13" fill="%s">%s</text>
          <text x="50%%" y="85" text-anchor="middle" font-family="Segoe UI,Helvetica,Arial,sans-serif,Apple Color Emoji,Segoe UI Emoji" font-weight="bold" font-size="10.5" fill="%s">%s</text>
          <text x="50%%" y="97" text-anchor="middle" font-family="Segoe UI,Helvetica,Arial,sans-serif,Apple Color Emoji,Segoe UI Emoji" font-weight="bold" font-size="10" fill="%s">%s</text>
          %s
        </svg>
        `,
		x, y, panelSize, panelSize, panelSize, panelSize,
		panelSize-1, panelSize-1,
		theme.Background, strokeOpacity, fillOpacity,
		trophyIcon,
		theme.Title, t.Title,
		theme.Text, t.TopMessage,
		theme.Text, t.BottomMessage,
		nextRankBar,
	)
}
