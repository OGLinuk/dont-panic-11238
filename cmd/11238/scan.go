package main

import (
	"fmt"

	dontpanic "gitlab.com/OGLinuk/dont-panic-11239"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("scan")
	x.Summary = `Scan local network`
	x.Method = func(args []string) {
		dontpanic.ScanLocalhost()

		for port := range activeLocalhostPorts {
			fmt.Printf("%d is active ...", port)
		}
	}
}
