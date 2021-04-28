package cmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/gordonbondon/exercises/takehome/webcrawler1/pkg/crawler"
	"github.com/gordonbondon/exercises/takehome/webcrawler1/pkg/html"
)

const (
	urlFlagHelp      = `Provide website url to crawl`
	parallelFlagHelp = `Set number of parallel crawler workers`
)

// Run is main entrypoint for webcrawler1 CLI
func Run() error {
	var urlFlag string
	var parallelFlag int

	flag.StringVar(&urlFlag, "url", "", urlFlagHelp)
	flag.IntVar(&parallelFlag, "parallelism", 10, parallelFlagHelp)
	flag.Parse()

	if urlFlag == "" {
		return errors.New("url required")
	}

	engineOpts := &html.HttpEngineOptions{}

	engine, err := html.NewHttpEngine(engineOpts)
	if err != nil {
		return fmt.Errorf("Failed configuring HTML engine: %w", err)
	}

	parserOpts := &html.ParserOptions{
		Engine: engine,
	}

	parser, err := html.NewParser(parserOpts)
	if err != nil {
		return fmt.Errorf("Failed configuring HTML parser: %w", err)
	}

	printOpts := &crawler.StdoutPrinterOptions{}

	printer, err := crawler.NewStdoutPrinter(printOpts)
	if err != nil {
		return fmt.Errorf("Failed configuring output: %w", err)
	}

	execOpts := &crawler.ExecutorOptions{
		Parser:  parser,
		Printer: printer,
		Url:     urlFlag,
		Runners: parallelFlag,
	}

	executor, err := crawler.NewExecutor(execOpts)
	if err != nil {
		return fmt.Errorf("Failed configuring crawler: %w", err)
	}

	_, err = executor.Run()
	if err != nil {
		return fmt.Errorf("Failed running crawler: %w", err)
	}

	return nil
}
