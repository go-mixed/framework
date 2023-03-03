package console

import (
	"errors"
	"os"
	"strings"

	"github.com/gookit/color"

	"gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/contracts/console/command"
	"gopkg.in/go-mixed/framework.v1/support/file"
	"gopkg.in/go-mixed/framework.v1/support/str"
)

type JobMakeCommand struct {
}

// Signature The name and signature of the console command.
func (c *JobMakeCommand) Signature() string {
	return "make:job"
}

// Description The console command description.
func (c *JobMakeCommand) Description() string {
	return "Create a new job class"
}

// Extend The console command extend.
func (c *JobMakeCommand) Extend() command.Extend {
	return command.Extend{
		Category: "make",
	}
}

// Handle Execute the console command.
func (c *JobMakeCommand) Handle(ctx console.Context) error {
	name := ctx.Argument(0)
	if name == "" {
		return errors.New("Not enough arguments (missing: name) ")
	}

	file.Create(c.getPath(name), c.populateStub(c.getStub(), name))
	color.Greenln("Job created successfully")

	return nil
}

func (c *JobMakeCommand) getStub() string {
	return JobStubs{}.Job()
}

// populateStub Populate the place-holders in the command stub.
func (c *JobMakeCommand) populateStub(stub string, name string) string {
	stub = strings.ReplaceAll(stub, "DummyJob", str.Case2Camel(name))
	stub = strings.ReplaceAll(stub, "DummyName", str.Camel2Case(name))

	return stub
}

// getPath Get the full path to the command.
func (c *JobMakeCommand) getPath(name string) string {
	pwd, _ := os.Getwd()

	return pwd + "/app/jobs/" + str.Camel2Case(name) + ".go"
}
