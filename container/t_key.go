package container

import (
	"reflect"
)

type VTConstraint interface {
	any
}

// Key the key type for the Injector.
type tKey struct {
	val       string
	isTypeKey bool
}

// create a string key
func newSKey(abstract string) tKey {
	return tKey{
		val:       abstract,
		isTypeKey: false,
	}
}

// create a type key
// abstract is unused, but it can show type in debugging mode
func newTKey[VT VTConstraint](abstract any) tKey {
	return tKey{
		val:       reflect.ValueOf(new(VT)).Type().Elem().String(),
		isTypeKey: true,
	}
}

func toTKey[VT VTConstraint](abstract any) tKey {
	switch abstract.(type) {
	case string:
		return newSKey(abstract.(string))
	default:
		return newTKey[VT](abstract)
	}
}

func (t tKey) String() string {
	return t.val
}
