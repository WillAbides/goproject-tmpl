package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var version string

const description = `` // add description here

type cliRoot struct {
	Version kong.VersionFlag `kong:"help=${VersionHelp}"`
}

var kongVars = kong.Vars{
	"version":     version,
	"VersionHelp": `Output the goproject-tmpl version and exit.`,
}

func main() {
	var cli cliRoot
	kctx := kong.Parse(&cli,
		kongVars,
		kong.Description(description),
	)

	fmt.Fprintln(kctx.Stdout, "hello, world")
}
