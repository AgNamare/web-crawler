package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"web-crawler/extractor"
)

type Options struct {
	MaxDepth int // 0 = unlimited
}

type item struct {
	url   string
	depth int
}

func getLinks(pageURL string) []string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return nil
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var links []string
	for r := range extractor.Extract_links(pageURL, resp.Body) {
		if r.Err != nil {
			break
		}
		links = append(links, r.Link)
	}
	return links
}

func Crawl(start string, opts Options) map[string]bool {
	visited := make(map[string]bool)
	queue := []item{{start, 0}}
	base, _ := url.Parse(start)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if visited[curr.url] {
			continue
		}
		visited[curr.url] = true
		fmt.Println("crawling:", curr.url)

		if opts.MaxDepth > 0 && curr.depth >= opts.MaxDepth {
			continue
		}

		for _, link := range getLinks(curr.url) {
			u, err := url.Parse(link)
			if err != nil || u.Host != base.Host {
				continue
			}
			if !visited[link] {
				queue = append(queue, item{link, curr.depth + 1})
			}
		}
	}

	return visited
}
