package console

// 此文件必须放这里，不然循环引用

import (
	"github.com/gookit/color"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/console"
	"gopkg.in/go-mixed/framework.v1/contracts/console/command"
	queue2 "gopkg.in/go-mixed/framework.v1/contracts/queue"
	"runtime"
)

type QueueWorkCommand struct {
}

// Signature The name and signature of the console command.
func (c *QueueWorkCommand) Signature() string {
	return "queue:work"
}

// Description The console command description.
func (c *QueueWorkCommand) Description() string {
	return "Run a new queue worker"
}

// Extend The console command extend.
func (c *QueueWorkCommand) Extend() command.Extend {
	return command.Extend{
		Category: "queue",
		Flags: []command.Flag{
			&command.StringFlag{
				Name:    "queue",
				Value:   "",
				Aliases: []string{},
				Usage:   "The names of the queues to work",
			},
			&command.IntFlag{
				Name:    "workers",
				Value:   runtime.NumCPU(),
				Aliases: []string{},
				Usage:   "The count of the workers",
			},
		},
	}
}

// Handle Execute the console command.
func (c *QueueWorkCommand) Handle(ctx console.Context) error {
	connectionName := ctx.Argument(0)
	queueName := ctx.StringOption("queue")
	workers := ctx.IntOption("workers")

	color.Greenln("Run queue worker")

	return container.MustMakeAs("queue", queue2.IBroker(nil)).Connection(connectionName).RunServe(queueName, workers)
}
