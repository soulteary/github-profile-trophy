# GitHub Profile Trophy (Implementazione Go)

[![GitHub](https://img.shields.io/badge/GitHub-soulteary%2Fgithub--profile--trophy-blue)](https://github.com/soulteary/github-profile-trophy)

![GitHub Profile Trophy](.github/assets/banner.png)

## Languages / 语言 / Sprachen / Lingue / 언어 / 言語

- [English](README.md)
- [简体中文](README.zh.md)
- [Deutsch](README.de.md)
- [Italiano](README.it.md)
- [한국어](README.kr.md)
- [日本語](README.ja.md)

## 🚀 Zero Configurazione, Sostituto Drop-in

**Nessuna distribuzione necessaria!** Questa è un'**implementazione Go 100% compatibile** del progetto [GitHub Profile Trophy](https://github.com/ryo-ma/github-profile-trophy). Puoi usarla come **sostituto diretto** del servizio originale - basta sostituire l'URL e tutti i tuoi parametri esistenti funzioneranno esattamente allo stesso modo.

### ✨ Perché scegliere questa implementazione?

| Funzionalità | Progetto Originale | Questo Progetto |
|--------------|-------------------|-----------------|
| **Distribuzione** | Richiede Vercel/Cloud hosting | ✅ Self-hosted, controllo completo |
| **Compatibilità API** | - | ✅ 100% compatibile, stessi parametri |
| **Prestazioni** | Runtime Node.js | ⚡ Runtime Go, più veloce e leggero |
| **Limiti di Rate** | Token singolo | 🔄 Supporto multi-token |
| **Cache** | Limitata | 💾 Supporto memoria + Redis |
| **Manutenzione** | Dipende dalla disponibilità del servizio | 🛡️ Tu controlli il servizio |
| **Costo** | Potrebbe richiedere hosting a pagamento | 💰 Self-hosting gratuito |

### 🎯 Vantaggi Chiave

- 🎯 **100% Compatibile API** - Usa gli stessi parametri URL del progetto originale
- 🚀 **Nessuna Distribuzione Richiesta** - Soluzione self-hosted, controllo completo sui tuoi dati
- ⚡ **Alte Prestazioni** - Costruito con Go per prestazioni migliori e minore utilizzo di risorse
- 🔄 **Supporto Multi-Token** - Gestisci limiti di rate API più alti con più token GitHub
- 💾 **Cache Intelligente** - Cache memoria integrata + supporto Redis opzionale per risposte più veloci
- 🎨 **20+ Temi Bellissimi** - Tutti i temi originali supportati e altro ancora
- 🛡️ **Pronto per la Produzione** - Meccanismi di retry, gestione errori e architettura robusta

### Guida Rapida - Basta Sostituire l'URL!

Se stai già usando il GitHub Profile Trophy originale, sostituisci semplicemente l'URL di base:

**Prima (Originale):**
```markdown
[![trophy](https://github-profile-trophy.vercel.app/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

**Dopo (Questo Progetto):**
```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

**Tutti i parametri funzionano esattamente allo stesso modo!** Nessuna modifica al tuo codice esistente necessaria.

Naturalmente, **consigliamo** di usare invece l'approccio GitHub Actions. Basta aggiornare i parametri della richiesta originale nel file Action:

```yml
...
- name: Generate trophy card
  uses: soulteary/github-profile-trophy-action@v1.0.0
    with:
      options: 'username=${{ github.repository_owner }}&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy'
      path: .github/assets/trophy.svg
      token: ${{ secrets.GITHUB_TOKEN }}
```

## Funzionalità

- ✅ Generazione carte trofei con più ranghi (SSS, SS, S, AAA, AA, A, B, C)
- ✅ 15+ tipi di trofei (Stars, Commits, Followers, Issues, PRs, Repositories, Reviews, ecc.)
- ✅ Trofei segreti (MultiLanguage, AllSuperRank, AncientAccount, ecc.)
- ✅ Supporto per 20+ temi
- ✅ Layout personalizzabile (colonna, riga, margini)
- ✅ Filtraggio per titolo e rango
- ✅ Supporto cache (memoria + Redis)
- ✅ Supporto API GitHub multi-token con meccanismo di retry

## 📖 Esempi di Utilizzo

Tutti gli esempi seguenti usano gli stessi parametri URL del progetto originale. Basta sostituire l'URL di base!

### Utilizzo Base

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma)](https://github.com/ryo-ma/github-profile-trophy)
```

![Trofeo Base](.github/assets/trophy-basic.svg)

### Con Tema

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&theme=onedark)](https://github.com/ryo-ma/github-profile-trophy)
```

![Trofeo con Tema](.github/assets/trophy-themed.svg)

### Filtra per Titoli

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&title=Stars,Followers)](https://github.com/ryo-ma/github-profile-trophy)
```

![Filtrato per Titoli](.github/assets/trophy-filtered-titles.svg)

### Filtra per Ranghi

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&rank=S,AAA)](https://github.com/ryo-ma/github-profile-trophy)
```

![Filtrato per Ranghi](.github/assets/trophy-filtered-ranks.svg)

### Layout Personalizzato

```markdown
[![trophy](http://localhost:8080/?username=ryo-ma&column=3&row=2&margin-w=15&margin-h=15)](https://github.com/ryo-ma/github-profile-trophy)
```

![Layout Personalizzato](.github/assets/trophy-custom-layout.svg)

> 💡 **Suggerimento:** Tutti i parametri URL del progetto originale funzionano identicamente qui. Non è necessario modificare il tuo codice README esistente!

### Utilizzo in GitHub Actions

Puoi usare [github-profile-trophy-action](https://github.com/soulteary/github-profile-trophy-action) per generare carte trofei nella tua pipeline CI/CD:

```yaml
name: Generate Trophy

on:
  schedule:
    - cron: "0 0 * * *" # Esegue una volta al giorno a mezzanotte
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

Quindi incorpora l'immagine generata nel tuo README:

```markdown
![Trophy](.github/assets/trophy.svg)
```

## 🚀 Guida Rapida

### Opzione 1: Docker (Consigliato - Più Facile)

```bash
# Esegui con Docker - nessuna installazione necessaria!
docker run -d \
  -p 8080:8080 \
  -e GITHUB_TOKEN1=your_github_token_here \
  --name github-profile-trophy \
  soulteary/github-profile-trophy:latest
```

Fatto! Il tuo servizio è ora in esecuzione su `http://localhost:8080` e pronto per essere usato con tutti i tuoi URL esistenti.

### Opzione 2: Build dal Sorgente

```bash
# Clona il repository
git clone https://github.com/soulteary/github-profile-trophy.git
cd github-profile-trophy

# Build
go build -o github-profile-trophy ./cmd/server

# Esegui (imposta il tuo token GitHub)
GITHUB_TOKEN1=your_github_token_here ./github-profile-trophy
```

### Opzione 3: Go Install

```bash
go install github.com/soulteary/github-profile-trophy/cmd/server@latest
```

### Variabili d'Ambiente

Crea un file `.env` o imposta le variabili d'ambiente:

```bash
# GitHub Personal Access Token (richiesto)
GITHUB_TOKEN1=your_github_token_here
# Puoi configurare più token per aumentare i limiti dell'API
GITHUB_TOKEN2=your_second_token_here

# Porta del server (opzionale, predefinito: 8080)
PORT=8080

# Configurazione cache (opzionale)
ENABLE_REDIS=false
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_USERNAME=
REDIS_PASSWORD=

# Modalità produzione (opzionale)
NODE_ENV=production
```

> ⚡ **Suggerimento Prestazioni:** Configura più `GITHUB_TOKEN1`, `GITHUB_TOKEN2`, ecc. per gestire automaticamente limiti di rate API più alti.

## 🎨 Temi Disponibili

Scegli tra 20+ temi bellissimi! Tutti i temi del progetto originale sono supportati.

### Temi Popolari

<details>
<summary>Clicca per visualizzare tutti i temi</summary>

## Temi Disponibili

### default

![tema default](.github/assets/theme-default.svg)

### flat

![tema flat](.github/assets/theme-flat.svg)

### onedark

![tema onedark](.github/assets/theme-onedark.svg)

### gruvbox

![tema gruvbox](.github/assets/theme-gruvbox.svg)

### dracula

![tema dracula](.github/assets/theme-dracula.svg)

### monokai

![tema monokai](.github/assets/theme-monokai.svg)

### chalk

![tema chalk](.github/assets/theme-chalk.svg)

### nord

![tema nord](.github/assets/theme-nord.svg)

### alduin

![tema alduin](.github/assets/theme-alduin.svg)

### darkhub

![tema darkhub](.github/assets/theme-darkhub.svg)

### juicyfresh

![tema juicyfresh](.github/assets/theme-juicyfresh.svg)

### oldie

![tema oldie](.github/assets/theme-oldie.svg)

### buddhism

![tema buddhism](.github/assets/theme-buddhism.svg)

### radical

![tema radical](.github/assets/theme-radical.svg)

### onestar

![tema onestar](.github/assets/theme-onestar.svg)

### discord

![tema discord](.github/assets/theme-discord.svg)

### algolia

![tema algolia](.github/assets/theme-algolia.svg)

### gitdimmed

![tema gitdimmed](.github/assets/theme-gitdimmed.svg)

### tokyonight

![tema tokyonight](.github/assets/theme-tokyonight.svg)

### matrix

![tema matrix](.github/assets/theme-matrix.svg)

### apprentice

![tema apprentice](.github/assets/theme-apprentice.svg)

### dark_dimmed

![tema dark_dimmed](.github/assets/theme-dark_dimmed.svg)

### dark_lover

![tema dark_lover](.github/assets/theme-dark_lover.svg)

### kimbie_dark

![tema kimbie_dark](.github/assets/theme-kimbie_dark.svg)

### aura

![tema aura](.github/assets/theme-aura.svg)

</details>

## 📋 Parametri API

**100% compatibile con il progetto originale!** Tutti i parametri funzionano esattamente allo stesso modo.

| Parametro | Descrizione | Predefinito | Esempio |
|-----------|-------------|-------------|---------|
| `username` | Nome utente GitHub (richiesto) | - | `?username=ryo-ma` |
| `theme` | Nome del tema | `"default"` | `&theme=onedark` |
| `title` | Filtra per titoli dei trofei (separati da virgola, usa il prefisso `-` per escludere) | Tutti | `&title=Stars,Followers` |
| `rank` | Filtra per ranghi (separati da virgola, usa il prefisso `-` per escludere) | Tutti | `&rank=S,AAA` |
| `column` | Numero massimo di colonne (usa `-1` per adattivo) | `8` | `&column=7` |
| `row` | Numero massimo di righe | `3` | `&row=2` |
| `margin-w` | Margine orizzontale tra i trofei | `0` | `&margin-w=15` |
| `margin-h` | Margine verticale tra i trofei | `0` | `&margin-h=15` |
| `no-bg` | Sfondo trasparente | `false` | `&no-bg=true` |
| `no-frame` | Nascondi cornici | `false` | `&no-frame=true` |

## 🏆 Tipi di Trofei

### Trofei Base
- Stars
- Commits
- Followers
- Issues
- Pull Requests
- Repositories
- Reviews

### Trofei Segreti
- MultiLanguage (10+ lingue)
- AllSuperRank (tutti i trofei base sono rango S o superiore)
- LongTimeUser (10+ anni)
- AncientUser (prima del 2010)
- OGUser (prima del 2008)
- Joined2020 (iscritto nel 2020)
- Organizations (3+ organizzazioni)
- Experience (durata dell'account)

## Sistema di Ranghi

I ranghi sono: `SECRET`, `SSS`, `SS`, `S`, `AAA`, `AA`, `A`, `B`, `C`, `UNKNOWN`

## Struttura del Progetto

```
.
├── cmd/
│   └── server/          # Punto di ingresso del server
│       └── main.go
├── internal/
│   ├── api/             # Gestori API
│   ├── cards/           # Logica di rendering delle carte
│   ├── common/          # Utility comuni
│   ├── fetchers/        # Recuperatori di dati (GitHub API)
│   ├── themes/          # Sistema di temi
│   └── trophies/        # Sistema di trofei
├── go.mod
└── README.md
```

## ✅ Stato di Sviluppo

**Pronto per la Produzione!** Tutte le funzionalità principali sono completate e testate.

- ✅ Struttura base del progetto
- ✅ Server HTTP (Framework Gin)
- ✅ Integrazione GitHub API (GraphQL)
- ✅ Meccanismo di retry e supporto multi-token
- ✅ Gestione cache (memoria + Redis)
- ✅ Sistema di temi (20+ temi)
- ✅ Tutti i tipi di trofei (15+ tipi)
- ✅ Calcolo del rango
- ✅ Tutti gli endpoint API
- ✅ 100% compatibilità API con il progetto originale

## Contribuire

I contributi sono benvenuti! Se hai idee o trovi problemi, per favore:
1. Fai un fork di questo progetto
2. Crea il tuo branch delle funzionalità
3. Committa le tue modifiche
4. Pusha al branch
5. Apri una Pull Request

## Licenza

Questo progetto è concesso in licenza sotto la licenza MIT.
