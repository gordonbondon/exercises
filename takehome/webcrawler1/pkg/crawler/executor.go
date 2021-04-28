package crawler

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type Parser interface {
	Run(string) ([]string, error)
}

type Printer interface {
	Run(*Page)
}

type Page struct {
	// Url of a page
	Url string
	// Links list of url parsed from page html
	Links []string
	// Error represents error encountered while trying to parse the page
	Error error
}

type Executor struct {
	// Options allows setting executor configuration
	Options *ExecutorOptions

	host   string
	scheme string

	unprocessed chan string
	queue       chan *Page
}

type ExecutorOptions struct {
	// Url to website we need to crawl
	Url string

	// Parser engine that will exctract urls from pages
	Parser Parser

	// Printer allows to configure execution outout
	Printer Printer

	// Runners set number of concurrent workers
	Runners int
}

func NewExecutor(options *ExecutorOptions) (*Executor, error) {
	executor := &Executor{}

	if options.Parser == nil {
		return nil, fmt.Errorf("Parser needs to be set")
	}

	if options.Runners == 0 {
		options.Runners = 1
	}

	entrypoint, err := url.Parse(options.Url)
	if err != nil {
		return nil, fmt.Errorf("Invalid target url provided: %w", err)
	}

	executor.Options = options

	executor.host = entrypoint.Host
	executor.scheme = entrypoint.Scheme

	return executor, nil
}

func (e *Executor) Run() ([]*Page, error) {
	output := make([]*Page, 0)
	e.unprocessed = make(chan string)
	e.queue = make(chan *Page)

	visited := make(map[string]bool)
	counter := 0

	for w := 1; w <= e.Options.Runners; w++ {
		go e.process()
	}

	counter++
	e.unprocessed <- e.Options.Url

	for counter > 0 {
		select {
		case page := <-e.queue:
			if _, ok := visited[page.Url]; ok {
				counter--

				continue
			}

			// mark visited page
			visited[page.Url] = true

			output = append(output, page)

			if e.Options.Printer != nil {
				e.Options.Printer.Run(page)
			}

			if page.Error != nil {
				counter--

				continue
			}

			for _, l := range page.Links {
				n, err := e.normalizeLink(l)
				if err != nil {
					continue
				}

				if n.Host != e.host {
					continue
				}

				if _, ok := visited[n.String()]; ok {
					continue
				}

				counter++

				e.unprocessed <- n.String()
			}

			counter--
		}

		if counter == 0 {
			close(e.unprocessed)

			break
		}
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Url < output[j].Url
	})

	return output, nil
}

func (e *Executor) process() {
	for url := range e.unprocessed {
		result, err := e.Options.Parser.Run(url)
		if err != nil {
			go e.enqueue(&Page{
				Url:   url,
				Error: err,
			})

			continue
		}

		go e.enqueue(&Page{
			Url:   url,
			Links: result,
		})
	}
}

func (e *Executor) enqueue(page *Page) {
	e.queue <- page
}

func (e *Executor) normalizeLink(link string) (*url.URL, error) {
	url, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	url.Path = strings.TrimSuffix(url.Path, "/")
	url.Fragment = ""
	url.RawQuery = ""

	if url.Host == "" {
		url.Host = e.host
	}

	if url.Scheme == "" {
		url.Scheme = e.scheme
	}

	return url, nil
}
