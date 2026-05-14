# news-cli

A fast command-line tool to fetch top news headlines for any keyword, powered by [NewsAPI.org](https://newsapi.org).

## Features

- Search headlines by any keyword
- Configure the number of results with `-n`
- Results sorted by most recent publication date
- Clean, readable terminal output

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

Fetch 1 headline about climate:

```bash
news -n 1 "climate change"
```

### Sample output

```
Top 5 headlines for "china":

1. [Reuters] China vows retaliation after new US tariffs take effect
   https://reuters.com/...

2. [BBC News] China's economy grows faster than expected in Q1
   https://bbc.com/...

3. [Bloomberg] China tech stocks rally on stimulus hopes
   https://bloomberg.com/...

4. [AP News] China and Taiwan tensions rise over military drills
   https://apnews.com/...

5. [The Guardian] China pledges to hit net-zero emissions by 2060
   https://theguardian.com/...
```

## API Limits

The free tier of NewsAPI.org allows **100 requests per day**. For higher limits, see [NewsAPI.org pricing](https://newsapi.org/pricing).

## License

MIT
