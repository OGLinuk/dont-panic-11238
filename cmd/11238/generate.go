package main

import "gitlab.com/rwxrob/cmdtab"

func init() {
	x := cmdtab.New("generate", "manifests", "services", "docker", "dockercompose", "gomain", "checksum")
	x.Usage = `OPTION`
	x.Summary = `generate file(s)`
}
