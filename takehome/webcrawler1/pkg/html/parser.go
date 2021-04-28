package html

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Engine interface {
	Run(string) ([]byte, error)
}

// Parser implements crawler.Parser interface
type Parser struct {
	Options *ParserOptions
}

type ParserOptions struct {
	Engine Engine
}

func NewParser(options *ParserOptions) (*Parser, error) {
	parser := &Parser{}

	if options.Engine == nil {
		return nil, fmt.Errorf("Engine needs to be set")
	}

	parser.Options = options

	return parser, nil
}

func (p *Parser) Run(url string) ([]string, error) {
	output := make([]string, 0)

	content, err := p.Options.Engine.Run(url)
	if err != nil {
		return output, fmt.Errorf("Failed querying link: %w", err)
	}

	z := html.NewTokenizer(strings.NewReader(string(content)))

	for {
		tt := z.Next()

		if tt == html.ErrorToken {
			break
		}

		switch tt {
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						output = append(output, attr.Val)
					}
				}
			}
		}
	}

	return output, nil
}
