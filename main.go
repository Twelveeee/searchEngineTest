package main

import (
	"fmt"
	"os"
	"searchEngineTest/script"

	"github.com/urfave/cli/v2"
)

var version = "v0.0.1 development"

const appName = "searchEngineTest"
const appAbout = "Twelveeee"

const appDescription = "搜索引擎测试"
const appCopyright = "(c) 2023 Twelveeee @ Twelveeee"

// Metadata contains build specific information.
var Metadata = map[string]interface{}{
	"Name":        appName,
	"About":       appAbout,
	"Description": appDescription,
	"Version":     version,
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("run recover %v \n", err)
			os.Exit(1)
		}
	}()

	app := cli.NewApp()
	app.Usage = appAbout
	app.Description = appDescription
	app.Version = version
	app.Copyright = appCopyright
	app.EnableBashCompletion = true
	app.Commands = script.Commands
	app.Flags = script.Flags
	app.Metadata = Metadata

	// os.Args = append(os.Args, "msi")

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("run error %v \n", err)
	}

	fmt.Printf("run done \n")

}
