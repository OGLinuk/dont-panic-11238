package main

import (
	"fmt"

	"gitlab.com/OGLinuk/dont-panic-11238/dontpanic"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("checkexists")
	x.Usage = `<path>`
	x.Summary = `checks if <path> exists`
	x.Method = func(args []string) error {
		if len(args) != 1 {
			return x.UsageError()
		}
		does := "DOES"
		if !dontpanic.CheckExists(args[0]) {
			does += " NOT"
		}
		fmt.Printf("%s %s exist ...\n", args[0], does)
		return nil
	}
}
