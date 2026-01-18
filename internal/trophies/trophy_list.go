package trophies

import (
	"github.com/soulteary/github-profile-trophy/internal/common"
	"github.com/soulteary/github-profile-trophy/internal/fetchers"
)

// TrophyList manages a collection of trophies
type TrophyList struct {
	trophies []*Trophy
}

// NewTrophyList creates a new trophy list from user info
func NewTrophyList(userInfo fetchers.UserInfo) *TrophyList {
	tl := &TrophyList{
		trophies: []*Trophy{},
	}

	// Base trophies
	tl.trophies = append(tl.trophies,
		NewTotalStarTrophy(userInfo.TotalStargazers),
		NewTotalCommitTrophy(userInfo.TotalCommits),
		NewTotalFollowerTrophy(userInfo.TotalFollowers),
		NewTotalIssueTrophy(userInfo.TotalIssues),
		NewTotalPullRequestTrophy(userInfo.TotalPullRequests),
		NewTotalRepositoryTrophy(userInfo.TotalRepositories),
		NewTotalReviewsTrophy(userInfo.TotalReviews),
	)

	// Secret trophies
	isAllSRank := tl.isAllSRank()
	tl.trophies = append(tl.trophies,
		NewAllSuperRankTrophy(isAllSRank),
		NewMultipleLangTrophy(userInfo.LanguageCount),
		NewLongTimeAccountTrophy(userInfo.DurationYear),
		NewAncientAccountTrophy(userInfo.AncientAccount),
		NewOGAccountTrophy(userInfo.OGAccount),
		NewJoined2020Trophy(userInfo.Joined2020),
		NewMultipleOrganizationsTrophy(userInfo.TotalOrganizations),
		NewAccountDurationTrophy(userInfo.DurationDays),
	)

	return tl
}

// isAllSRank checks if all base trophies are S rank or higher
func (tl *TrophyList) isAllSRank() int {
	if len(tl.trophies) < 7 {
		return 0
	}
	baseTrophies := tl.trophies[:7] // First 7 are base trophies
	for _, trophy := range baseTrophies {
		rankStr := string(trophy.Rank)
		if len(rankStr) == 0 || rankStr[0] != 'S' {
			return 0
		}
	}
	return 1
}

// FilterByHidden filters out hidden trophies unless they have a non-UNKNOWN rank
func (tl *TrophyList) FilterByHidden() {
	filtered := []*Trophy{}
	for _, trophy := range tl.trophies {
		if !trophy.Hidden || trophy.Rank != common.RankUnknown {
			filtered = append(filtered, trophy)
		}
	}
	tl.trophies = filtered
}

// FilterByTitles filters trophies by title
func (tl *TrophyList) FilterByTitles(titles []string) {
	filtered := []*Trophy{}
	for _, trophy := range tl.trophies {
		for _, filterTitle := range titles {
			for _, trophyTitle := range trophy.FilterTitles {
				if trophyTitle == filterTitle {
					filtered = append(filtered, trophy)
					goto next
				}
			}
		}
	next:
	}
	tl.trophies = filtered
}

// FilterByRanks filters trophies by rank
func (tl *TrophyList) FilterByRanks(ranks []string) {
	// Check if any rank starts with "-" (exclusion)
	hasExclusion := false
	for _, rank := range ranks {
		if len(rank) > 0 && rank[0] == '-' {
			hasExclusion = true
			break
		}
	}

	if hasExclusion {
		// Exclusion mode
		excludeRanks := []string{}
		for _, rank := range ranks {
			if len(rank) > 0 && rank[0] == '-' {
				excludeRanks = append(excludeRanks, rank[1:])
			}
		}
		filtered := []*Trophy{}
		for _, trophy := range tl.trophies {
			excluded := false
			for _, excludeRank := range excludeRanks {
				if string(trophy.Rank) == excludeRank {
					excluded = true
					break
				}
			}
			if !excluded {
				filtered = append(filtered, trophy)
			}
		}
		tl.trophies = filtered
	} else {
		// Inclusion mode
		filtered := []*Trophy{}
		for _, trophy := range tl.trophies {
			for _, filterRank := range ranks {
				if string(trophy.Rank) == filterRank {
					filtered = append(filtered, trophy)
					break
				}
			}
		}
		tl.trophies = filtered
	}
}

// FilterByExclusionTitles filters out trophies with excluded titles
func (tl *TrophyList) FilterByExclusionTitles(titles []string) {
	excludeTitles := []string{}
	for _, title := range titles {
		if len(title) > 0 && title[0] == '-' {
			excludeTitles = append(excludeTitles, title[1:])
		}
	}
	if len(excludeTitles) == 0 {
		return
	}

	filtered := []*Trophy{}
	for _, trophy := range tl.trophies {
		excluded := false
		for _, excludeTitle := range excludeTitles {
			if trophy.Title == excludeTitle {
				excluded = true
				break
			}
		}
		if !excluded {
			filtered = append(filtered, trophy)
		}
	}
	tl.trophies = filtered
}

// SortByRank sorts trophies by rank (highest first)
func (tl *TrophyList) SortByRank() {
	// Use stable sort to maintain order for same ranks
	for i := 0; i < len(tl.trophies)-1; i++ {
		for j := i + 1; j < len(tl.trophies); j++ {
			if common.GetRankIndex(tl.trophies[i].Rank) > common.GetRankIndex(tl.trophies[j].Rank) {
				tl.trophies[i], tl.trophies[j] = tl.trophies[j], tl.trophies[i]
			}
		}
	}
}

// GetArray returns the trophy array
func (tl *TrophyList) GetArray() []*Trophy {
	return tl.trophies
}

// Length returns the number of trophies
func (tl *TrophyList) Length() int {
	return len(tl.trophies)
}
