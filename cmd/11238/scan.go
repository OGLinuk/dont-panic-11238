package main

import (
	"log"

	"gitlab.com/OGLinuk/dont-panic-11238/dontpanic"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("scan")
	x.Summary = `Scan localhost ports and print the active ones`
	x.Method = func(args []string) error {
		log.Printf("Scanning localhost ...")
		for _, port := range dontpanic.ScanLocalhost() {
			log.Printf("%s is active ...\n", port)
		}
		return nil
	}
}
