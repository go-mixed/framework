package console

import (
	"gopkg.in/go-mixed/framework.v1/contracts/console/command"
	"time"
)

type Command interface {
	//Signature The name and signature of the console command.
	Signature() string
	//Description The console command description.
	Description() string
	//Extend The console command extend.
	Extend() command.Extend
	//Handle Execute the console command.
	Handle(ctx Context) error
}

//go:generate mockery --name=Context
type Context interface {
	Argument(index int) string
	Arguments() []string

	Option(name string) any
	PathOption(name string) string
	TimeOption(name string) *time.Time
	DurationOption(name string) time.Duration
	StringOption(name string) string
	StringSliceOption(name string) []string
	IntSliceOption(name string) []int
	Int64Option(name string) int64
	Uint64Option(name string) uint64
	IntOption(name string) int
	BoolOption(name string) bool
	UintOption(name string) uint
	Float64Option(name string) float64
	Float64SliceOption(name string) []float64

	NumOptions() int
}
