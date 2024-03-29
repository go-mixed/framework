package command

import "github.com/urfave/cli/v2"

type Extend struct {
	Category string
	Flags    []Flag
}

type Flag = cli.Flag

type IntFlag = cli.IntFlag
type Int64Flag = cli.Int64Flag
type IntSliceFlag = cli.IntSliceFlag
type Int64SliceFlag = cli.Int64SliceFlag
type UintFlag = cli.UintFlag
type Uint64Flag = cli.Uint64Flag
type Float64Flag = cli.Float64Flag
type Float64SliceFlag = cli.Float64SliceFlag

type StringFlag = cli.StringFlag
type StringSliceFlag = cli.StringSliceFlag

type BoolFlag = cli.BoolFlag

type DurationFlag = cli.DurationFlag
type TimeFlag = cli.TimestampFlag

type ExitCoder = cli.ExitCoder
type MultiError = cli.MultiError

func Exit(message any, code int) ExitCoder {
	return cli.Exit(message, code)
}
