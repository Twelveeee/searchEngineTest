package script

import "github.com/urfave/cli/v2"

var Commands = []*cli.Command{
	&MeillSearchImportCommand,
	&MeillSearchOnceCommand,
	&MeillSearchDeleteIndexCommand,
	&MeillSearchTestCommand,
	&TypeSenseImportCommand,
	&TypeSenseOnceCommand,
}

var Flags = []cli.Flag{
	&configFlag,
}

var MeillSearchImportCommand = cli.Command{
	Name:    "MeillSearchImport",
	Aliases: []string{"msi"},
	Usage:   "init MeillSearch data",
	Action:  MeillSearchImportAction,
}

var MeillSearchOnceCommand = cli.Command{
	Name:    "MeillSearchOnce",
	Aliases: []string{"mso"},
	Usage:   "search MeillSearch data once",
	Flags:   searchOnceFlags,
	Action:  MeillSearchOnceAction,
}

var MeillSearchDeleteIndexCommand = cli.Command{
	Name:    "MeillSearchDeleteIndex",
	Aliases: []string{"msdi"},
	Usage:   "delete MeillSearch index",
	Action:  MeillSearchDeleteIndexAction,
}
var MeillSearchTestCommand = cli.Command{
	Name:    "MeillSearchTest",
	Aliases: []string{"mst"},
	Usage:   "test MeillSearch",
	Action:  MeillSearchTestAction,
}

var TypeSenseImportCommand = cli.Command{
	Name:    "TypeSenseImport",
	Aliases: []string{"tsi"},
	Usage:   "init TypeSense data",
	Action:  TypeSenseImportAction,
}

var TypeSenseOnceCommand = cli.Command{
	Name:    "TypeSenseOnce",
	Aliases: []string{"tso"},
	Usage:   "search TypeSense data once",
	Flags:   searchOnceFlags,
	Action:  TypeSenseOnceAction,
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
