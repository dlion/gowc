package parser

import (
	"flag"
	"os"
)

type Parser struct {
	fs *flag.FlagSet
}

func NewParser(fs *flag.FlagSet) *Parser {
	return &Parser{fs}
}

func (p Parser) GetParameters() (map[string]string, error) {
	err := p.fs.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	parametersMap := make(map[string]string)
	p.fs.Visit(func(f *flag.Flag) {
		parametersMap[f.Name] = f.Value.String()
	})

	return parametersMap, nil
}
