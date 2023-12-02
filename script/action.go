package script

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

var createIndexActionConf = map[string]func(*cli.Context) error{
	"meillsearch": meillSearchCreateIndexAction,
	"typesense":   typeSenseCreateIndexAction,
	"algolia":     algoliaCreateIndexAction,
}

var importDataActionConf = map[string]func(*cli.Context) error{
	"meillsearch": meillSearchImportAction,
	"typesense":   typeSenseImportAction,
	"algolia":     algoliaImportAction,
}

var searchActionConf = map[string]func(*cli.Context) error{
	"meillsearch": meillSearchSearchAction,
	"typesense":   typeSenseSearchAction,
	"algolia":     algoliaSearchAction,
}

var deleteIndexActionConf = map[string]func(*cli.Context) error{
	"meillsearch": meillSearchDeleteIndexAction,
	"typesense":   typeSenseDeleteIndexAction,
	"algolia":     algoliaDeleteIndexAction,
}

var pressureTestActionConf = map[string]func(*cli.Context) error{
	"meillsearch": meillSearchTestAction,
	"typesense":   typeSenseTestAction,
	"algolia":     algoliaTestAction,
}

func getEngineName(ctx *cli.Context) string {
	engine := ctx.String("engine")
	engine = strings.ToLower(engine)
	switch engine {
	case "meillsearch", "ms", "m":
		return "meillsearch"
	case "typesense", "ts", "t":
		return "typesense"
	case "algolia", "a":
		return "algolia"
	default:
		return engine
	}
}

func importDataAction(ctx *cli.Context) error {
	return executeAction(ctx, importDataActionConf)
}

func createIndexAction(ctx *cli.Context) error {
	return executeAction(ctx, createIndexActionConf)
}

func deleteIndexAction(ctx *cli.Context) error {
	return executeAction(ctx, deleteIndexActionConf)
}

func searchAction(ctx *cli.Context) error {
	return executeAction(ctx, searchActionConf)
}

func pressureTestAction(ctx *cli.Context) error {
	return executeAction(ctx, pressureTestActionConf)
}

func executeAction(ctx *cli.Context, actions map[string]func(*cli.Context) error) error {
	engine := getEngineName(ctx)
	if engine == "all" {
		for _, action := range actions {
			err := action(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	}

	action, ok := actions[engine]
	if !ok {
		return fmt.Errorf("engine %s not supported", engine)
	}
	return action(ctx)
}
