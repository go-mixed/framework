package console

type Stubs struct {
}

func (r Stubs) Model() string {
	return `package models

import (
	"gopkg.in/go-mixed/framework/database/orm"
)

type DummyModel struct {
	orm.Model
}
`
}
