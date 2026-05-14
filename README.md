# news-cli

A fast command-line tool to fetch top news headlines for any keyword, powered by [NewsAPI.org](https://newsapi.org).

## Features

- Search headlines by any keyword
- Clickable source links in supported terminals
- Configure the number of results with `-n`
- Results sorted by most recent publication date
- API key loaded automatically from a `.env` file

## Requirements

- Go 1.18 or later
- A free [NewsAPI.org](https://newsapi.org/register) API key

## Installation

### Build from source

```bash
git clone https://github.com/mahdiaryayi/news-cli.git
cd news-cli
go build -o news .
```

Optionally move the binary to your PATH:

```bash
mv news /usr/local/bin/news
```

### Install with go install

```bash
go install github.com/mahdiaryayi/news-cli@latest
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

### Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--count` | `-n` | `5` | Number of headlines to fetch |
| `--help` | `-h` | — | Show help |

### Examples

Fetch 5 headlines about China (default):

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

### Sample output

```
Top 5 headlines for "bitcoin":

1. Bitcoin back above $81,000 after hot CPI print — CoinDesk
2. Strait of Hormuz closure boosts Bitcoin, silver appeal — Crypto Briefing
3. Donald Trump to visit China with 16 CEOs — Crypto Briefing
4. Iran war, AI spending could push Bitcoin back to $126K — Cointelegraph
5. XRP Price Finds Support Again — newsBTC
```

Source names at the end of each line are clickable hyperlinks in terminals that support OSC 8 (iTerm2, GNOME Terminal, Windows Terminal, Kitty, and most modern terminals).

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
