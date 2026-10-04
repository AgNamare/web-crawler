package extractor

import (
	"io"
	"net/url"

	"golang.org/x/net/html"
)

type LinkResult struct {
	Link string
	Err  error
}

func Extract_links(base string, html_body io.Reader) <-chan LinkResult {
	ch := make(chan LinkResult)
	tokenizer := html.NewTokenizer(html_body)
	baseURL, _ := url.Parse(base)

	go func() {
		defer close(ch)
		for {
			tt := tokenizer.Next()
			if tt == html.ErrorToken {
				ch <- LinkResult{Err: tokenizer.Err()}
				return
			}
			if tt == html.StartTagToken {
				t := tokenizer.Token()
				if t.Data == "a" {
					for _, attr := range t.Attr {
						if attr.Key == "href" {
							link, err := url.Parse(attr.Val)
							if err != nil {
								continue
							}
							resolved := baseURL.ResolveReference(link)
							if resolved.Scheme == "http" || resolved.Scheme == "https" {
								ch <- LinkResult{Link: resolved.String()}
							}
						}
					}
				}
			}
		}
	}()
	return ch
}
