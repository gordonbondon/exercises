package crawler

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestParser struct {
	pages map[string][]string
}

func (p *TestParser) Run(url string) ([]string, error) {
	if len(p.pages[url]) > 0 && p.pages[url][0] == "error" {
		return nil, fmt.Errorf("Failed loading: %v", errors.New("403"))
	}

	return p.pages[url], nil
}

func TestExecutor(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                 string
		parser               *TestParser
		url                  string
		expectedOptionsError string
		expectedError        string
		expectedPages        []*Page
	}{
		{
			name:                 "wrong url",
			parser:               &TestParser{},
			url:                  "%dsfdsfsdfdsf",
			expectedOptionsError: "Invalid target url provided",
		},
		{
			name: "one page",
			parser: &TestParser{
				pages: map[string][]string{
					"https://example.com": {},
				},
			},
			url:                  "https://example.com",
			expectedOptionsError: "",
			expectedPages: []*Page{
				{Url: "https://example.com", Links: []string{}},
			},
		},
		{
			name: "error page",
			parser: &TestParser{
				pages: map[string][]string{
					"https://example.com": {"error"},
				},
			},
			url:                  "https://example.com",
			expectedOptionsError: "",
			expectedPages: []*Page{
				{Url: "https://example.com", Error: fmt.Errorf("Failed loading: %v", errors.New("403"))},
			},
		},
		{
			name: "recursive pages",
			parser: &TestParser{
				pages: map[string][]string{
					"https://example.com":          {"https://example.com/contacts", "https://example.com/blog"},
					"https://example.com/blog":     {"https://example.com", "https://example.com/contacts"},
					"https://example.com/contacts": {"https://example.com", "https://example.com/blog"},
				},
			},
			url:                  "https://example.com",
			expectedOptionsError: "",
			expectedPages: []*Page{
				{Url: "https://example.com", Links: []string{"https://example.com/contacts", "https://example.com/blog"}},
				{Url: "https://example.com/blog", Links: []string{"https://example.com", "https://example.com/contacts"}},
				{Url: "https://example.com/contacts", Links: []string{"https://example.com", "https://example.com/blog"}},
			},
		},
		{
			name: "trailing slashes or params",
			parser: &TestParser{
				pages: map[string][]string{
					"https://example.com":      {"https://example.com/blog/?utm=foo"},
					"https://example.com/blog": {"https://example.com/", "https://example.com/?utm=foo"},
				},
			},
			url:                  "https://example.com",
			expectedOptionsError: "",
			expectedPages: []*Page{
				{Url: "https://example.com", Links: []string{"https://example.com/blog/?utm=foo"}},
				{Url: "https://example.com/blog", Links: []string{"https://example.com/", "https://example.com/?utm=foo"}},
			},
		},
		{
			name: "outgoing links",
			parser: &TestParser{
				pages: map[string][]string{
					"https://example.com": {"https://community.example.com/blog", "https://facebook.com/example"},
				},
			},
			url:                  "https://example.com",
			expectedOptionsError: "",
			expectedPages: []*Page{
				{Url: "https://example.com", Links: []string{"https://community.example.com/blog", "https://facebook.com/example"}},
			},
		},
		{
			name: "relative links",
			parser: &TestParser{
				pages: map[string][]string{
					"https://example.com":      {"/blog"},
					"https://example.com/blog": {"/"},
				},
			},
			url:                  "https://example.com",
			expectedOptionsError: "",
			expectedPages: []*Page{
				{Url: "https://example.com", Links: []string{"/blog"}},
				{Url: "https://example.com/blog", Links: []string{"/"}},
			},
		},
		{
			name: "more links then channels",
			parser: &TestParser{
				pages: map[string][]string{
					"https://example.com":   {"/1", "/2", "/3", "/4", "/5", "/6", "/7"},
					"https://example.com/1": {},
					"https://example.com/2": {},
					"https://example.com/3": {},
					"https://example.com/4": {},
					"https://example.com/5": {},
					"https://example.com/6": {},
					"https://example.com/7": {},
				},
			},
			url:                  "https://example.com",
			expectedOptionsError: "",
			expectedPages: []*Page{
				{Url: "https://example.com", Links: []string{"/1", "/2", "/3", "/4", "/5", "/6", "/7"}},
				{Url: "https://example.com/1", Links: []string{}},
				{Url: "https://example.com/2", Links: []string{}},
				{Url: "https://example.com/3", Links: []string{}},
				{Url: "https://example.com/4", Links: []string{}},
				{Url: "https://example.com/5", Links: []string{}},
				{Url: "https://example.com/6", Links: []string{}},
				{Url: "https://example.com/7", Links: []string{}},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			options := ExecutorOptions{
				Parser: tc.parser,
				Url:    tc.url,
				// less runners than links to test exhausting channels
				Runners: 2,
			}

			exec, err := NewExecutor(&options)

			if tc.expectedOptionsError == "" {
				assert.Nil(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.expectedOptionsError, "errors not matching")

				return
			}

			pages, err := exec.Run()
			if tc.expectedError == "" {
				assert.Nil(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.expectedError, "errors not matching")
			}

			assert.Equal(t, tc.expectedPages, pages, "wrong pages result")
		})
	}
}
