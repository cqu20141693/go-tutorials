package main

import (
	"flag"
	"fmt"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/cmd"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source"
	mcli "go-micro.dev/v4/config/source/cli"
	"log"
	"os"
)

func main() {
	//
	fmt.Println(os.Args)
	testArgs()

	testApp()

}

func testArgs() {
	set := flag.NewFlagSet("app", 0)
	set.Int("flags", 12, "doc")
	ctx := cli.NewContext(nil, set, nil)
	app := cmd.App()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "cc.profiles.active",
			Aliases: []string{"active"},
			EnvVars: []string{"cc.profiles.active"},
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "oss",
			EnvVars: []string{"oss"},
			Value:   "",
		},
	}

	ctx.App = app
	var src source.Source
	app.Action = func(c *cli.Context) error {
		src = mcli.WithContext(c)
		return nil
	}
	//os.Args = args
	_ = app.Run(os.Args)
	conf, _ := config.NewConfig()
	err := conf.Load(src)
	if err != nil {
		log.Fatal(err)
		return
	}
	m := conf.Map()
	for s, i := range m {
		fmt.Println(s, i)
	}
	fmt.Println(m)
}

func testApp() {
	set := flag.NewFlagSet("test", 0)
	set.Int("flags", 12, "doc")
	parentSet := flag.NewFlagSet("test", 0)
	parentSet.Int("top-flag", 13, "doc")
	parentCtx := cli.NewContext(nil, parentSet, nil)
	c := cli.NewContext(nil, set, parentCtx)
	app := cmd.App()
	c.App = app
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "oss",
			EnvVars: []string{"oss"},
			Value:   "file",
		},
	}

	args := []string{"run", "-oss", "aliyun"}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("oss=%s", c.String("oss"))
		return nil
	}
	//os.Args = args
	app.Run(args)
}
