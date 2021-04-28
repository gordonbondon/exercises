package crawler

import (
	"fmt"
	"strings"
)

// StdoutPrinter implements crawler.Printer
type StdoutPrinter struct {
	// Options allows setting printer configuration
	Options *StdoutPrinterOptions
}

type StdoutPrinterOptions struct{}

func NewStdoutPrinter(options *StdoutPrinterOptions) (*StdoutPrinter, error) {
	printer := &StdoutPrinter{}

	return printer, nil
}

func (p *StdoutPrinter) Run(page *Page) {
	var output strings.Builder

	output.WriteString(fmt.Sprintf("Page: %s\n", page.Url))

	if page.Error != nil {
		output.WriteString(fmt.Sprintf("\tLoading failed with error: %v\n", page.Error))
	}

	if len(page.Links) > 0 {
		output.WriteString("\tLinks:\n")

		for _, l := range page.Links {
			output.WriteString(fmt.Sprintf("\t\t%s\n", l))
		}
	}

	fmt.Println(output.String())
}
