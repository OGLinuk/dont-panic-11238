package main

import (
	"log"

	"gitlab.com/OGLinuk/dont-panic-11238/dontpanic"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("fileserver")
	x.Usage = `<name> <port> <link>`
	x.Summary = `generate a main.go file that serves a fileserver`
	x.Method = func(args []string) error {
		if len(args) != 3 {
			return x.UsageError()
		}
		if err := dontpanic.GenerateFileServer(args[0], args[1], args[2]); err != nil {
			return err
		}
		log.Printf("Successfully generated main.go at %s ...", args[2])
		return nil
	}
}
