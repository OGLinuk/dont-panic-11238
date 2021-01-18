package main

import (
	"log"

	"gitlab.com/OGLinuk/dont-panic-11238/dontpanic"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("docker")
	x.Usage = `<name> <port> <path>`
	x.Summary = `generate a Dockerfile using name, port, and path`
	x.Method = func(args []string) error {
		if len(args) != 3 {
			return x.UsageError()
		}
		if err := dontpanic.GenerateDockerfile(args[0], args[1], args[2]); err != nil {
			return err
		}
		log.Printf("Successfully generated Dockerfile at %s ...", args[2])
		return nil
	}
}
