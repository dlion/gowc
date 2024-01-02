package parser

import (
	"flag"
	"os"
)

type Parser struct {
	flagSet *flag.FlagSet
}

func NewParser(flagSet *flag.FlagSet) *Parser {
	return &Parser{flagSet}
}

func (p Parser) GetParameters() (map[string]string, error) {
	err := p.flagSet.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	parametersMap := make(map[string]string)
	p.flagSet.Visit(func(f *flag.Flag) {
		parametersMap[f.Name] = f.Value.String()
	})

	return parametersMap, nil
}
