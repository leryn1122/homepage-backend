package main

import (
	"os"

	"github.com/leryn1122/homepage/v2/pkg/web"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Usage = "Leryn homepage"

	app.Action = func(c *cli.Context) error {
		return run(c)
	}

	app.ExitErrHandler = func(c *cli.Context, err error) {
		logrus.Fatal(err)
	}
	app.Run(os.Args)
}

func run(cli *cli.Context) error {
	logrus.Infof("Homepage version %s is running.", "0.1.0")
	logrus.Infof("Homepage arguments: ")

	web.StartWebService()
	return nil
}
