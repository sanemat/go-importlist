package importlist

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"

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
		fmt.Fprint(outStream, list)
	} else { // stdin
		if len(argv) >= 1 {
			return xerrors.New("We have no subcommand")
		}
		list, err2 = importList(data)
		if err2 != nil {
			return err2
		}
		fmt.Fprint(outStream, list)
	}
	if *nullTerminators {
		fmt.Fprintf(outStream, "null terminator")
	}
	return nil
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s v%s (rev:%s)\n", cmdName, version, revision)
	return err
}

func importList(data []byte) ([]string, error) {
	var result []string
	return result, nil
}
