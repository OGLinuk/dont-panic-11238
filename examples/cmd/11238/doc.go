package main

import (
	"log"

	dontpanic "gitlab.com/OGLinuk/dont-panic-11238"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("doc")
	x.Usage = `<path>`
	x.Summary = `generate an HTML file from a markdown file`
	x.Method = func(args []string) error {
		if len(args) != 1 {
			return x.UsageError()
		}
		if err := dontpanic.GenerateDoc(args[0]); err != nil {
			return err
		}
		log.Printf("Successfully rendered %s ...", args[0])
		return nil
	}
}
