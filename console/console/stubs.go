package console

type Stubs struct {
}

func (receiver Stubs) Command() string {
	return `package commands

import (
	"gopkg.in/go-mixed/framework/contracts/console"
	"gopkg.in/go-mixed/framework/contracts/console/command"
)

type DummyCommand struct {
}

//Signature The name and signature of the console command.
func (receiver *DummyCommand) Signature() string {
	return "command:name"
}

//Description The console command description.
func (receiver *DummyCommand) Description() string {
	return "Command description"
}

//Extend The console command extend.
func (receiver *DummyCommand) Extend() command.Extend {
	return command.Extend{}
}

//Handle Execute the console command.
func (receiver *DummyCommand) Handle(ctx console.Context) error {
	
	return nil
}
`
}
