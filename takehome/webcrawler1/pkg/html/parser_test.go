package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestEngine struct {
	pages map[string]string
}

func (p *TestEngine) Run(url string) ([]byte, error) {
	return []byte(p.pages[url]), nil
}

func TestParser(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		engine               *TestEngine
		url                  string
		expectedOptionsError string
		expectedError        string
		expectedLinks        []string
	}{
		{
			name: "simple page",
			engine: &TestEngine{
				pages: map[string]string{
					"https://example.com": `
						<!DOCTYPE html>
						<html class="no-js " lang="en" dir="ltr">
						<head>
							<link rel="canonical" href="https://example.com/">
						</head>
						<body>
							<a href="/blog/" class="main-navigation__links__link">Blog</a>
						</body>
					</html>
					`,
				},
			},
			url:           "https://example.com",
			expectedLinks: []string{"/blog/"},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			options := ParserOptions{
				Engine: tc.engine,
			}

			parser, err := NewParser(&options)

			if tc.expectedOptionsError == "" {
				assert.Nil(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.expectedOptionsError, "errors not matching")

				return
			}

			links, err := parser.Run(tc.url)
			if tc.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "errors not matching")
			}

			assert.Equal(t, tc.expectedLinks, links, "wrong pages result")

		})
	}
}
