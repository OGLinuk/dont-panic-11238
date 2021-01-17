package main

import (
	"fmt"

	"gitlab.com/OGLinuk/dont-panic-11238/dontpanic"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("scan")
	x.Summary = `Scan local network`
	x.Method = func(args []string) error {
		for _, port := range dontpanic.ScanLocalhost() {
			fmt.Printf("%s is active ...\n", port)
		}
		return nil
	}
}
