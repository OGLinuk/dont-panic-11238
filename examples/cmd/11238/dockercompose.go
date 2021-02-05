package main

import (
	"log"

	dontpanic "gitlab.com/OGLinuk/dont-panic-11238"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("dockercompose")
	x.Summary = `generate  docker-compose.yml file from manifest files`
	x.Method = func(args []string) error {
		if err := dontpanic.GenerateDockerCompose(); err != nil {
			return err
		}
		log.Printf("Successfully generated docker-compose.yml ...")
		return nil
	}
}
