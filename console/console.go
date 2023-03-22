package console

import (
	"os"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

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

func (c *Console) Register(commands []console.Command) {
	for _, item := range commands {
		item := item
		cliCommand := cli.Command{
			Name:  item.Signature(),
			Usage: item.Description(),
			Action: func(ctx *cli.Context) error {
				return item.Handle(&CliContext{ctx})
			},
		}

		cliCommand.Category = item.Extend().Category
		cliCommand.Flags = item.Extend().Flags
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

type CliContext struct {
	instance *cli.Context
}

var _ console.Context = (*CliContext)(nil)

func (r *CliContext) Option(name string) any {
	return r.instance.Value(name)
}

func (r *CliContext) PathOption(name string) string {
	return r.instance.Path(name)
}

func (r *CliContext) TimeOption(name string) *time.Time {
	return r.instance.Timestamp(name)
}

func (r *CliContext) DurationOption(name string) time.Duration {
	return r.instance.Duration(name)
}

func (r *CliContext) StringOption(name string) string {
	return r.instance.String(name)
}

func (r *CliContext) StringSliceOption(name string) []string {
	return r.instance.StringSlice(name)
}

func (r *CliContext) IntSliceOption(name string) []int {
	return r.instance.IntSlice(name)
}

func (r *CliContext) Int64Option(name string) int64 {
	return r.instance.Int64(name)
}

func (r *CliContext) Uint64Option(name string) uint64 {
	return r.instance.Uint64(name)
}

func (r *CliContext) IntOption(name string) int {
	return r.instance.Int(name)
}

func (r *CliContext) BoolOption(name string) bool {
	return r.instance.Bool(name)
}

func (r *CliContext) UintOption(name string) uint {
	return r.instance.Uint(name)
}

func (r *CliContext) Float64Option(name string) float64 {
	return r.instance.Float64(name)
}

func (r *CliContext) Float64SliceOption(name string) []float64 {
	return r.instance.Float64Slice(name)
}

func (r *CliContext) NumOptions() int {
	return r.instance.NumFlags()
}

func (r *CliContext) Argument(index int) string {
	return r.instance.Args().Get(index)
}

func (r *CliContext) Arguments() []string {
	return r.instance.Args().Slice()
}
