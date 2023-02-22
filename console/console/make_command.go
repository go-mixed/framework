package console

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/contracts/console/command"
	"gopkg.in/go-mixed/framework.v1/support/file"
	"gopkg.in/go-mixed/framework.v1/support/str"
)

type MakeCommand struct {
}

// Signature The name and signature of the console command.
func (receiver *MakeCommand) Signature() string {
	return "make:command"
}

// Description The console command description.
func (receiver *MakeCommand) Description() string {
	return "Create a new Artisan command"
}

// Extend The console command extend.
func (receiver *MakeCommand) Extend() command.Extend {
	return command.Extend{
		Category: "make",
	}
}

// Handle Execute the console command.
func (receiver *MakeCommand) Handle(ctx console.Context) error {
	name := ctx.Argument(0)
	if name == "" {
		return errors.New("Not enough arguments (missing: name) ")
	}

	file.Create(receiver.getPath(name), receiver.populateStub(receiver.getStub(), name))

	return nil
}

func (receiver *MakeCommand) getStub() string {
	return Stubs{}.Command()
}

// populateStub Populate the place-holders in the command stub.
func (receiver *MakeCommand) populateStub(stub string, name string) string {
	return strings.ReplaceAll(stub, "DummyCommand", str.Case2Camel(name))
}

// getPath Get the full path to the command.
func (receiver *MakeCommand) getPath(name string) string {
	pwd, _ := os.Getwd()

	return pwd + "/app/console/commands/" + str.Camel2Case(name) + ".go"
}
