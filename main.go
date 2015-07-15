package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/codegangsta/cli"

	"github.com/containerops/crew/cmd"
	"github.com/containerops/wrench/setting"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	//Init Config
	if err := setting.SetConfig("conf/containerops.conf"); err != nil {
		fmt.Println("Setting config value error: %s", err.Error())
		os.Exit(1)
	}

	app := cli.NewApp()

	app.Name = setting.AppName
	app.Usage = setting.Usage
	app.Version = setting.Version
	app.Author = setting.Author
	app.Email = setting.Email

	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
