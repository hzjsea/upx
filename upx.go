package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/urfave/cli"
)

const VERSION = "v0.3.6"

func main() {
	initProgress()
	progress.Start()
	defer progress.Stop()

	app := cli.NewApp()
	app.Name = "upx"
	app.Usage = "a tool for driving UpYun Storage"
	app.Author = "Hongbo.Mo"
	app.Email = "zjutpolym@gmail.com"
	app.Version = fmt.Sprintf("%s %s/%s %s", VERSION,
		runtime.GOOS, runtime.GOARCH, runtime.Version())
	app.EnableBashCompletion = true
	app.Compiled = time.Now()
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "quiet, q", Usage: "not verbose"},
		cli.StringFlag{Name: "auth", Usage: "auth string"},
		cli.StringFlag{Name: "config, f", Usage: "Set User Config"},
	}
	app.Before = func(c *cli.Context) error {
		if c.Bool("q") {
			isVerbose = false
		}
		if c.String("auth") != "" {
			err := authStrToConfig(c.String("auth"))
			if err != nil {
				PrintErrorAndExit("%s: invalid auth string", c.Command.FullName())
			}
		}
		if c.String("f") != "" {
			var group, fpath string
			fpathSlice := strings.Split(c.String("f"), ",")
			fpath = fpathSlice[0]
			if len(fpathSlice) == 2 {
				group = fpathSlice[1]
			}
			if _, err := os.Stat(fpath); err != nil {
				PrintErrorAndExit("%s: file path ", err)
			}

			ConfigFromFile(fpath, group)
		}
		return nil
	}
	app.Commands = []cli.Command{
		NewLoginCommand(),
		NewLogoutCommand(),
		NewListSessionsCommand(),
		NewSwitchSessionCommand(),
		NewInfoCommand(),
		NewCdCommand(),
		NewPwdCommand(),
		NewMkdirCommand(),
		NewLsCommand(),
		NewTreeCommand(),
		NewGetCommand(),
		NewPutCommand(),
		NewRmCommand(),
		NewSyncCommand(),
		NewAuthCommand(),
		NewPostCommand(),
		NewPurgeCommand(),
		NewGetDBCommand(),
		NewCleanDBCommand(),
		NewUpgradeCommand(),
	}

	app.Run(os.Args)
}
