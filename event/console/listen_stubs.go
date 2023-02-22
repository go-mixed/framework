package console

type ListenerStubs struct {
}

func (receiver ListenerStubs) Listener() string {
	return `package listeners

import (
	"gopkg.in/go-mixed/framework/contracts/event"
)

type DummyListener struct {
}

func (receiver *DummyListener) Signature() string {
	return "DummyName"
}

func (receiver *DummyListener) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     false,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *DummyListener) Handle(args ...any) error {
	return nil
}
`
}
