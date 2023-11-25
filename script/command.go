package script

import "github.com/urfave/cli/v2"

var Commands = []*cli.Command{
	&MeillSearchImportCommand,
	&MeillSearchSearchCommand,
	&MeillSearchDeleteIndexCommand,
	&MeillSearchTestCommand,
	&TypeSenseImportCommand,
	&TypeSenseSearchCommand,
}

var Flags = []cli.Flag{
	&configFlag,
}

var MeillSearchImportCommand = cli.Command{
	Name:    "MeillSearchImport",
	Aliases: []string{"mi"},
	Usage:   "init MeillSearch data",
	Action:  MeillSearchImportAction,
}

var MeillSearchSearchCommand = cli.Command{
	Name:    "MeillSearchSearch",
	Aliases: []string{"ms"},
	Usage:   "search MeillSearch data once",
	Flags:   searchOnceFlags,
	Action:  MeillSearchSearchAction,
}

var MeillSearchDeleteIndexCommand = cli.Command{
	Name:    "MeillSearchDeleteIndex",
	Aliases: []string{"mdi"},
	Usage:   "delete MeillSearch index",
	Action:  MeillSearchDeleteIndexAction,
}
var MeillSearchTestCommand = cli.Command{
	Name:    "MeillSearchTest",
	Aliases: []string{"mt"},
	Usage:   "test MeillSearch",
	Action:  MeillSearchTestAction,
}

var TypeSenseImportCommand = cli.Command{
	Name:    "TypeSenseImport",
	Aliases: []string{"ti"},
	Usage:   "init TypeSense data",
	Action:  TypeSenseImportAction,
}

var TypeSenseSearchCommand = cli.Command{
	Name:    "TypeSenseSearch",
	Aliases: []string{"ts"},
	Usage:   "search TypeSense data once",
	Flags:   searchOnceFlags,
	Action:  TypeSenseSearchAction,
}

var queryFlag = cli.StringFlag{
	Name:     "query",
	Aliases:  []string{"q"},
	Usage:    "input `QUERY`",
	Required: true,
}

var configFlag = cli.StringFlag{
	Name:    "config",
	Aliases: []string{"c"},
	Usage:   "config file path",
	Value:   "./config.yaml",
}

var searchOnceFlags = []cli.Flag{
	&queryFlag,
}
