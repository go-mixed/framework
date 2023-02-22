package support

import (
	"github.com/pkg/errors"
	"gopkg.in/go-mixed/framework/contracts/container"
)

func As[T any](result container.InstanceResult, as T) (T, error) {
	instance, err := result.Instance, result.Error
	if err != nil {
		return nil, err
	}

	_i, ok := instance.(T)
	if !ok {
		return nil, errors.Errorf("cannot cast %t to %t", instance, as)
	}

	return _i, nil
}

func MustAs[T any](result container.InstanceResult, as T) T {
	instance, err := As(result, as)
	if err != nil {
		panic(err)
	}

	return instance
}
