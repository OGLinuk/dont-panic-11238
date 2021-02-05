package main

import "gitlab.com/rwxrob/cmdtab"

func init() {
	x := cmdtab.New("gomain", "fileserver")
	x.Usage = `OPTION`
	x.Summary = `generate a main.go file of type OPTION`
}
