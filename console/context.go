package console

import (
	"github.com/urfave/cli/v2"
	"gopkg.in/go-mixed/framework.v1/contracts/console"
	"time"
)

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

func (r *CliContext) HasOption(name string) bool {
	return r.instance.IsSet(name)
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
