package broker

import (
	"fmt"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework.v1/contracts/event"
	"gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/facades/config"
)

func encodeArgs(args []queue.Argument) []tasks.Arg {
	var mArgs []tasks.Arg
	for _, arg := range args {
		mArgs = append(mArgs, tasks.Arg{
			Name:  arg.Name,
			Type:  arg.Type,
			Value: arg.Value,
		})
	}

	return mArgs
}

func decodeArgs(args []queue.Argument) []any {
	var mArgs []any
	for _, arg := range args {
		mArgs = append(mArgs, arg.Value)
	}

	return mArgs
}

func eventsToJobMap(events map[event.Event][]event.Listener) (map[string]any, error) {
	jobMap := make(map[string]any)

	for _, listeners := range events {
		for _, listener := range listeners {
			if listener.Signature() == "" {
				return nil, errors.New("the Signature of listener can't be empty")
			}

			if jobMap[listener.Signature()] != nil {
				continue
			}

			jobMap[listener.Signature()] = listener.Handle
		}
	}

	return jobMap, nil
}

func GetQueueName(connection, queue string) string {
	appName := config.GetString("app.name")
	if appName == "" {
		appName = "laravel"
	}
	if connection == "" {
		connection = config.GetString("queue.default")
	}
	if queue == "" {
		queue = config.GetString(fmt.Sprintf("queue.connections.%s.queue", connection), "default")
	}

	return fmt.Sprintf("%s_%s:%s", appName, "queues", queue)
}
