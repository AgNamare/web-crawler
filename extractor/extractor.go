package extractor

import (
	"io"

	"golang.org/x/net/html"
)

func Extract_links(html_body io.Reader) (links []string, err error) {
	var extracted_links []string
	tokenizer := html.NewTokenizer(html_body)
	for {
		token_type := tokenizer.Next()
		if token_type == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				break
			} else {
				return extracted_links, tokenizer.Err()
			}
		}
		token := tokenizer.Token()
		data := token.Data
		if token_type == html.StartTagToken {
			if data == "a" {
				attributes := token.Attr
				for _, attr := range attributes {
					if attr.Key == "href" {
						extracted_links = append(extracted_links, attr.Val)
					}
				}
			}
		}
	}
	return extracted_links, nil
}
