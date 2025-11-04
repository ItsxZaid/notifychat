package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(string(field.Tag.Get("json")), ",", 2)[0]

		if name == "-" {
			return ""
		}
		return name
	})

	return &Validator{
		validate: validate,
	}
}

func (v *Validator) Validate(s interface{}) error {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return err
	}

	if len(validationErrors) == 0 {
		return err
	}

	fieldErr := validationErrors[0]

	fieldName := fieldErr.Field()

	switch fieldErr.Tag() {
	case "required":
		return fmt.Errorf("%s is required", fieldName)
	case "min":
		return fmt.Errorf("%s must be at least %s characters", fieldName, fieldErr.Param())
	case "max":
		return fmt.Errorf("%s must be no more than %s characters", fieldName, fieldErr.Param())
	case "email":
		return fmt.Errorf("%s must be a valid email address", fieldName)
	case "url":
		return fmt.Errorf("%s must be a valid URL", fieldName)
	default:
		return fmt.Errorf("%s is not valid (failed on %s)", fieldName, fieldErr.Tag())
	}
}
