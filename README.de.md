# GitHub Profile Trophy (Go Implementierung)

[![GitHub](https://img.shields.io/badge/GitHub-soulteary%2Fgithub--profile--trophy-blue)](https://github.com/soulteary/github-profile-trophy)

![GitHub Profile Trophy](.github/assets/banner.png)

## Languages / 语言 / Sprachen / Lingue / 언어 / 言語

- [English](README.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)
- [Italiano](README.it.md)
- [한국어](README.kr.md)
- [日本語](README.ja.md)

## 🚀 Null-Konfiguration, Drop-in-Ersatz

**Keine Bereitstellung erforderlich!** Dies ist eine **100% kompatible Go-Implementierung** des [GitHub Profile Trophy](https://github.com/ryo-ma/github-profile-trophy) Projekts. Sie können es als **direkten Ersatz** für den ursprünglichen Service verwenden - tauschen Sie einfach die URL aus und alle Ihre vorhandenen Parameter funktionieren genau gleich.

### ✨ Warum diese Implementierung wählen?

| Funktion | Originalprojekt | Dieses Projekt |
|----------|----------------|----------------|
| **Bereitstellung** | Erfordert Vercel/Cloud-Hosting | ✅ Self-hosted, volle Kontrolle |
| **API-Kompatibilität** | - | ✅ 100% kompatibel, gleiche Parameter |
| **Leistung** | Node.js-Laufzeit | ⚡ Go-Laufzeit, schneller & leichter |
| **Rate Limits** | Einzelner Token | 🔄 Multi-Token-Unterstützung |
| **Caching** | Begrenzt | 💾 Speicher + Redis-Unterstützung |
| **Wartung** | Abhängig von Service-Verfügbarkeit | 🛡️ Sie kontrollieren den Service |
| **Kosten** | Erfordert möglicherweise kostenpflichtiges Hosting | 💰 Kostenloses Self-Hosting |

### 🎯 Hauptvorteile

- 🎯 **100% API-kompatibel** - Verwenden Sie exakt die gleichen URL-Parameter wie das Originalprojekt
- 🚀 **Keine Bereitstellung erforderlich** - Self-hosted-Lösung, volle Kontrolle über Ihre Daten
- ⚡ **Hohe Leistung** - Mit Go erstellt für bessere Leistung und geringeren Ressourcenverbrauch
- 🔄 **Multi-Token-Unterstützung** - Behandeln Sie höhere API-Rate-Limits mit mehreren GitHub-Tokens
- 💾 **Intelligentes Caching** - Integrierter Speicher-Cache + optionale Redis-Unterstützung für schnellere Antworten
- 🎨 **20+ schöne Themes** - Alle Original-Themes unterstützt und mehr
- 🛡️ **Produktionsbereit** - Wiederholungsmechanismen, Fehlerbehandlung und robuste Architektur

### Schnellstart - Ersetzen Sie einfach die URL!

Wenn Sie bereits das ursprüngliche GitHub Profile Trophy verwenden, ersetzen Sie einfach die Basis-URL:

**Vorher (Original):**
```markdown
[![trophy](https://github-profile-trophy.vercel.app/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

**Nachher (Dieses Projekt):**
```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

**Alle Parameter funktionieren genau gleich!** Keine Änderungen an Ihrem vorhandenen Code erforderlich.

Natürlich **empfehlen wir**, stattdessen den GitHub Actions-Ansatz zu verwenden. Aktualisieren Sie einfach die ursprünglichen Anforderungsparameter in der Action-Datei:

```yml
...
- name: Generate trophy card
  uses: soulteary/github-profile-trophy-action@v1.0.0
    with:
      options: 'username=${{ github.repository_owner }}&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy'
      path: .github/assets/trophy.svg
      token: ${{ secrets.GITHUB_TOKEN }}
```

## Funktionen

- ✅ Trophäen-Karten-Generierung mit mehreren Rängen (SSS, SS, S, AAA, AA, A, B, C)
- ✅ 15+ Trophäen-Typen (Stars, Commits, Followers, Issues, PRs, Repositories, Reviews, etc.)
- ✅ Geheime Trophäen (MultiLanguage, AllSuperRank, AncientAccount, etc.)
- ✅ 20+ Themen-Unterstützung
- ✅ Anpassbares Layout (Spalte, Zeile, Ränder)
- ✅ Filterung nach Titel und Rang
- ✅ Cache-Unterstützung (Speicher + Redis)
- ✅ Multi-Token GitHub API-Unterstützung mit Wiederholungsmechanismus

## 📖 Verwendungsbeispiele

Alle folgenden Beispiele verwenden die gleichen URL-Parameter wie das Originalprojekt. Ersetzen Sie einfach die Basis-URL!

### Grundlegende Verwendung

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma)](https://github.com/ryo-ma/github-profile-trophy)
```

![Grundlegende Trophäe](.github/assets/trophy-basic.svg)

### Mit Theme

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

![Themen-Trophäe](.github/assets/trophy-themed.svg)

### Nach Titeln filtern

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&title=Stars,Followers)](https://github.com/ryo-ma/github-profile-trophy)
```

![Nach Titeln gefiltert](.github/assets/trophy-filtered-titles.svg)

### Nach Rängen filtern

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&rank=S,AAA)](https://github.com/ryo-ma/github-profile-trophy)
```

![Nach Rängen gefiltert](.github/assets/trophy-filtered-ranks.svg)

### Benutzerdefiniertes Layout

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&column=3&row=2&margin-w=15&margin-h=15)](https://github.com/ryo-ma/github-profile-trophy)
```

![Benutzerdefiniertes Layout](.github/assets/trophy-custom-layout.svg)

> 💡 **Tipp:** Alle URL-Parameter des Originalprojekts funktionieren hier identisch. Keine Notwendigkeit, Ihren vorhandenen README-Code zu ändern!

### Verwendung in GitHub Actions

Sie können [github-profile-trophy-action](https://github.com/soulteary/github-profile-trophy-action) verwenden, um Trophäen-Karten in Ihrer CI/CD-Pipeline zu generieren:

```yaml
name: Generate Trophy

on:
  schedule:
    - cron: "0 0 * * *" # Läuft täglich um Mitternacht
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

Dann können Sie das generierte Bild in Ihrem README einbetten:

```markdown
![Trophy](.github/assets/trophy.svg)
```

## 🚀 Schnellstart

### Option 1: Docker (Empfohlen - Am einfachsten)

```bash
# Mit Docker ausführen - keine Installation erforderlich!
docker run -d \
  -p 8080:8080 \
  -e GITHUB_TOKEN1=your_github_token_here \
  --name github-profile-trophy \
  soulteary/github-profile-trophy:latest
```

Fertig! Ihr Service läuft jetzt auf `http://localhost:8080` und ist bereit, mit allen Ihren vorhandenen URLs verwendet zu werden.

### Option 2: Build aus Quellcode

```bash
# Repository klonen
git clone https://github.com/soulteary/github-profile-trophy.git
cd github-profile-trophy

# Build
go build -o github-profile-trophy ./cmd/server

# Ausführen (setzen Sie Ihren GitHub-Token)
GITHUB_TOKEN1=your_github_token_here ./github-profile-trophy
```

### Option 3: Go Install

```bash
go install github.com/soulteary/github-profile-trophy/cmd/server@latest
```

### Umgebungsvariablen

Erstellen Sie eine `.env` Datei oder setzen Sie Umgebungsvariablen:

```bash
# GitHub Personal Access Token (erforderlich)
GITHUB_TOKEN1=your_github_token_here
# Sie können mehrere Tokens konfigurieren, um API-Rate-Limits zu erhöhen
GITHUB_TOKEN2=your_second_token_here

# Server-Port (optional, Standard: 8080)
PORT=8080

# Cache-Konfiguration (optional)
ENABLE_REDIS=false
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_USERNAME=
REDIS_PASSWORD=

# Produktionsmodus (optional)
NODE_ENV=production
```

> ⚡ **Leistungstipp:** Konfigurieren Sie mehrere `GITHUB_TOKEN1`, `GITHUB_TOKEN2` usw., um automatisch höhere API-Rate-Limits zu behandeln.

## 🎨 Verfügbare Themes

Wählen Sie aus 20+ schönen Themes! Alle Themes des Originalprojekts werden unterstützt.

### Beliebte Themes

<details>
<summary>Klicken Sie, um alle Themes anzuzeigen</summary>

## Verfügbare Themes

### default

![default Theme](.github/assets/theme-default.svg)

### flat

![flat Theme](.github/assets/theme-flat.svg)

### onedark

![onedark Theme](.github/assets/theme-onedark.svg)

### gruvbox

![gruvbox Theme](.github/assets/theme-gruvbox.svg)

### dracula

![dracula Theme](.github/assets/theme-dracula.svg)

### monokai

![monokai Theme](.github/assets/theme-monokai.svg)

### chalk

![chalk Theme](.github/assets/theme-chalk.svg)

### nord

![nord Theme](.github/assets/theme-nord.svg)

### alduin

![alduin Theme](.github/assets/theme-alduin.svg)

### darkhub

![darkhub Theme](.github/assets/theme-darkhub.svg)

### juicyfresh

![juicyfresh Theme](.github/assets/theme-juicyfresh.svg)

### oldie

![oldie Theme](.github/assets/theme-oldie.svg)

### buddhism

![buddhism Theme](.github/assets/theme-buddhism.svg)

### radical

![radical Theme](.github/assets/theme-radical.svg)

### onestar

![onestar Theme](.github/assets/theme-onestar.svg)

### discord

![discord Theme](.github/assets/theme-discord.svg)

### algolia

![algolia Theme](.github/assets/theme-algolia.svg)

### gitdimmed

![gitdimmed Theme](.github/assets/theme-gitdimmed.svg)

### tokyonight

![tokyonight Theme](.github/assets/theme-tokyonight.svg)

### matrix

![matrix Theme](.github/assets/theme-matrix.svg)

### apprentice

![apprentice Theme](.github/assets/theme-apprentice.svg)

### dark_dimmed

![dark_dimmed Theme](.github/assets/theme-dark_dimmed.svg)

### dark_lover

![dark_lover Theme](.github/assets/theme-dark_lover.svg)

### kimbie_dark

![kimbie_dark Theme](.github/assets/theme-kimbie_dark.svg)

### aura

![aura Theme](.github/assets/theme-aura.svg)

</details>

## 📋 API-Parameter

**100% kompatibel mit dem Originalprojekt!** Alle Parameter funktionieren genau gleich.

| Parameter | Beschreibung | Standard | Beispiel |
|-----------|-------------|----------|----------|
| `username` | GitHub-Benutzername (erforderlich) | - | `?username=ryo-ma` |
| `theme` | Themenname | `"default"` | `&theme=onedark` |
| `title` | Nach Trophäen-Titeln filtern (kommagetrennt, verwenden Sie `-` Präfix zum Ausschließen) | Alle | `&title=Stars,Followers` |
| `rank` | Nach Rängen filtern (kommagetrennt, verwenden Sie `-` Präfix zum Ausschließen) | Alle | `&rank=S,AAA` |
| `column` | Maximale Anzahl von Spalten (verwenden Sie `-1` für adaptiv) | `8` | `&column=7` |
| `row` | Maximale Anzahl von Zeilen | `3` | `&row=2` |
| `margin-w` | Horizontaler Abstand zwischen Trophäen | `0` | `&margin-w=15` |
| `margin-h` | Vertikaler Abstand zwischen Trophäen | `0` | `&margin-h=15` |
| `no-bg` | Transparenter Hintergrund | `false` | `&no-bg=true` |
| `no-frame` | Rahmen ausblenden | `false` | `&no-frame=true` |

## 🏆 Trophäen-Typen

### Basis-Trophäen
- Stars
- Commits
- Followers
- Issues
- Pull Requests
- Repositories
- Reviews

### Geheime Trophäen
- MultiLanguage (10+ Sprachen)
- AllSuperRank (alle Basis-Trophäen sind S-Rang oder höher)
- LongTimeUser (10+ Jahre)
- AncientUser (vor 2010)
- OGUser (vor 2008)
- Joined2020 (2020 beigetreten)
- Organizations (3+ Organisationen)
- Experience (Kontodauer)

## Rang-System

Ränge sind: `SECRET`, `SSS`, `SS`, `S`, `AAA`, `AA`, `A`, `B`, `C`, `UNKNOWN`

## Projektstruktur

```
.
├── cmd/
│   └── server/          # Server-Einstiegspunkt
│       └── main.go
├── internal/
│   ├── api/             # API-Handler
│   ├── cards/           # Karten-Rendering-Logik
│   ├── common/          # Gemeinsame Utilities
│   ├── fetchers/        # Datenabrufer (GitHub API)
│   ├── themes/          # Themen-System
│   └── trophies/        # Trophäen-System
├── go.mod
└── README.md
```

## ✅ Entwicklungsstatus

**Produktionsbereit!** Alle Kernfunktionen sind abgeschlossen und getestet.

- ✅ Projekt-Basisstruktur
- ✅ HTTP-Server (Gin-Framework)
- ✅ GitHub API-Integration (GraphQL)
- ✅ Wiederholungsmechanismus und Multi-Token-Unterstützung
- ✅ Cache-Verwaltung (Speicher + Redis)
- ✅ Themen-System (20+ Themen)
- ✅ Alle Trophäen-Typen (15+ Typen)
- ✅ Rangberechnung
- ✅ Alle API-Endpunkte
- ✅ 100% API-Kompatibilität mit Originalprojekt

## Beitragen

Beiträge sind willkommen! Wenn Sie Ideen haben oder Probleme finden, bitte:
1. Forken Sie dieses Projekt
2. Erstellen Sie Ihren Feature-Branch
3. Committen Sie Ihre Änderungen
4. Pushen Sie zum Branch
5. Öffnen Sie einen Pull Request

## Lizenz

Dieses Projekt ist unter der MIT-Lizenz lizenziert.
