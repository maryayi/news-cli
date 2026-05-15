# news-cli

A fast command-line tool to fetch top news headlines, powered by [NewsAPI.org](https://newsapi.org). Run it with a keyword to search, or with no arguments to get the latest top headlines from around the world.

## Features

- General world headlines with no keyword required
- Search headlines by any keyword
- Clickable, blue-colored source links in supported terminals
- Configure the number of results with `-n`
- Results sorted by most recent publication date
- API key loaded automatically from a `.env` file

## Requirements

- Go 1.18 or later
- A free [NewsAPI.org](https://newsapi.org/register) API key

## Installation

### go install

```bash
go install github.com/maryayi/news-cli@latest
```

Then run it as:

```bash
news-cli --version
```

Make sure `$(go env GOPATH)/bin` is in your `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Build from source

```bash
git clone https://github.com/maryayi/news-cli.git
cd news-cli
go build -o $(go env GOPATH)/bin/news .
```

This builds the binary directly into your Go bin directory so `news` is available globally.

Make sure `$(go env GOPATH)/bin` is in your `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Verify with:

```bash
news --version
```

## Setup

1. Register for a free API key at [https://newsapi.org/register](https://newsapi.org/register)
2. Set the key using **either** method:

**Option A — `.env` file** (recommended for local development):

Create a `.env` file in the same directory as the binary:

```
NEWS_API_KEY=your_api_key_here
```

**Option B — environment variable**:

```bash
export NEWS_API_KEY=your_api_key_here
```

To persist it across sessions, add the line above to your `~/.bashrc`, `~/.zshrc`, or equivalent shell config file.

## Usage

```
news [keyword] [flags]
```

The keyword is optional. Omitting it returns the current top headlines from around the world.

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--count` | `-n` | `5` | Number of headlines to fetch |
| `--version` | `-v` | — | Print version |
| `--help` | `-h` | — | Show help |

### Examples

Fetch top world headlines (no keyword):

```bash
news
```

Fetch 5 headlines about China:

```bash
news china
```

Fetch 10 headlines about bitcoin:

```bash
news -n 10 bitcoin
```

Fetch 3 headlines about climate change (multi-word keyword):

```bash
news -n 3 "climate change"
```

Fetch 10 general world headlines:

```bash
news -n 10
```

### Sample output

Without a keyword:

```
Top 5 headlines around the world:

1. Nvidia says CEO Jensen Huang is joining Trump's China trip — CNBC
2. Amid war, Trump says he doesn't 'think about Americans' financial situation' — The Washington Post
3. Memphis Grizzlies forward Brandon Clarke dies at 29 — Associated Press
4. Fed holds rates steady amid economic uncertainty — Reuters
5. Scientists discover new deep-sea species off Pacific coast — BBC News
```

With a keyword:

```
Top 5 headlines for "bitcoin":

1. Bitcoin back above $81,000 after hot CPI print — CoinDesk
2. Strait of Hormuz closure boosts Bitcoin, silver appeal — Crypto Briefing
3. Donald Trump to visit China with 16 CEOs — Crypto Briefing
4. Iran war, AI spending could push Bitcoin back to $126K — Cointelegraph
5. XRP Price Finds Support Again — newsBTC
```

Source names are rendered in blue and are clickable hyperlinks in terminals that support OSC 8 (iTerm2, GNOME Terminal, Windows Terminal, Kitty, and most modern terminals).

## API Limits

The free tier of NewsAPI.org allows **100 requests per day**. For higher limits, see [NewsAPI.org pricing](https://newsapi.org/pricing).

> **Note:** NewsAPI.org free tier only returns articles from the last month and restricts sources for non-paying users. Upgrade to a paid plan for real-time results.

## Dependencies

| Package | Purpose |
|---------|---------|
| [spf13/cobra](https://github.com/spf13/cobra) | CLI framework |
| [joho/godotenv](https://github.com/joho/godotenv) | `.env` file loading |

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Commit your changes: `git commit -m 'Add my feature'`
4. Push to the branch: `git push origin feature/my-feature`
5. Open a pull request

## License

Copyright (c) 2026 Mahdi Aryayi — released under the [MIT License](./LICENSE).
