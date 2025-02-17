package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var version = "unknown"

const description = `` // add description here

type cmdRoot struct {
	Version kong.VersionFlag `kong:"help=${VersionHelp}"`
}

var kongVars = kong.Vars{
	"version":     version,
	"VersionHelp": `Output the goproject-tmpl version and exit.`,
}

func main() {
	var cli cmdRoot
	k := kong.Parse(&cli,
		kongVars,
		kong.Description(description),
	)

	//nolint:errcheck // stdout
	_, _ = fmt.Fprintln(k.Stdout, "hello, world")
}
