package console

import (
	"gopkg.in/go-mixed/framework/contracts/console"
	"gopkg.in/go-mixed/framework/contracts/console/command"
	"gopkg.in/go-mixed/framework/facades"
)

type ListCommand struct {
}

// Signature The name and signature of the console command.
func (receiver *ListCommand) Signature() string {
	return "list"
}

// Description The console command description.
func (receiver *ListCommand) Description() string {
	return "List commands"
}

// Extend The console command extend.
func (receiver *ListCommand) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *ListCommand) Handle(ctx console.Context) error {
	facades.Artisan.Call("--help")

	return nil
}
