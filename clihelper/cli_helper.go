package clihelper

import (
	"errors"
	"flag"
	"fmt"
	"github.com/xtforgame/restfs/fshelper"
	"os"
)

type CliHelper struct {
	h bool
	v bool

	args        []string
	storagePath string
}

func (ch *CliHelper) H() bool             { return ch.h }
func (ch *CliHelper) V() bool             { return ch.v }
func (ch *CliHelper) Args() []string      { return ch.args }
func (ch *CliHelper) StoragePath() string { return ch.storagePath }

func (ch *CliHelper) SetFlag() {
	flag.BoolVar(&ch.h, "h", false, "this help")
	flag.BoolVar(&ch.v, "v", false, "show version and exit")
	flag.Usage = Usage
}

func (ch *CliHelper) Usage() { flag.Usage() }

func (ch *CliHelper) Parse() {
	flag.Parse()
	ch.args = flag.Args()
}

func (ch *CliHelper) Validate() (int, error) {
	if ch.h || ch.v {
		return 0, nil
	}

	if len(ch.args) < 1 {
		return 2, errors.New("command error: missing <storage-path>")
	}
	var err error
	ch.storagePath, err = fshelper.NormalizePath(ch.args[0])
	if err != nil {
		return 2, errors.New("command error: invalid <storage-path> :" + ch.args[0])
	}

	return 0, nil
}

func CreateCliHelper() *CliHelper {
	return &CliHelper{
		args: []string{},
	}
}

func Usage() {
	flag.CommandLine.SetOutput(os.Stdout)
	fmt.Fprintf(flag.CommandLine.Output(), `bptd version: bptd/1.0.0
Usage: bptd [-hv] <storage-path>

Options:
`)
	flag.PrintDefaults()
}
