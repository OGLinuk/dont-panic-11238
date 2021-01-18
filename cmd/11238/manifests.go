package main

import (
	"log"

	"gitlab.com/OGLinuk/dont-panic-11238/dontpanic"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("manifests")
	x.Usage = `<environment>`
	x.Summary = `generate manifests directory and default manifests of type`
	x.Method = func(args []string) error {
		if len(args) != 1 {
			return x.UsageError()
		}
		if err := dontpanic.GenerateManifests(args[0]); err != nil {
			return err
		}
		log.Printf("Successfully generated %s manifests ...", args[0])
		return nil
	}
}
