package parameters

import (
	"flag"
	"fmt"
	"gowc/parser"
	"gowc/utils"
	"io/fs"
	"os"
)

func GetParameters() map[string]string {
	f := flag.NewFlagSet("gowc parameters", flag.ExitOnError)
	setParameterDefinitions(f)

	p := parser.NewParser(f)

	parameters, err := p.GetParameters()
	if err != nil {
		fmt.Printf("Can't read the parameters correctly\n")
		os.Exit(1)
	}
	return parameters
}

func setParameterDefinitions(f *flag.FlagSet) {
	f.String("c", "", "Path to the file you want to count the bytes")
	f.String("l", "", "Path to the file you want to count the lines")
	f.String("w", "", "Path to the file you want to count the words")
}

func GetFSandFilenameFromParameter(path string) (fs.FS, string) {
	dirPath, fileName := utils.SplitPath(path)
	return os.DirFS(dirPath), fileName
}
