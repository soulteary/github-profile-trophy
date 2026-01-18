package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/soulteary/github-profile-trophy/internal/api"
	"github.com/soulteary/github-profile-trophy/internal/common"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		// .env file is optional
		log.Println("No .env file found, using environment variables")
	}

	// Initialize Redis if enabled
	common.InitRedis()

	// Check if running in CLI mode
	if len(os.Args) > 1 {
		if err := runCLI(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Set Gin mode
	if os.Getenv("NODE_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	r := gin.Default()

	// API routes
	r.GET("/", api.TrophyHandler)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Server running on port %s", port)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func runCLI() error {
	var outputPath string
	var params map[string]string

	// Check for help flag
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help") {
		printUsage()
		return nil
	}

	// Check if first argument is a URL
	if len(os.Args) > 1 && (strings.HasPrefix(os.Args[1], "/") || strings.HasPrefix(os.Args[1], "?")) {
		// URL format
		urlStr := os.Args[1]
		// If it starts with "?", prepend "/" to make it a valid path
		if strings.HasPrefix(urlStr, "?") {
			urlStr = "/" + urlStr
		}
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
			return fmt.Errorf("invalid URL: %v", err)
		}

		// Parse query parameters
		params = make(map[string]string)
		for key, values := range parsedURL.Query() {
			if len(values) > 0 {
				params[key] = values[0]
			}
		}

		// Check for --output flag in remaining args
		flagSet := flag.NewFlagSet("", flag.ContinueOnError)
		flagSet.StringVar(&outputPath, "output", "", "Output file path (default: stdout)")
		if len(os.Args) > 2 {
			flagSet.Parse(os.Args[2:])
		}
	} else {
		// Flag format - parse all arguments manually for flexibility
		args := os.Args[1:]
		params = make(map[string]string)

		// Parse arguments
		for i := 0; i < len(args); i++ {
			arg := args[i]
			if strings.HasPrefix(arg, "--") {
				if strings.Contains(arg, "=") {
					// --key=value format
					parts := strings.SplitN(arg[2:], "=", 2)
					if len(parts) == 2 {
						key := parts[0]
						value := parts[1]
						if key == "output" {
							outputPath = value
						} else {
							params[key] = value
						}
					}
				} else {
					// --key value format
					key := arg[2:]
					if key == "output" {
						if i+1 < len(args) {
							outputPath = args[i+1]
							i++
						}
					} else {
						if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
							params[key] = args[i+1]
							i++
						} else {
							// Boolean flag without value
							params[key] = "true"
						}
					}
				}
			} else if !strings.HasPrefix(arg, "-") {
				// Positional argument - treat as error or ignore
				return fmt.Errorf("unexpected positional argument: %s (use --key=value format or URL format)", arg)
			}
		}
	}

	// Generate trophy card
	return api.GenerateTrophyCard(params, outputPath)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage:
  %s [OPTIONS]                    # Start HTTP server
  %s "/?username=xxx&..."        # Generate card from URL
  %s --username=xxx [OPTIONS]     # Generate card with flags

URL Format:
  /?username=xxx&theme=gruvbox&column=7&margin-w=15&margin-h=15&title=AllSuperRank,MultiLanguage,Stars,Commits,Follower,Issues,PullRequest
  ?username=xxx&theme=gruvbox&column=7&margin-w=15&margin-h=15&title=AllSuperRank,MultiLanguage,Stars,Commits,Follower,Issues,PullRequest

Flag Format:
  --output=FILE                   # Output file path (default: stdout)
  --username=USER                 # GitHub username (required)
  --theme=THEME                  # Theme name (default: "default")
  --column=NUMBER                 # Maximum number of columns (default: 8)
  --row=NUMBER                    # Maximum number of rows (default: 3)
  --margin-w=NUMBER               # Horizontal margin between trophies (default: 0)
  --margin-h=NUMBER               # Vertical margin between trophies (default: 0)
  --title=TITLES                  # Filter by trophy titles (comma-separated, use - prefix to exclude)
  --rank=RANKS                    # Filter by ranks (comma-separated, use - prefix to exclude)
  --no-bg=BOOLEAN                 # Transparent background (default: false)
  --no-frame=BOOLEAN              # Hide frames (default: false)
  [other parameters...]           # Any other API parameters

Examples:
  %s "/?username=soulteary&theme=gruvbox&column=7&margin-w=15&margin-h=15&title=AllSuperRank,MultiLanguage,Stars,Commits,Follower,Issues,PullRequest" --output=trophy.svg
  %s --username=soulteary --theme=gruvbox --column=7 --margin-w=15 --margin-h=15 --title=AllSuperRank,MultiLanguage,Stars,Commits,Follower,Issues,PullRequest --output=trophy.svg
  %s "/?username=soulteary&theme=gruvbox" > trophy.svg

`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}
