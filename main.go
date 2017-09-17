package main

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/mkideal/cli"
	"os"
)

var SelfVersion = "no version"

type nextArgs struct {
	cli.Helper
	Build   []string `cli:"b,build" usage:"build metadata"`
	Change  string   `cli:"c,change" usage:"What version has changed?\n\t\t\t\tM - Major	m - minor\n\t\t\t\tp  -patch	n - none" dft:"p"`
	Current string   `cli:"current" usage:"Current version" dft:"0.0.0"`
	KeepPre bool     `cli:"k,keep-pre" usage:"Keep Prerelease version" dft:"false"`
	Pre     []string `cli:"p,pre" usage:"Prerelease version. incompatible with --keep-pre=true"`
}

func (a *nextArgs) Validate(*cli.Context) error {
	switch a.Change {
	case "M":
	case "m":
	case "p":
	case "n":
	default:
		return fmt.Errorf("--change  M|m|p|n is allowed")
	}
	if _, err := semver.New(a.Current); err != nil {
		return err
	}

	if a.KeepPre && len(a.Pre) > 0 {
		return fmt.Errorf("--keep-pre=true is not compatible with --pre")
	}

	return nil
}

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
	Argv: func() interface{} { return new(nextArgs) },
	Fn: func(ctx *cli.Context) error {
		params := ctx.Argv().(*nextArgs)
		version, err := Next(params)
		if err != nil {
			return err
		}
		ctx.String("%s\n", version.String())
		return nil
	},
}

func Next(params *nextArgs) (semver.Version, error) {
	oldVersion, _ := semver.Parse(params.Current)
	newVersion := semver.Version{}

	switch params.Change {
	case "M":
		newVersion.Major = oldVersion.Major + 1
		newVersion.Minor = 0
		newVersion.Patch = 0
	case "m":
		newVersion.Major = oldVersion.Major
		newVersion.Minor = oldVersion.Minor + 1
		newVersion.Patch = 0
	case "p":
		newVersion.Major = oldVersion.Major
		newVersion.Minor = oldVersion.Minor
		newVersion.Patch = oldVersion.Patch + 1
	case "n":
		newVersion.Major = oldVersion.Major
		newVersion.Minor = oldVersion.Minor
		newVersion.Patch = oldVersion.Patch
	default:
		return newVersion, fmt.Errorf("change param is incorrect. Please, see --help")
	}
	if params.KeepPre {
		newVersion.Pre = oldVersion.Pre
	} else if len(params.Pre) > 0 {
		for _, pre := range params.Pre {
			if pr, err := semver.NewPRVersion(pre); err == nil {
				newVersion.Pre = append(newVersion.Pre, pr)
			}
		}
	}
	if len(params.Build) > 0 {
		newVersion.Build = params.Build
	}
	return newVersion, nil
}
