package dto_test

import (
	"testing"

	dto "github.com/dlqProcessor/src/DTO"
	"github.com/dlqProcessor/src/entity"
	"github.com/stretchr/testify/assert"
)

func TestToResponse(t *testing.T) {
	workflow := &entity.Workflow{
		ID:           "abc-123",
		BU:           "some-bu",
		WorkflowName: "some-name",
	}

	response := dto.ToResponse(workflow)
	assert.Equal(t, workflow.ID, response.ID)
	assert.Equal(t, workflow.BU, response.BU)
	assert.Equal(t, workflow.WorkflowName, response.WorkflowName)
}

func TestToDomain(t *testing.T) {
	workflowRequest := &dto.WorkflowRequest{
		WorkflowName: "some-name",
		BU:           "some-bu",
	}

	workflow := dto.ToDomain(workflowRequest)
	assert.Equal(t, workflow.BU, workflowRequest.BU)
	assert.Equal(t, workflow.WorkflowName, workflowRequest.WorkflowName)
}
