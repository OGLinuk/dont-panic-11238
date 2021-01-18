package main

import (
	"fmt"

	"gitlab.com/OGLinuk/dont-panic-11238/dontpanic"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("checksum")
	x.Usage = `<file>`
	x.Summary = `output crc32 checksum of <file>`
	x.Method = func(args []string) error {
		if len(args) != 1 {
			return x.UsageError()
		}
		checksum, err := dontpanic.Checksum(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("Checksum of %s: %s\n", args[0], checksum)
		return nil
	}
}
