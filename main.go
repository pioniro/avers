package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"os"
)

var SelfVersion = "no version"

func main() {
	if err := cli.Root(
		root,
		cli.Tree(help),
		cli.Tree(next),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type rootT struct {
	cli.Helper
	Version bool `cli:"version" usage:"Show program version"`
}

func (argv *rootT) AutoHelp() bool {
	return argv.Help
}

var root = &cli.Command{
	Name: "root",
	Desc: "Avers - AutoVERSion program",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		if argv.Version {
			ctx.String("version %s\n", SelfVersion)
		}
		return nil
	},
}

var help = cli.HelpCommand("display help information")

var next = &cli.Command{
	Name: "next",
	Desc: "Next version generation tools",
	Argv: func() interface{} { return new(NextArgs) },
	Fn: func(ctx *cli.Context) error {
		params := ctx.Argv().(*NextArgs)
		version, err := Next(params)
		if err != nil {
			return err
		}
		ctx.String("%s\n", version.String())
		return nil
	},
}
