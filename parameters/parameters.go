package parameters

import "flag"

const (
	BytesFlag = "c"
	LinesFlag = "l"
	WordsFlag = "w"
	CharsFlag = "m"
)

type Parameters struct {
	defParams      map[string]string
	noFlagProvided bool
}

func NewParameters(defParams map[string]string) Parameters {
	return Parameters{defParams: defParams}
}

func (p *Parameters) Parse() {
	for paramName, usage := range p.defParams {
		flag.String(paramName, "", usage)
	}

	if flag.Parse(); len(flag.Args()) != 0 {
		p.noFlagProvided = true
	}
}

func (p *Parameters) NoFlagsProvided() bool {
	return p.noFlagProvided
}

func (p *Parameters) CountBytes() (bool, string) {
	if p.noFlagProvided {
		return true, flag.Args()[0]
	}

	lookup := flag.Lookup(BytesFlag)
	return getParameterValueFromFlag(lookup)
}

func (p *Parameters) CountLines() (bool, string) {
	if p.noFlagProvided {
		return true, flag.Args()[0]
	}

	lookup := flag.Lookup(LinesFlag)
	return getParameterValueFromFlag(lookup)
}

func (p *Parameters) CountWords() (bool, string) {
	if p.noFlagProvided {
		return true, flag.Args()[0]
	}

	lookup := flag.Lookup(WordsFlag)
	return getParameterValueFromFlag(lookup)
}

func (p *Parameters) CountChars() (bool, string) {
	lookup := flag.Lookup(CharsFlag)
	return getParameterValueFromFlag(lookup)
}

func getParameterValueFromFlag(lookup *flag.Flag) (bool, string) {
	if lookup == nil || lookup.Value.String() == "" {
		return false, ""
	}
	return true, lookup.Value.String()
}
