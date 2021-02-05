package main

import (
	"log"

	dontpanic "gitlab.com/OGLinuk/dont-panic-11238"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("ip")
	x.Summary = `Get the host machines IPv4 address`
	x.Method = func(args []string) error {
		log.Printf("Host machines IP: %s", dontpanic.GetLocalIP())
		return nil
	}
}
