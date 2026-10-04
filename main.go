package main

import (
	"flag"
	"fmt"
	"os"
	"web-crawler/crawler"
)

func main() {
	depth := flag.Int("depth", 0, "max crawl depth (0 = unlimited)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: crawler [--depth N] <url>")
		os.Exit(1)
	}

	url := flag.Arg(0)
	pages := crawler.Crawl(url, crawler.Options{MaxDepth: *depth})
	fmt.Println("total pages crawled:", len(pages))
}
