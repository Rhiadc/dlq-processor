package entity

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Workflow struct {
	ID           string `json:"id"`
	WorkflowName string `json:"workflow_name"`
	BU           string `json:"bu"`
	Steps        []Step `json:"steps"`
}

type Step struct {
	Plugin string `json:"plugin"`
}

func (w Workflow) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.ID, validation.Required, validation.Length(3, 50)),
		validation.Field(&w.WorkflowName, validation.Required, validation.Length(5, 50)),
		validation.Field(&w.BU, validation.Required, validation.Length(3, 50)),
		validation.Field(&w.Steps, validation.Required, validation.NotNil, validation.Each(validation.Required)),
	)
}

func (s Step) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Plugin, validation.Required, validation.NotNil),
	)
}

type WorkflowRepository interface {
	CreateWorkflow(workflow *Workflow) (*Workflow, error)
	DeleteWorkflow(WorkflowID string) error
	GetWorkflow(WorkflowID string) (*Workflow, error)
	GetWorkflows() ([]*Workflow, error)
	UpdateWorkflow(ctx context.Context, workflowID string, data map[string]interface{}) error
}
