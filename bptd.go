package main

import (
	"fmt"
	"github.com/xtforgame/restfs/clihelper"
	"github.com/xtforgame/restfs/httpserver"
	"log"
	"os"
	"runtime"
	// "time"
)

// FinalReport is the panic handler and provides the last report
func FinalReport() {
	if err := recover(); err != nil {
		trace := make([]byte, 1024)
		count := runtime.Stack(trace, true)
		errMsg := fmt.Sprintf("%s", err)
		fmt.Printf("Recover from panic: %s\n", errMsg)
		fmt.Printf("Stack of %d bytes: %s\n", count, string(trace[:count]))
		os.Exit(2)
	} else {
	}
}

var cliHelper = clihelper.CreateCliHelper()

func init() {
	cliHelper.SetFlag()
}

func main() {
	cliHelper.Parse()

	if cliHelper.H() {
		cliHelper.Usage()
		return
	}

	if cliHelper.V() {
		fmt.Println("restfs v1.0.0")
		return
	}

	exitCode, err := cliHelper.Validate()
	if err != nil {
		l := log.New(os.Stderr, "", 0)
		l.Println(err)
		os.Exit(exitCode)
	}

	defer FinalReport()
	// os.Exit(0)

	hs := httpserver.NewHttpServer()
	hs.Init()
	hs.Start()
}
