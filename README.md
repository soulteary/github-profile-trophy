# GitHub Profile Trophy (Go Implementation)

[![GitHub](https://img.shields.io/badge/GitHub-soulteary%2Fgithub--profile--trophy-blue)](https://github.com/soulteary/github-profile-trophy)

![GitHub Profile Trophy](.github/assets/banner.png)

## Languages / 语言 / Sprachen / Lingue / 언어 / 言語

- [English](README.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)
- [Italiano](README.it.md)
- [한국어](README.kr.md)
- [日本語](README.ja.md)

## 🚀 Zero-Configuration, Drop-in Replacement

**No deployment needed!** This is a **100% compatible Go implementation** of the [GitHub Profile Trophy](https://github.com/ryo-ma/github-profile-trophy) project. You can use it as a **direct replacement** for the original service - just swap the URL and all your existing parameters will work exactly the same.

### ✨ Why Choose This Implementation?

| Feature | Original Project | This Project |
|---------|-----------------|--------------|
| **Deployment** | Requires Vercel/Cloud hosting | ✅ Self-hosted, full control |
| **API Compatibility** | - | ✅ 100% compatible, same parameters |
| **Performance** | Node.js runtime | ⚡ Go runtime, faster & lighter |
| **Rate Limits** | Single token | 🔄 Multi-token support |
| **Caching** | Limited | 💾 Memory + Redis support |
| **Maintenance** | Depends on service availability | 🛡️ You control the service |
| **Cost** | May require paid hosting | 💰 Free self-hosting |

### 🎯 Key Advantages

- 🎯 **100% API Compatible** - Use the exact same URL parameters as the original project
- 🚀 **No Deployment Required** - Self-hosted solution, full control over your data
- ⚡ **High Performance** - Built with Go for better performance and lower resource usage
- 🔄 **Multi-Token Support** - Handle higher API rate limits with multiple GitHub tokens
- 💾 **Smart Caching** - Built-in memory cache + optional Redis support for faster responses
- 🎨 **20+ Beautiful Themes** - All original themes supported plus more
- 🛡️ **Production Ready** - Retry mechanisms, error handling, and robust architecture

### Quick Start - Just Replace the URL!

If you're already using the original GitHub Profile Trophy, simply replace the base URL:

**Before (Original):**
```markdown
[![trophy](https://github-profile-trophy.vercel.app/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

**After (This Project):**
```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

**All parameters work exactly the same!** No changes needed to your existing code.

Of course, **we recommend** using the GitHub Actions approach instead. Simply update the original request parameters in the Action file:

```yml
...
- name: Generate trophy card
  uses: soulteary/github-profile-trophy-action@v1.0.0
    with:
      options: 'username=${{ github.repository_owner }}&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy'
      path: .github/assets/trophy.svg
      token: ${{ secrets.GITHUB_TOKEN }}
```

## Features

- ✅ Trophy card generation with multiple ranks (SSS, SS, S, AAA, AA, A, B, C)
- ✅ 15+ trophy types (Stars, Commits, Followers, Issues, PRs, Repositories, Reviews, etc.)
- ✅ Secret trophies (MultiLanguage, AllSuperRank, AncientAccount, etc.)
- ✅ 20+ themes support
- ✅ Customizable layout (column, row, margins)
- ✅ Filtering by title and rank
- ✅ Caching support (memory + Redis)
- ✅ Multi-token GitHub API support with retry mechanism

## 📖 Usage Examples

All examples below use the same URL parameters as the original project. Just replace the base URL!

### Basic Usage

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma)](https://github.com/ryo-ma/github-profile-trophy)
```

![Basic Trophy](.github/assets/trophy-basic.svg)

### With Theme

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

![Themed Trophy](.github/assets/trophy-themed.svg)

### Filter by Titles

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&title=Stars,Followers)](https://github.com/ryo-ma/github-profile-trophy)
```

![Filtered by Titles](.github/assets/trophy-filtered-titles.svg)

### Filter by Ranks

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&rank=S,AAA)](https://github.com/ryo-ma/github-profile-trophy)
```

![Filtered by Ranks](.github/assets/trophy-filtered-ranks.svg)

### Custom Layout

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&column=3&row=2&margin-w=15&margin-h=15)](https://github.com/ryo-ma/github-profile-trophy)
```

![Custom Layout](.github/assets/trophy-custom-layout.svg)

> 💡 **Tip:** All URL parameters from the original project work identically here. No need to change your existing README code!

### Using in GitHub Actions

You can use [github-profile-trophy-action](https://github.com/soulteary/github-profile-trophy-action) to generate trophy cards in your CI/CD pipeline:

```yaml
name: Generate Trophy

on:
  schedule:
    - cron: "0 0 * * *" # Runs once daily at midnight
  workflow_dispatch:

jobs:
  generate:
    runs-on: ubuntu-latest
    
    permissions:
      contents: write
    
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Generate trophy card
        uses: soulteary/github-profile-trophy-action@v1.0.0
        with:
          options: 'username=${{ github.repository_owner }}&theme=gruvbox&column=7&margin-w=15&margin-h=15'
          path: .github/assets/trophy.svg
          token: ${{ secrets.GITHUB_TOKEN }}

      - name: Commit trophy card
        run: |
          git config user.name "github-actions[bot]"
          git config user.email "41898282+github-actions[bot]@users.noreply.github.com"
          git add .github/assets/trophy.svg
          git commit -m "Update trophy card" || exit 0
          git push
```

Then embed the generated image in your README:

```markdown
![Trophy](.github/assets/trophy.svg)
```

## 🚀 Quick Start

### Option 1: Docker (Recommended - Easiest)

```bash
# Run with Docker - no installation needed!
docker run -d \
  -p 8080:8080 \
  -e GITHUB_TOKEN1=your_github_token_here \
  --name github-profile-trophy \
  soulteary/github-profile-trophy:latest
```

That's it! Your service is now running at `http://localhost:8080` and ready to use with all your existing URLs.

### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/soulteary/github-profile-trophy.git
cd github-profile-trophy

# Build
go build -o github-profile-trophy ./cmd/server

# Run (set your GitHub token)
GITHUB_TOKEN1=your_github_token_here ./github-profile-trophy
```

### Option 3: Go Install

```bash
go install github.com/soulteary/github-profile-trophy/cmd/server@latest
```

### Environment Variables

Create a `.env` file or set environment variables:

```bash
# GitHub Personal Access Token (required)
GITHUB_TOKEN1=your_github_token_here
# You can configure multiple tokens to increase API rate limits
GITHUB_TOKEN2=your_second_token_here

# Server port (optional, default: 8080)
PORT=8080

# Cache configuration (optional)
ENABLE_REDIS=false
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_USERNAME=
REDIS_PASSWORD=

# Production mode (optional)
NODE_ENV=production
```

> ⚡ **Performance Tip:** Configure multiple `GITHUB_TOKEN1`, `GITHUB_TOKEN2`, etc. to handle higher API rate limits automatically.

## 🎨 Available Themes

Choose from 20+ beautiful themes! All themes from the original project are supported.

### Popular Themes

<details>
<summary>Click to view all themes</summary>

### default

![default theme](.github/assets/theme-default.svg)

### flat

![flat theme](.github/assets/theme-flat.svg)

### onedark

![onedark theme](.github/assets/theme-onedark.svg)

### gruvbox

![gruvbox theme](.github/assets/theme-gruvbox.svg)

### dracula

![dracula theme](.github/assets/theme-dracula.svg)

### monokai

![monokai theme](.github/assets/theme-monokai.svg)

### chalk

![chalk theme](.github/assets/theme-chalk.svg)

### nord

![nord theme](.github/assets/theme-nord.svg)

### alduin

![alduin theme](.github/assets/theme-alduin.svg)

### darkhub

![darkhub theme](.github/assets/theme-darkhub.svg)

### juicyfresh

![juicyfresh theme](.github/assets/theme-juicyfresh.svg)

### oldie

![oldie theme](.github/assets/theme-oldie.svg)

### buddhism

![buddhism theme](.github/assets/theme-buddhism.svg)

### radical

![radical theme](.github/assets/theme-radical.svg)

### onestar

![onestar theme](.github/assets/theme-onestar.svg)

### discord

![discord theme](.github/assets/theme-discord.svg)

### algolia

![algolia theme](.github/assets/theme-algolia.svg)

### gitdimmed

![gitdimmed theme](.github/assets/theme-gitdimmed.svg)

### tokyonight

![tokyonight theme](.github/assets/theme-tokyonight.svg)

### matrix

![matrix theme](.github/assets/theme-matrix.svg)

### apprentice

![apprentice theme](.github/assets/theme-apprentice.svg)

### dark_dimmed

![dark_dimmed theme](.github/assets/theme-dark_dimmed.svg)

### dark_lover

![dark_lover theme](.github/assets/theme-dark_lover.svg)

### kimbie_dark

![kimbie_dark theme](.github/assets/theme-kimbie_dark.svg)

### aura

![aura theme](.github/assets/theme-aura.svg)

</details>

## 📋 API Parameters

**100% compatible with the original project!** All parameters work exactly the same.

| Parameter | Description | Default | Example |
|-----------|-------------|---------|---------|
| `username` | GitHub username (required) | - | `?username=ryo-ma` |
| `theme` | Theme name | `"default"` | `&theme=onedark` |
| `title` | Filter by trophy titles (comma-separated, use `-` prefix to exclude) | All | `&title=Stars,Followers` |
| `rank` | Filter by ranks (comma-separated, use `-` prefix to exclude) | All | `&rank=S,AAA` |
| `column` | Maximum number of columns (use `-1` for adaptive) | `8` | `&column=7` |
| `row` | Maximum number of rows | `3` | `&row=2` |
| `margin-w` | Horizontal margin between trophies | `0` | `&margin-w=15` |
| `margin-h` | Vertical margin between trophies | `0` | `&margin-h=15` |
| `no-bg` | Transparent background | `false` | `&no-bg=true` |
| `no-frame` | Hide frames | `false` | `&no-frame=true` |

## 🏆 Trophy Types

### Base Trophies
- Stars
- Commits
- Followers
- Issues
- Pull Requests
- Repositories
- Reviews

### Secret Trophies
- MultiLanguage (10+ languages)
- AllSuperRank (all base trophies are S rank or higher)
- LongTimeUser (10+ years)
- AncientUser (before 2010)
- OGUser (before 2008)
- Joined2020 (joined in 2020)
- Organizations (3+ organizations)
- Experience (account duration)

## Rank System

Ranks are: `SECRET`, `SSS`, `SS`, `S`, `AAA`, `AA`, `A`, `B`, `C`, `UNKNOWN`

## Project Structure

```
.
├── cmd/
│   └── server/          # Server entry point
│       └── main.go
├── internal/
│   ├── api/             # API handlers
│   ├── cards/           # Card rendering logic
│   ├── common/          # Common utilities
│   ├── fetchers/        # Data fetchers (GitHub API)
│   ├── themes/          # Theme system
│   └── trophies/        # Trophy system
├── go.mod
└── README.md
```

## ✅ Development Status

**Production Ready!** All core features are complete and tested.

- ✅ Project base structure
- ✅ HTTP server (Gin framework)
- ✅ GitHub API integration (GraphQL)
- ✅ Retry mechanism and multi-token support
- ✅ Cache handling (memory + Redis)
- ✅ Theme system (20+ themes)
- ✅ All trophy types (15+ types)
- ✅ Rank calculation
- ✅ All API endpoints
- ✅ 100% API compatibility with original project

## Contributing

Contributions are welcome! If you have any ideas or find issues, please:
1. Fork this project
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## License

This project is licensed under the MIT License.
