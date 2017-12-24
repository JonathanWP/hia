/**
* @file main.go
* @Synopsis
* @author alan lin
* @version 1.0
* @date 2017-11-30
 */
package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
	"hia/cmd/spark"
	http "hia/http/httpraw"
	"hia/lab"
	"hia/redis"
)

func helloCommand(c *cli.Context) error {
	fmt.Println("run hello command: ", c.Args().First())
	return nil
}

func lab1Command(c *cli.Context) error {
	fmt.Println("lab1Command: ", c.Args().First())
	lab.Lab1Command()
	return nil
}

func lab2Command(c *cli.Context) error {
	fmt.Println("lab1Command: ", c.Args().First())
	lab.Lab2Command()
	return nil
}

func testNetCommand(c *cli.Context) error {
	fmt.Println("test net: ", c.Args().First())
	return nil
}

func testRedisCommand(c *cli.Context) error {
	fmt.Println("test redis: ", c.Args().First())
	redis.TestRedis()
	return nil
}

func testHttpcCommand(c *cli.Context) error {
	http.TestHttpcCommand()
	return nil
}

func testHttpdCommand(c *cli.Context) error {
	http.TestHttpdCommand()
	return nil
}

var (
	app = cli.NewApp()

	addCommand = cli.Command{
		Name:    "hello",
		Aliases: []string{"a"},
		Usage:   "hello test command",
		Action:  helloCommand,
	}

	labCommand = cli.Command{
		Name:    "lab",
		Aliases: []string{"l"},
		Usage:   "options for task templates",
		Subcommands: []cli.Command{
			{
				Name:   "lab1",
				Usage:  "lab net",
				Action: lab1Command,
			},
			{
				Name:   "lab2",
				Usage:  "lab net",
				Action: lab2Command,
			},
		},
	}

	sparkCommand = cli.Command{
		Name:    "spark",
		Aliases: []string{"s"},
		Usage:   "options for task templates",
		Action:  spark.NewSpark,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "network2",
				Usage: "name of the network to administer",
				Value: "hello world",
			},
			cli.BoolFlag{
				Name:  "rc",
				Usage: "Enable the HTTP-RPC server",
			},
		},
	}
)

func init() {
	app.Action = spark.NewSpark
	app.Commands = []cli.Command{
		addCommand,
		labCommand,
		sparkCommand,
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "network1",
			Usage: "name of the network to administer",
			Value: "hello world",
		},
		cli.BoolFlag{
			Name:  "rpc",
			Usage: "Enable the HTTP-RPC server",
		},
	}

}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
