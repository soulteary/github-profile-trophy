package common

const (
	// Cache settings
	CacheMaxAge          = 18800
	CDNCacheMaxAge       = 28800              // 8 hours for CDN edge cache
	StaleWhileRevalidate = 86400              // 24 hours - serve stale while revalidating
	RevalidateTime       = 60 * 60 * 1000     // 1 hour in milliseconds
	RedisTTL             = 60 * 60 * 1000 * 4 // 4 hours in milliseconds

	// Panel settings
	DefaultPanelSize    = 110
	DefaultMaxColumn    = 8
	DefaultMaxRow       = 3
	DefaultMarginW      = 0
	DefaultMarginH      = 0
	DefaultNoBackground = false
	DefaultNoFrame      = false

	// GitHub API settings
	DefaultGitHubAPI        = "https://api.github.com/graphql"
	DefaultGitHubRetryDelay = 500 // milliseconds
)

// Rank represents trophy rank levels
type Rank string

const (
	RankSecret  Rank = "SECRET"
	RankSSS     Rank = "SSS"
	RankSS      Rank = "SS"
	RankS       Rank = "S"
	RankAAA     Rank = "AAA"
	RankAA      Rank = "AA"
	RankA       Rank = "A"
	RankB       Rank = "B"
	RankC       Rank = "C"
	RankUnknown Rank = "?"
)

// RankOrder defines the order of ranks from highest to lowest
var RankOrder = []Rank{
	RankSecret,
	RankSSS,
	RankSS,
	RankS,
	RankAAA,
	RankAA,
	RankA,
	RankB,
	RankC,
	RankUnknown,
}

// GetRankIndex returns the index of a rank in RankOrder
func GetRankIndex(rank Rank) int {
	for i, r := range RankOrder {
		if r == rank {
			return i
		}
	}
	return len(RankOrder) - 1 // Return last index for unknown ranks
}
