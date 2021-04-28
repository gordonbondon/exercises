package html

import (
	"io"
	"net/http"
)

// HttpEngine implements html.Engine interface
type HttpEngine struct {
	Options *HttpEngineOptions
}

type HttpEngineOptions struct {
	HttpEngine HttpEngine
}

func NewHttpEngine(options *HttpEngineOptions) (*HttpEngine, error) {
	engine := &HttpEngine{}

	engine.Options = options

	return engine, nil
}

func (p *HttpEngine) Run(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	output, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return output, nil
}
