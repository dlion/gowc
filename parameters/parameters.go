package parameters

import (
	"flag"
	"os"
)

const (
	BytesFlag = "c"
	LinesFlag = "l"
	WordsFlag = "w"
	CharsFlag = "m"
)

func HasProvided() bool {
	return len(os.Args) > 1
}

func GetFilename() string {
	return os.Args[len(os.Args)-1]
}

func GetFlags() map[string]*bool {
	flags := map[string]*bool{
		BytesFlag: flag.Bool(BytesFlag, false, "Count bytes of the file"),
		LinesFlag: flag.Bool(LinesFlag, false, "Count lines of the file"),
		WordsFlag: flag.Bool(WordsFlag, false, "Count words of the file"),
		CharsFlag: flag.Bool(CharsFlag, false, "Count chars of the file"),
	}

	flag.Parse()

	return flags
}

func HaveBeenPassed(flags map[string]*bool) (string, bool) {
	for flagName, flagValue := range flags {
		if *flagValue == true {
			return flagName, true
		}
	}
	return "", false
}
