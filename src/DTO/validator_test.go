package dto_test

import (
	"testing"

	dto "github.com/dlqProcessor/src/DTO"
	"github.com/stretchr/testify/assert"
)

//@TODO validar se o ID também esta sendo feito um update

func TestValidateValidWorkflowRequest(t *testing.T) {
	entrypoint := dto.EntryPoint{
		Request:  map[string]interface{}{"something": []string{"some", "thing"}},
		Response: []string{"some", "thing"},
	}
	step := dto.Step{
		Plugin:      "plugin-name",
		Events:      []string{"some-string", "another-string"},
		EntryPoints: []dto.EntryPoint{entrypoint},
	}
	validWorkflow := dto.WorkflowRequest{
		WorkflowName: "some-name",
		BU:           "some-bu",
		Steps:        []dto.Step{step},
	}

	err := validWorkflow.Validate()
	assert.Nil(t, err)
}

func TestValidateValidUpdateWorkflowRequest(t *testing.T) {
	//update without step
	validWorkflow := dto.UpdateWorkflowRequest{
		WorkflowName: "some-name",
		BU:           "some-bu",
	}
	err := validWorkflow.Validate()
	assert.Nil(t, err)

	//validate workflow with step without entrypoints
	step := dto.UpdateStep{
		Plugin: "plugin-name",
		Events: []string{"some-string", "another-string"},
	}
	validWorkflow.UpdateSteps = []dto.UpdateStep{step}

	err = validWorkflow.Validate()
	assert.Nil(t, err)

	//validate workflow with valid entrypoint
	entrypoint := dto.UpdateEntryPoint{
		Request:  map[string]interface{}{"something": []string{"some", "thing"}},
		Response: []string{"teste"},
	}

	err = entrypoint.Validate()
	assert.Nil(t, err)

	step.UpdateEntryPoints = []dto.UpdateEntryPoint{entrypoint}
	err = step.Validate()
	assert.Nil(t, err)
}

func TestValidateStepWithoutFields(t *testing.T) {
	step := dto.UpdateStep{
		Events: []string{"some-string", "another-string"},
	}
	err := step.Validate()
	assert.Nil(t, err)
}

func TestValidateInvalidUpdateEntrypoint(t *testing.T) {
	//Response with empty array
	entrypoint := dto.UpdateEntryPoint{
		Request:  map[string]interface{}{"something": []string{"some", "thing"}},
		Response: []string{},
	}

	err := entrypoint.Validate()
	assert.NotNil(t, err)
}

func TestUpdateStepWithoutEntrypoint(t *testing.T) {
	step := dto.UpdateStep{
		Plugin: "plugin-name",
		Events: []string{"some-string", "another-string"},
	}

	err := step.Validate()
	assert.Nil(t, err)
}

func TestValidDeleteRequest(t *testing.T) {
	deleteRequest := dto.DeleteRequest{
		WorkflowID: "ashduiy89rhofu849urodj",
	}
	err := deleteRequest.Validate()
	assert.Nil(t, err)
}

func TestInvalidDeleteRequest(t *testing.T) {
	deleteRequest := dto.DeleteRequest{
		WorkflowID: "",
	}
	err := deleteRequest.Validate()
	assert.NotNil(t, err)
}
