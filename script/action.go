package script

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func importDataAction(ctx *cli.Context) error {
	engine := ctx.String("engine")
	engine = strings.ToLower(engine)
	switch engine {
	case "meillsearch", "ms", "m":
		return meillSearchImportAction(ctx)
	case "typesense", "ts", "t":
		return typeSenseImportAction(ctx)
	case "all":
		err := meillSearchImportAction(ctx)
		if err != nil {
			return err
		}
		err = typeSenseImportAction(ctx)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("engine %s not support", engine)
	}
}

func createIndexAction(ctx *cli.Context) error {
	engine := ctx.String("engine")
	engine = strings.ToLower(engine)
	switch engine {
	case "meillsearch", "ms", "m":
		return meillSearchCreateIndexAction(ctx)
	case "typesense", "ts", "t":
		return typeSenseCreateIndexAction(ctx)
	case "all":
		err := meillSearchCreateIndexAction(ctx)
		if err != nil {
			return err
		}
		err = typeSenseCreateIndexAction(ctx)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("engine %s not support", engine)
	}
}

func deleteIndexAction(ctx *cli.Context) error {
	engine := ctx.String("engine")
	engine = strings.ToLower(engine)
	switch engine {
	case "meillsearch", "ms", "m":
		return meillSearchDeleteIndexAction(ctx)
	case "typesense", "ts", "t":
		return typeSenseDeleteIndexAction(ctx)
	case "all":
		err := meillSearchDeleteIndexAction(ctx)
		if err != nil {
			return err
		}
		err = typeSenseDeleteIndexAction(ctx)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("engine %s not support", engine)
	}
}

func searchAction(ctx *cli.Context) error {
	engine := ctx.String("engine")
	engine = strings.ToLower(engine)
	switch engine {
	case "meillsearch", "ms", "m":
		return meillSearchSearchAction(ctx)
	case "typesense", "ts", "t":
		return typeSenseSearchAction(ctx)
	case "all":
		err := meillSearchSearchAction(ctx)
		if err != nil {
			return err
		}
		err = typeSenseSearchAction(ctx)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("engine %s not support", engine)
	}
}

func pressureTestAction(ctx *cli.Context) error {
	engine := ctx.String("engine")
	engine = strings.ToLower(engine)
	switch engine {
	case "meillsearch", "ms", "m":
		return meillSearchTestAction(ctx)
	case "typesense", "ts", "t":
		return typeSenseTestAction(ctx)
	case "all":
		err := meillSearchTestAction(ctx)
		if err != nil {
			return err
		}
		err = typeSenseTestAction(ctx)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("engine %s not support", engine)
	}
}
