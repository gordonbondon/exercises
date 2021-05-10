package cron

import (
	"fmt"
	"strings"
)

type Printer struct {
	Options *PrinterOptions
}

type PrinterOptions struct{}

func NewPrinter(options *PrinterOptions) (*Printer, error) {
	engine := &Printer{}

	engine.Options = options

	return engine, nil
}

func (p *Printer) Run(s *Schedule) error {
	var output strings.Builder

	output.WriteString(fmt.Sprintf("minute \t\t%v\n", s.Minutes))
	output.WriteString(fmt.Sprintf("hour \t\t%v\n", s.Hours))
	output.WriteString(fmt.Sprintf("day of month \t%v\n", s.DaysOfMonth))
	output.WriteString(fmt.Sprintf("month \t\t%v\n", s.Months))
	output.WriteString(fmt.Sprintf("day of week \t%v\n", s.DaysOfWeek))
	output.WriteString(fmt.Sprintf("command \t%v\n", s.Command))

	fmt.Println(output.String())

	return nil
}
