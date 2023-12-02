package script

import "github.com/urfave/cli/v2"

var Commands = []*cli.Command{
	&importDataCommand,
	&createIndexCommand,
	&deleteIndexCommand,
	&searchCommand,
	&pressureTestCommand,
}

var Flags = []cli.Flag{
	&configFlag,
	&engineFlag,
}

var importDataCommand = cli.Command{
	Name:    "importData",
	Aliases: []string{"i"},
	Usage:   "init data",
	Action:  importDataAction,
}

var createIndexCommand = cli.Command{
	Name:    "createIndex",
	Aliases: []string{"ci"},
	Usage:   "create index",
	Action:  createIndexAction,
}

var deleteIndexCommand = cli.Command{
	Name:    "deleteIndex",
	Aliases: []string{"di"},
	Usage:   "delete index",
	Action:  deleteIndexAction,
}

var searchCommand = cli.Command{
	Name:    "search",
	Aliases: []string{"s"},
	Usage:   "search",
	Action:  searchAction,
	Flags:   searchFlags,
}

var pressureTestCommand = cli.Command{
	Name:    "pressureTest",
	Aliases: []string{"pt"},
	Usage:   "pressureTest",
	Action:  pressureTestAction,
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

var engineFlag = cli.StringFlag{
	Name:    "engine",
	Aliases: []string{"e"},
	Usage:   "set search engine; m as meillsearch, t as typesense, a as aligolia ,all",
}

var searchFlags = []cli.Flag{
	&queryFlag,
}
