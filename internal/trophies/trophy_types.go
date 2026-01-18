package trophies

import "github.com/soulteary/github-profile-trophy/internal/common"

// NewTotalStarTrophy creates a Stars trophy
func NewTotalStarTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSSS, "Super Stargazer", 2000},
		{common.RankSS, "High Stargazer", 700},
		{common.RankS, "Stargazer", 200},
		{common.RankAAA, "Super Star", 100},
		{common.RankAA, "High Star", 50},
		{common.RankA, "You are a Star", 30},
		{common.RankB, "Middle Star", 10},
		{common.RankC, "First Star", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "Stars"
	t.FilterTitles = []string{"Star", "Stars"}
	return t
}

// NewTotalCommitTrophy creates a Commits trophy
func NewTotalCommitTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSSS, "God Committer", 4000},
		{common.RankSS, "Deep Committer", 2000},
		{common.RankS, "Super Committer", 1000},
		{common.RankAAA, "Ultra Committer", 500},
		{common.RankAA, "Hyper Committer", 200},
		{common.RankA, "High Committer", 100},
		{common.RankB, "Middle Committer", 10},
		{common.RankC, "First Commit", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "Commits"
	t.FilterTitles = []string{"Commit", "Commits"}
	return t
}

// NewTotalFollowerTrophy creates a Followers trophy
func NewTotalFollowerTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSSS, "Super Celebrity", 1000},
		{common.RankSS, "Ultra Celebrity", 400},
		{common.RankS, "Hyper Celebrity", 200},
		{common.RankAAA, "Famous User", 100},
		{common.RankAA, "Active User", 50},
		{common.RankA, "Dynamic User", 20},
		{common.RankB, "Many Friends", 10},
		{common.RankC, "First Friend", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "Followers"
	t.FilterTitles = []string{"Follower", "Followers"}
	return t
}

// NewTotalIssueTrophy creates an Issues trophy
func NewTotalIssueTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSSS, "God Issuer", 1000},
		{common.RankSS, "Deep Issuer", 500},
		{common.RankS, "Super Issuer", 200},
		{common.RankAAA, "Ultra Issuer", 100},
		{common.RankAA, "Hyper Issuer", 50},
		{common.RankA, "High Issuer", 20},
		{common.RankB, "Middle Issuer", 10},
		{common.RankC, "First Issue", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "Issues"
	t.FilterTitles = []string{"Issue", "Issues"}
	return t
}

// NewTotalPullRequestTrophy creates a Pull Requests trophy
func NewTotalPullRequestTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSSS, "God Puller", 1000},
		{common.RankSS, "Deep Puller", 500},
		{common.RankS, "Super Puller", 200},
		{common.RankAAA, "Ultra Puller", 100},
		{common.RankAA, "Hyper Puller", 50},
		{common.RankA, "High Puller", 20},
		{common.RankB, "Middle Puller", 10},
		{common.RankC, "First Pull", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "PullRequest"
	t.FilterTitles = []string{"PR", "PullRequest", "Pulls", "Puller"}
	return t
}

// NewTotalRepositoryTrophy creates a Repositories trophy
func NewTotalRepositoryTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSSS, "God Repo Creator", 50},
		{common.RankSS, "Deep Repo Creator", 45},
		{common.RankS, "Super Repo Creator", 40},
		{common.RankAAA, "Ultra Repo Creator", 35},
		{common.RankAA, "Hyper Repo Creator", 30},
		{common.RankA, "High Repo Creator", 20},
		{common.RankB, "Middle Repo Creator", 10},
		{common.RankC, "First Repository", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "Repositories"
	t.FilterTitles = []string{"Repo", "Repository", "Repositories"}
	return t
}

// NewTotalReviewsTrophy creates a Reviews trophy
func NewTotalReviewsTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSSS, "God Reviewer", 70},
		{common.RankSS, "Deep Reviewer", 57},
		{common.RankS, "Super Reviewer", 45},
		{common.RankAAA, "Ultra Reviewer", 30},
		{common.RankAA, "Hyper Reviewer", 20},
		{common.RankA, "Active Reviewer", 8},
		{common.RankB, "Intermediate Reviewer", 3},
		{common.RankC, "New Reviewer", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "Reviews"
	t.FilterTitles = []string{"Review", "Reviews"}
	return t
}

// NewMultipleLangTrophy creates a MultiLanguage trophy (secret)
func NewMultipleLangTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSecret, "Rainbow Lang User", 10},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "MultiLanguage"
	t.FilterTitles = []string{"MultipleLang", "MultiLanguage"}
	t.Hidden = true
	return t
}

// NewAllSuperRankTrophy creates an AllSuperRank trophy (secret)
func NewAllSuperRankTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSecret, "S Rank Hacker", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "AllSuperRank"
	t.FilterTitles = []string{"AllSuperRank"}
	t.BottomMessage = "All S Rank"
	t.Hidden = true
	return t
}

// NewJoined2020Trophy creates a Joined2020 trophy (secret)
func NewJoined2020Trophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSecret, "Everything started...", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "Joined2020"
	t.FilterTitles = []string{"Joined2020"}
	t.BottomMessage = "Joined 2020"
	t.Hidden = true
	return t
}

// NewAncientAccountTrophy creates an AncientAccount trophy (secret)
func NewAncientAccountTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSecret, "Ancient User", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "AncientUser"
	t.FilterTitles = []string{"AncientUser"}
	t.BottomMessage = "Before 2010"
	t.Hidden = true
	return t
}

// NewLongTimeAccountTrophy creates a LongTimeAccount trophy (secret)
func NewLongTimeAccountTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSecret, "Village Elder", 10},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "LongTimeUser"
	t.FilterTitles = []string{"LongTimeUser"}
	t.Hidden = true
	return t
}

// NewMultipleOrganizationsTrophy creates a MultipleOrganizations trophy (secret)
func NewMultipleOrganizationsTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSecret, "Jack of all Trades", 3},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "Organizations"
	t.FilterTitles = []string{"Organizations", "Orgs", "Teams"}
	t.Hidden = true
	return t
}

// NewOGAccountTrophy creates an OGAccount trophy (secret)
func NewOGAccountTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSecret, "OG User", 1},
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "OGUser"
	t.FilterTitles = []string{"OGUser"}
	t.BottomMessage = "Joined 2008"
	t.Hidden = true
	return t
}

// NewAccountDurationTrophy creates an AccountDuration trophy
func NewAccountDurationTrophy(score int) *Trophy {
	rankConditions := []RankCondition{
		{common.RankSSS, "Seasoned Veteran", 70}, // 20 years
		{common.RankSS, "Grandmaster", 55},       // 15 years
		{common.RankS, "Master Dev", 40},         // 10 years
		{common.RankAAA, "Expert Dev", 28},       // 7.5 years
		{common.RankAA, "Experienced Dev", 18},   // 5 years
		{common.RankA, "Intermediate Dev", 11},   // 3 years
		{common.RankB, "Junior Dev", 6},          // 1.5 years
		{common.RankC, "Newbie", 2},              // 0.5 year
	}
	t := NewTrophy(score, rankConditions)
	t.Title = "Experience"
	t.FilterTitles = []string{"Experience", "Duration", "Since"}
	return t
}
