package validator

import (
	"learn-go/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task, validation.Field(&task.Title, validation.Required.Error("title is required"), validation.RuneLength(1, 64).Error("limited max 64 chars")))
}
