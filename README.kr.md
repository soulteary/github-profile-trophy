# GitHub Profile Trophy (Go 구현)

[![GitHub](https://img.shields.io/badge/GitHub-soulteary%2Fgithub--profile--trophy-blue)](https://github.com/soulteary/github-profile-trophy)

![GitHub Profile Trophy](.github/assets/banner.png)

## Languages / 语言 / Sprachen / Lingue / 언어 / 言語

- [English](README.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)
- [Italiano](README.it.md)
- [한국어](README.kr.md)
- [日本語](README.ja.md)

## 🚀 제로 설정, 드롭인 대체

**배포가 필요 없습니다!** 이것은 [GitHub Profile Trophy](https://github.com/ryo-ma/github-profile-trophy) 프로젝트의 **100% 호환 Go 구현** 버전입니다. 원본 서비스의 **직접 대체품**으로 사용할 수 있습니다 - URL만 교체하면 모든 기존 매개변수가 정확히 동일하게 작동합니다.

### ✨ 이 구현을 선택하는 이유?

| 기능 | 원본 프로젝트 | 이 프로젝트 |
|------|-------------|------------|
| **배포** | Vercel/클라우드 호스팅 필요 | ✅ 셀프 호스팅, 완전한 제어 |
| **API 호환성** | - | ✅ 100% 호환, 동일한 매개변수 |
| **성능** | Node.js 런타임 | ⚡ Go 런타임, 더 빠르고 가벼움 |
| **속도 제한** | 단일 토큰 | 🔄 다중 토큰 지원 |
| **캐싱** | 제한적 | 💾 메모리 + Redis 지원 |
| **유지보수** | 서비스 가용성에 의존 | 🛡️ 서비스를 직접 제어 |
| **비용** | 유료 호스팅 필요할 수 있음 | 💰 무료 셀프 호스팅 |

### 🎯 주요 장점

- 🎯 **100% API 호환** - 원본 프로젝트와 정확히 동일한 URL 매개변수 사용
- 🚀 **배포 불필요** - 셀프 호스팅 솔루션, 데이터에 대한 완전한 제어
- ⚡ **고성능** - 더 나은 성능과 낮은 리소스 사용을 위해 Go로 구축
- 🔄 **다중 토큰 지원** - 여러 GitHub 토큰으로 더 높은 API 속도 제한 처리
- 💾 **스마트 캐싱** - 내장 메모리 캐시 + 선택적 Redis 지원으로 더 빠른 응답
- 🎨 **20+ 아름다운 테마** - 모든 원본 테마 지원 및 더 많은 테마
- 🛡️ **프로덕션 준비 완료** - 재시도 메커니즘, 오류 처리 및 견고한 아키텍처

### 빠른 시작 - URL만 교체하세요!

이미 원본 GitHub Profile Trophy를 사용 중이라면 기본 URL만 교체하세요:

**이전 (원본):**
```markdown
[![trophy](https://github-profile-trophy.vercel.app/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

**이후 (이 프로젝트):**
```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

**모든 매개변수가 정확히 동일하게 작동합니다!** 기존 코드를 변경할 필요가 없습니다.

물론, **권장**하는 방법은 GitHub Actions 방식을 사용하는 것입니다. Action 파일에서 원본 요청 매개변수만 업데이트하면 됩니다:

```yml
...
- name: Generate trophy card
  uses: soulteary/github-profile-trophy-action@v1.0.0
    with:
      options: 'username=${{ github.repository_owner }}&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy'
      path: .github/assets/trophy.svg
      token: ${{ secrets.GITHUB_TOKEN }}
```

## 기능

- ✅ 여러 등급의 트로피 카드 생성 (SSS, SS, S, AAA, AA, A, B, C)
- ✅ 15개 이상의 트로피 유형 (Stars, Commits, Followers, Issues, PRs, Repositories, Reviews 등)
- ✅ 비밀 트로피 (MultiLanguage, AllSuperRank, AncientAccount 등)
- ✅ 20개 이상의 테마 지원
- ✅ 사용자 정의 가능한 레이아웃 (열, 행, 여백)
- ✅ 제목 및 등급별 필터링
- ✅ 캐시 지원 (메모리 + Redis)
- ✅ 재시도 메커니즘이 있는 다중 토큰 GitHub API 지원

## 📖 사용 예제

아래의 모든 예제는 원본 프로젝트와 동일한 URL 매개변수를 사용합니다. 기본 URL만 교체하세요!

### 기본 사용법

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma)](https://github.com/ryo-ma/github-profile-trophy)
```

![기본 트로피](.github/assets/trophy-basic.svg)

### 테마 사용

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

![테마 트로피](.github/assets/trophy-themed.svg)

### 제목별 필터링

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&title=Stars,Followers)](https://github.com/ryo-ma/github-profile-trophy)
```

![제목별 필터링](.github/assets/trophy-filtered-titles.svg)

### 등급별 필터링

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&rank=S,AAA)](https://github.com/ryo-ma/github-profile-trophy)
```

![등급별 필터링](.github/assets/trophy-filtered-ranks.svg)

### 사용자 정의 레이아웃

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&column=3&row=2&margin-w=15&margin-h=15)](https://github.com/ryo-ma/github-profile-trophy)
```

![사용자 정의 레이아웃](.github/assets/trophy-custom-layout.svg)

> 💡 **팁:** 원본 프로젝트의 모든 URL 매개변수가 여기서도 동일하게 작동합니다. 기존 README 코드를 변경할 필요가 없습니다!

### GitHub Actions에서 사용

[github-profile-trophy-action](https://github.com/soulteary/github-profile-trophy-action)을 사용하여 CI/CD 파이프라인에서 트로피 카드를 생성할 수 있습니다:

```yaml
name: Generate Trophy

on:
  schedule:
    - cron: "0 0 * * *" # 매일 자정에 한 번 실행
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

그런 다음 생성된 이미지를 README에 임베드하세요:

```markdown
![Trophy](.github/assets/trophy.svg)
```

## 🚀 빠른 시작

### 옵션 1: Docker (권장 - 가장 쉬움)

```bash
# Docker로 실행 - 설치 불필요!
docker run -d \
  -p 8080:8080 \
  -e GITHUB_TOKEN1=your_github_token_here \
  --name github-profile-trophy \
  soulteary/github-profile-trophy:latest
```

완료! 서비스가 이제 `http://localhost:8080`에서 실행 중이며 모든 기존 URL과 함께 사용할 준비가 되었습니다.

### 옵션 2: 소스에서 빌드

```bash
# 저장소 클론
git clone https://github.com/soulteary/github-profile-trophy.git
cd github-profile-trophy

# 빌드
go build -o github-profile-trophy ./cmd/server

# 실행 (GitHub 토큰 설정)
GITHUB_TOKEN1=your_github_token_here ./github-profile-trophy
```

### 옵션 3: Go 설치

```bash
go install github.com/soulteary/github-profile-trophy/cmd/server@latest
```

### 환경 변수

`.env` 파일을 생성하거나 환경 변수를 설정하세요:

```bash
# GitHub Personal Access Token (필수)
GITHUB_TOKEN1=your_github_token_here
# API 제한을 늘리기 위해 여러 토큰을 구성할 수 있습니다
GITHUB_TOKEN2=your_second_token_here

# 서버 포트 (선택 사항, 기본값: 8080)
PORT=8080

# 캐시 구성 (선택 사항)
ENABLE_REDIS=false
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_USERNAME=
REDIS_PASSWORD=

# 프로덕션 모드 (선택 사항)
NODE_ENV=production
```

> ⚡ **성능 팁:** 여러 `GITHUB_TOKEN1`, `GITHUB_TOKEN2` 등을 구성하여 더 높은 API 속도 제한을 자동으로 처리할 수 있습니다.

## 🎨 사용 가능한 테마

20+ 아름다운 테마 중에서 선택하세요! 원본 프로젝트의 모든 테마가 지원됩니다.

### 인기 테마

<details>
<summary>모든 테마 보기</summary>

## 사용 가능한 테마

### default

![default 테마](.github/assets/theme-default.svg)

### flat

![flat 테마](.github/assets/theme-flat.svg)

### onedark

![onedark 테마](.github/assets/theme-onedark.svg)

### gruvbox

![gruvbox 테마](.github/assets/theme-gruvbox.svg)

### dracula

![dracula 테마](.github/assets/theme-dracula.svg)

### monokai

![monokai 테마](.github/assets/theme-monokai.svg)

### chalk

![chalk 테마](.github/assets/theme-chalk.svg)

### nord

![nord 테마](.github/assets/theme-nord.svg)

### alduin

![alduin 테마](.github/assets/theme-alduin.svg)

### darkhub

![darkhub 테마](.github/assets/theme-darkhub.svg)

### juicyfresh

![juicyfresh 테마](.github/assets/theme-juicyfresh.svg)

### oldie

![oldie 테마](.github/assets/theme-oldie.svg)

### buddhism

![buddhism 테마](.github/assets/theme-buddhism.svg)

### radical

![radical 테마](.github/assets/theme-radical.svg)

### onestar

![onestar 테마](.github/assets/theme-onestar.svg)

### discord

![discord 테마](.github/assets/theme-discord.svg)

### algolia

![algolia 테마](.github/assets/theme-algolia.svg)

### gitdimmed

![gitdimmed 테마](.github/assets/theme-gitdimmed.svg)

### tokyonight

![tokyonight 테마](.github/assets/theme-tokyonight.svg)

### matrix

![matrix 테마](.github/assets/theme-matrix.svg)

### apprentice

![apprentice 테마](.github/assets/theme-apprentice.svg)

### dark_dimmed

![dark_dimmed 테마](.github/assets/theme-dark_dimmed.svg)

### dark_lover

![dark_lover 테마](.github/assets/theme-dark_lover.svg)

### kimbie_dark

![kimbie_dark 테마](.github/assets/theme-kimbie_dark.svg)

### aura

![aura 테마](.github/assets/theme-aura.svg)

</details>

## 📋 API 매개변수

**원본 프로젝트와 100% 호환됩니다!** 모든 매개변수가 정확히 동일하게 작동합니다.

| 매개변수 | 설명 | 기본값 | 예제 |
|---------|------|--------|------|
| `username` | GitHub 사용자 이름 (필수) | - | `?username=ryo-ma` |
| `theme` | 테마 이름 | `"default"` | `&theme=onedark` |
| `title` | 트로피 제목별 필터링 (쉼표로 구분, `-` 접두사를 사용하여 제외) | 모두 | `&title=Stars,Followers` |
| `rank` | 등급별 필터링 (쉼표로 구분, `-` 접두사를 사용하여 제외) | 모두 | `&rank=S,AAA` |
| `column` | 최대 열 수 (적응형은 `-1` 사용) | `8` | `&column=7` |
| `row` | 최대 행 수 | `3` | `&row=2` |
| `margin-w` | 트로피 간 수평 여백 | `0` | `&margin-w=15` |
| `margin-h` | 트로피 간 수직 여백 | `0` | `&margin-h=15` |
| `no-bg` | 투명 배경 | `false` | `&no-bg=true` |
| `no-frame` | 프레임 숨기기 | `false` | `&no-frame=true` |

## 🏆 트로피 유형

### 기본 트로피
- Stars
- Commits
- Followers
- Issues
- Pull Requests
- Repositories
- Reviews

### 비밀 트로피
- MultiLanguage (10개 이상의 언어)
- AllSuperRank (모든 기본 트로피가 S 등급 이상)
- LongTimeUser (10년 이상)
- AncientUser (2010년 이전)
- OGUser (2008년 이전)
- Joined2020 (2020년 가입)
- Organizations (3개 이상의 조직)
- Experience (계정 기간)

## 등급 시스템

등급은 다음과 같습니다: `SECRET`, `SSS`, `SS`, `S`, `AAA`, `AA`, `A`, `B`, `C`, `UNKNOWN`

## 프로젝트 구조

```
.
├── cmd/
│   └── server/          # 서버 진입점
│       └── main.go
├── internal/
│   ├── api/             # API 핸들러
│   ├── cards/           # 카드 렌더링 로직
│   ├── common/          # 공통 유틸리티
│   ├── fetchers/        # 데이터 페처 (GitHub API)
│   ├── themes/          # 테마 시스템
│   └── trophies/        # 트로피 시스템
├── go.mod
└── README.md
```

## ✅ 개발 상태

**프로덕션 준비 완료!** 모든 핵심 기능이 완료되고 테스트되었습니다.

- ✅ 프로젝트 기본 구조
- ✅ HTTP 서버 (Gin 프레임워크)
- ✅ GitHub API 통합 (GraphQL)
- ✅ 재시도 메커니즘 및 다중 토큰 지원
- ✅ 캐시 처리 (메모리 + Redis)
- ✅ 테마 시스템 (20+ 테마)
- ✅ 모든 트로피 유형 (15+ 유형)
- ✅ 등급 계산
- ✅ 모든 API 엔드포인트
- ✅ 원본 프로젝트와 100% API 호환

## 기여

기여를 환영합니다! 아이디어가 있거나 문제를 발견한 경우:
1. 이 프로젝트를 포크하세요
2. 기능 브랜치를 만드세요
3. 변경 사항을 커밋하세요
4. 브랜치에 푸시하세요
5. Pull Request를 열어주세요

## 라이선스

이 프로젝트는 MIT 라이선스 하에 있습니다.
