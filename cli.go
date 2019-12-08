package importlist

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/xerrors"
)

const cmdName = "import-list"

// Run command
func Run(argv []string, data []byte, outStream, errStream io.Writer) error {
	log.SetOutput(errStream)
	log.SetPrefix(fmt.Sprintf("[%s] ", cmdName))
	nameAndVer := fmt.Sprintf("%s (v%s rev:%s)", cmdName, version, revision)
	fs := flag.NewFlagSet(nameAndVer, flag.ContinueOnError)
	fs.SetOutput(errStream)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage of %s:\n", nameAndVer)
		fs.PrintDefaults()
	}

	var ver = fs.Bool("version", false, "display version")
	var nullTerminators = fs.Bool("z", false, "use NULs as output field terminators")

	if err := fs.Parse(argv); err != nil {
		return err
	}
	if *ver {
		return printVersion(outStream)
	}

	var list []string
	var err2 error
	argv = fs.Args()
	if len(data) == 0 { // read file
		if len(argv) != 1 {
			return xerrors.New("require one target golang file")
		}
		data2, err := ioutil.ReadFile(argv[0])
		if err != nil {
			return err
		}
		list, err2 = importList(data2)
		if err2 != nil {
			return err2
		}
	} else { // stdin
		if len(argv) >= 1 {
			return xerrors.New("We have no subcommand")
		}
		list, err2 = importList(data)
		if err2 != nil {
			return err2
		}
	}

	if *nullTerminators {
		fmt.Fprint(outStream, strings.Join(list, "\x00"))
	} else {
		last := len(list) - 1
		for i, r := range list {
			if i == last {
				fmt.Fprint(outStream, r)
			} else {
				fmt.Fprintln(outStream, r)
			}
		}
	}
	return nil
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s v%s (rev:%s)\n", cmdName, version, revision)
	return err
}

func importList(data []byte) ([]string, error) {
	var result []string
	result = []string{"a", "b"}
	return result, nil
}
