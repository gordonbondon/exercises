package cmd

import (
	"errors"
	"flag"
	"fmt"

	"github.com/gordonbondon/exercises/takehome/cron_parser1/pkg/cron"
)

const (
	cronFlagHelp = `Provide cron string in crontab format: M H DM M DW CMD`
)

// Run is main entrypoint for cron_parser1 CLI
func Run() error {
	var cronFlag string

	flag.StringVar(&cronFlag, "crontab", "", cronFlagHelp)
	flag.Parse()

	if cronFlag == "" {
		return errors.New("cron parmeter is required")
	}

	parserOpts := &cron.ParserOptions{}

	parser, err := cron.NewParser(parserOpts)
	if err != nil {
		return fmt.Errorf("Failed configuring cron parser: %w", err)
	}

	schedule, err := parser.Run(cronFlag)
	if err != nil {
		return fmt.Errorf("Failed parsing crontab: %w", err)
	}

	printerOpts := &cron.PrinterOptions{}

	printer, err := cron.NewPrinter(printerOpts)
	if err != nil {
		return fmt.Errorf("Failed configuring stdout printer: %w", err)
	}

	err = printer.Run(schedule)
	if err != nil {
		return fmt.Errorf("Failed printing crontab: %w", err)
	}

	return nil
}
