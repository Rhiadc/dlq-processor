package dto

import validation "github.com/go-ozzo/ozzo-validation"

func (d DeleteRequest) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.WorkflowID, validation.Required, validation.Length(3, 50)),
	)
}

func (w WorkflowRequest) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.WorkflowName, validation.Required, validation.Length(5, 50)),
		validation.Field(&w.BU, validation.Required, validation.Length(3, 50)),
		validation.Field(&w.Steps, validation.Required, validation.NotNil, validation.Each(validation.Required)),
	)
}

func (s Step) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Plugin, validation.Required, validation.Length(5, 50)),
		validation.Field(&s.Events, validation.Required, validation.Each(validation.NotNil)),
		validation.Field(&s.EntryPoints, validation.Required, validation.NotNil, validation.Each(validation.NotNil)),
	)
}

func (e EntryPoint) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Response, validation.Each(validation.NotNil)),
	)
}

func (w UpdateWorkflowRequest) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.WorkflowName, validation.Length(5, 50)),
		validation.Field(&w.BU, validation.Length(3, 50)),
		validation.Field(&w.UpdateSteps),
	)
}

func (s UpdateStep) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Plugin, validation.Length(5, 50)),
		validation.Field(&s.Events, validation.Each(validation.NotNil, validation.Length(5, 50))),
		validation.Field(&s.UpdateEntryPoints, validation.Each(validation.NotNil)),
	)
}

func (e UpdateEntryPoint) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Response, validation.Required, validation.Each(validation.NotNil, validation.Length(5, 50))),
	)
}
