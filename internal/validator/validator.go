package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError map[string]string

func (v ValidationError) Error() string {
	errs := []string{}

	for _, err := range v {
		errs = append(errs, err)
	}

	return strings.Join(errs, ",")
}

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

	errMap := make(ValidationError)

	for _, err := range validationErrors {
		fieldName := err.Field()

		var errDescription (string)

		switch err.Tag() {
		case "required":
			errDescription = fmt.Errorf("%s is required", fieldName).Error()
		case "min":
			errDescription = fmt.Errorf("%s must be at least %s characters", fieldName, err.Param()).Error()
		case "max":
			errDescription = fmt.Errorf("%s must be no more than %s characters", fieldName, err.Param()).Error()
		case "email":
			errDescription = fmt.Errorf("%s must be a valid email address", fieldName).Error()
		case "url":
			errDescription = fmt.Errorf("%s must be a valid URL", fieldName).Error()
		default:
			errDescription = fmt.Errorf("%s is not valid (failed on %s)", fieldName, err.Tag()).Error()
		}

		errMap[fieldName] = errDescription
	}

	return errMap
}
