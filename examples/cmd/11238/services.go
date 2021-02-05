package main

import (
	"log"

	dontpanic "gitlab.com/OGLinuk/dont-panic-11238"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("services")
	x.Summary = `generate services directory and git clone/pull services defined in manifest files`
	x.Method = func(args []string) error {
		if err := dontpanic.GenerateServices(); err != nil {
			return err
		}
		log.Printf("Successfully generated services ...")
		return nil
	}
}
