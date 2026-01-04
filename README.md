# AniMagnet

AniMagnet is a lightweight command-line tool for fetching anime-related data and magnet links. It allows users to search for anime entries, display the results, and operate on magnet links for downloading content.

## Features
- **Search for Anime**: Fetch anime entries by providing the anime name.
- **Interact with Results**: Display and copy magnet links directly from the command line.
- **Error Handling**: Clear feedback on invalid inputs or errors during execution.

## Requirements
- [Go](https://golang.org/) version 1.16 or higher

## Installation
Clone this repository into your local machine:

```bash
git clone https://github.com/LeDavidcb/AnimeScraper.git
```

Navigate to the project directory:

```bash
cd AnimeScraper
```

Build the project:

```bash
go build -o animagnet ./cmd/main.go 
```

## Usage
Run the program from the terminal:

```bash
./animagnet --name="<anime_name>"
```

### Flags
- `--name`: The name of the anime to fetch. This flag is required.

## Motivation
AniMagnet was developed as a personal project to learn the Go programming language, so if there are any clear deviations from Go standards, I apologize. This project is also a rewrite of my earlier project, [NyaaScraper](https://github.com/LeDavidcb/nyaa-scraper), which is outdated and no longer functional.
