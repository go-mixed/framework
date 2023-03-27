package console

import (
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
	"os"
	"strings"

	"gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/support"
)

type Console struct {
	instance *cli.App
}

func NewConsole() console.IArtisan {
	instance := cli.NewApp()
	instance.Name = "Laravel Framework"
	instance.Usage = support.Version
	instance.UsageText = "artisan [global options] command [options] [arguments...]"

	return &Console{instance}
}

func (c *Console) Register(commands []console.ICommand) {
	for _, item := range commands {
		cliCommand := WrapICommand(item)

		c.instance.Commands = append(c.instance.Commands, &cliCommand)
	}
}

// Call Run an Artisan console command by name.
func (c *Console) Call(command string) {
	c.Run(append([]string{os.Args[0], "artisan"}, strings.Split(command, " ")...), false)
}

// CallAndExit Run an Artisan console command by name and exit.
func (c *Console) CallAndExit(command string) {
	c.Run(append([]string{os.Args[0], "artisan"}, strings.Split(command, " ")...), true)
}

// Run a command. Args come from os.Args.
func (c *Console) Run(args []string, exitIfArtisan bool) {
	if len(args) >= 2 {
		if args[1] == "artisan" {
			if len(args) == 2 {
				args = append(args, "--help")
			}

			if args[2] != "-V" && args[2] != "--version" {
				cliArgs := append([]string{args[0]}, args[2:]...)
				if err := c.instance.Run(cliArgs); err != nil {
					panic(err.Error())
				}
			}

			printResult(args[2])

			if exitIfArtisan {
				os.Exit(0)
			}
		}
	}
}

func printResult(command string) {
	switch command {
	case "make:command":
		color.Greenln("Console command created successfully")
	case "-V", "--version":
		color.Greenln("Laravel Framework " + support.Version)
	}
}
