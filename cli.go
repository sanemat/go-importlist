package importlist

import (
	"flag"
	"fmt"
	"io"
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

	argv = fs.Args()
	if len(data) == 0 {
		if len(argv) >= 2 {
			return xerrors.New("We have no subcommand")
		}
		fmt.Fprintf(outStream, "read gofile")
	} else {
		if len(argv) >= 1 {
			return xerrors.New("We have no subcommand")
		}
		fmt.Fprintf(outStream, "read from stdin")
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
