package main

import (
	"os"
	"strings"

	checkprocs "github.com/mackerelio/go-check-plugins/check-procs/lib"
)

func main() {
	ckr := checkprocs.Run(os.Args[1:])
	ckr.Name = ""
	if strings.Contains(ckr.Message, "Failure to parse flags") {
		os.Exit(1)
	}
	ckr.Exit()
}
