# web-crawler

A simple BFS web crawler written in Go. Give it a URL and it'll crawl every page it can find on that domain. You can also cap how deep it goes with a depth flag.

## Requirements

- Go 1.21+

## Getting started

Clone the repo and grab dependencies:

```bash
git clone <your-repo-url>
cd web-crawler
go mod tidy
```

## Running it

```bash
go run . <url>
```

For example:

```bash
go run . https://books.toscrape.com/
```

This will crawl the entire site and print every URL it visits. When it's done it'll tell you how many pages it found.

### Limiting depth

If you don't want it crawling the whole site, pass `--depth`:

```bash
go run . --depth 2 https://books.toscrape.com/
```

Depth 0 (the default) means no limit — it'll go until there's nothing left to visit.

### Building a binary

```bash
go build -o crawler .
./crawler --depth 3 https://books.toscrape.com/
```

## Using it as a package

You can import the crawler into your own project:

```go
import "web-crawler/crawler"

pages := crawler.Crawl("https://example.com", crawler.Options{
    MaxDepth: 2,
})

fmt.Println("pages found:", len(pages))
```

`Crawl` returns a `map[string]bool` of every URL it visited. `MaxDepth: 0` means unlimited.

## Project structure

```
web-crawler/
├── main.go           # CLI entry point
├── crawler/
│   └── crawler.go    # crawl logic, importable as a package
└── extractor/
    └── extractor.go  # pulls links out of an HTML page
```

## Notes

- Only crawls pages on the same domain as the starting URL
- Handles relative and absolute URLs correctly
- Won't visit the same page twice
