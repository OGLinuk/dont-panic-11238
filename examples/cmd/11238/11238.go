package main

import (
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("11238", "scan", "ip", "generate", "checkexists")
	x.Default = "scan"
	x.Summary = `utility to interact with dont-panic-11238`
	x.Version = "1.0.0"
	x.Author = "oglinuk"
	x.Git = "gitlab.com/OGLinuk/dont-panic-11238"
}
