package cron

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
)

type Parser struct {
	Options *ParserOptions
}

type ParserOptions struct{}

type field int

type fieldSpec struct {
	min int
	max int
	// TODO: extend to support named ranges and singular names
}

const (
	minute field = iota
	hour
	dayOfMonth
	month
	dayOfWeek
)

var (
	fieldTypes = map[field]fieldSpec{
		minute:     {min: 0, max: 59},
		hour:       {min: 0, max: 23},
		dayOfMonth: {min: 1, max: 31},
		month:      {min: 1, max: 12},
		dayOfWeek:  {min: 0, max: 7},
	}
)

func NewParser(options *ParserOptions) (*Parser, error) {
	engine := &Parser{}

	engine.Options = options

	return engine, nil
}

func (p *Parser) Run(crontab string) (*Schedule, error) {
	parts := strings.Split(crontab, " ")

	if len(parts) < 6 {
		return nil, fmt.Errorf("Wrong format of crontab: should be `M H DM M DW CMD`, got: %s", crontab)
	}

	var parseError error
	parsed := make([][]int, 5)

	for f, s := range fieldTypes {
		r, err := parseField(parts[f], s)
		if err != nil {
			parseError = multierror.Append(parseError, err)
		}

		parsed[f] = r
	}

	if parseError != nil {
		return nil, fmt.Errorf("Wrong format of crontab: %w", parseError)
	}

	result := &Schedule{
		Minutes:     parsed[0],
		Hours:       parsed[1],
		DaysOfMonth: parsed[2],
		Months:      parsed[3],
		DaysOfWeek:  parsed[4],
		Command:     parts[5],
	}

	return result, nil
}

func parseField(field string, spec fieldSpec) ([]int, error) {
	result := make([]int, 0)
	resultCheck := make(map[int]bool)

	// check for "," separated list of spans
	list := strings.Split(field, ",")

	// parse each span
	for _, span := range list {
		var err error
		var min int
		var max int

		step := 1

		// check for "/" separated step value
		spanIter := strings.Split(span, "/")

		if len(spanIter) > 2 {
			return nil, fmt.Errorf("expected up to one step, got: %s", span)
		} else if len(spanIter) > 1 {
			if spanIter[1] != "" {
				step, err = strconv.Atoi(spanIter[1])
				if err != nil {
					return nil, fmt.Errorf("expected step as a number, got: %s", span)
				}
			} else {
				return nil, fmt.Errorf("expected step provided, got empty: %s", span)
			}
		}

		// check for "-" separated range
		spanRange := strings.Split(spanIter[0], "-")

		if len(spanRange) > 2 {
			return nil, fmt.Errorf("expected one number or range, got: %s", span)
		} else if len(spanRange) > 1 {
			min, err = strconv.Atoi(spanRange[0])
			if err != nil {
				return nil, fmt.Errorf("expected range from numbers, got: %s", span)
			}

			if min < spec.min {
				return nil, fmt.Errorf("expected minimum of %d, got: %s", spec.min, span)
			}

			max, err = strconv.Atoi(spanRange[1])
			if err != nil {
				return nil, fmt.Errorf("expected range from numbers, got: %s", span)
			}

			if max > spec.max {
				return nil, fmt.Errorf("expected maximum of %d, got: %s", spec.max, span)
			}
		} else if spanRange[0] != "*" {
			oneNumber, err := strconv.Atoi(spanRange[0])
			if err != nil {
				return nil, fmt.Errorf("expected number, got: %s", span)
			}

			if oneNumber < spec.min || oneNumber > spec.max {
				return nil, fmt.Errorf("expected number in range of %d to %d, got: %s", spec.min, spec.max, span)
			}

			if _, ok := resultCheck[oneNumber]; ok {
				continue
			}

			result = append(result, oneNumber)

			resultCheck[oneNumber] = true

			continue
		}

		// check for glob
		if spanIter[0] == "*" {
			min = spec.min
			max = spec.max
		}

		for i := min; i <= max; i = i + step {
			if _, ok := resultCheck[i]; ok {
				continue
			}

			result = append(result, i)

			resultCheck[i] = true
		}
	}

	sort.Ints(result)

	return result, nil
}
