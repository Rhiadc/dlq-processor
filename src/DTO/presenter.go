package dto

import "github.com/dlqProcessor/src/entity"

type CreateWorkflowResponse struct {
	ID           string `json:"id"`
	WorkflowName string `json:"workflow_name,omitempty"`
	BU           string `json:"bu,omitempty"`
}

type AllWorkflows struct {
	Workflows []CreateWorkflowResponse
}

type WorkflowCreated struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	ID      string `json:"id"`
}

type SuccessResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

func ToResponse(entity *entity.Workflow) *CreateWorkflowResponse {
	return &CreateWorkflowResponse{
		ID:           entity.ID,
		WorkflowName: entity.WorkflowName,
		BU:           entity.BU,
	}
}

//@TODO: precisamos incluir lógica dos plugins
func ToDomain(req *WorkflowRequest) *entity.Workflow {
	return &entity.Workflow{
		WorkflowName: req.WorkflowName,
		BU:           req.BU,
	}
}
