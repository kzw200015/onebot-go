package onebot

import "github.com/go-playground/validator/v10"

type EventValidator struct {
	validator *validator.Validate
}

func (ev *EventValidator) Validate(i interface{}) error {
	return ev.validator.Struct(i)
}

func DefaultEventValidator() *EventValidator {
	return &EventValidator{validator: validator.New()}
}
