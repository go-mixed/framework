package validation

import (
	"net/url"

	"github.com/gookit/validate"

	httpvalidate "gopkg.in/go-mixed/framework.v1/contracts/validation"
)

func init() {
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		opt.SkipOnEmpty = false
	})
}

type Validator struct {
	instance *validate.Validation
	data     validate.DataFace
}

func NewValidator(instance *validate.Validation, data validate.DataFace) *Validator {
	instance.Validate()

	return &Validator{instance: instance, data: data}
}

func (v *Validator) Bind(ptr any) error {
	var data any
	if _, ok := v.data.Src().(url.Values); ok {
		values := make(map[string]string)
		for key, value := range v.data.Src().(url.Values) {
			if len(value) > 0 {
				values[key] = value[0]
			}
		}

		data = values
	} else {
		data = v.data.Src()
	}

	bts, err := validate.Marshal(data)
	if err != nil {
		return err
	}

	return validate.Unmarshal(bts, ptr)
}

func (v *Validator) Errors() httpvalidate.Errors {
	if v.instance.Errors == nil || len(v.instance.Errors) == 0 {
		return nil
	}

	return NewErrors(v.instance.Errors)
}

func (v *Validator) Fails() bool {
	return v.instance.IsFail()
}
