package router

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo"
)

type Validator struct{}

func (v *Validator) Validate(i interface{}) error {
	if c, ok := i.(validation.Validatable); ok {
		return c.Validate()
	}
	return nil
}

func SetValidator(e *echo.Echo) {
	e.Validator = &Validator{}
}

func BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	return c.Validate(i)
}
