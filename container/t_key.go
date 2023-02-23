package container

import "reflect"

// Key the key type for the Injector.
type tKey struct {
	val       string
	isTypeKey bool
}

func newSKey(abstract string) tKey {
	return tKey{
		val:       abstract,
		isTypeKey: false,
	}
}

func newTKey(abstract any) tKey {
	return tKey{
		val:       reflect.ValueOf(abstract).Type().Elem().String(),
		isTypeKey: true,
	}
}

func toTKey(abstract any) tKey {
	switch abstract.(type) {
	case string:
		return newSKey(abstract.(string))
	default:
		return newTKey(abstract)
	}
}

func (t tKey) String() string {
	return t.val
}
