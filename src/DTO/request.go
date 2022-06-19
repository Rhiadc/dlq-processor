package dto

type WorkflowRequest struct {
	WorkflowName string `json:"workflow_name"`
	BU           string `json:"bu"`
	Steps        []Step `json:"steps"`
}

type UpdateWorkflowRequest struct {
	ID           string       `json:"id"`
	WorkflowName string       `json:"workflow_name"`
	BU           string       `json:"bu"`
	UpdateSteps  []UpdateStep `json:"steps"`
}

type Step struct {
	Plugin      string       `json:"plugin"`
	Events      []string     `json:"events"`
	EntryPoints []EntryPoint `json:"entry_point"`
}

type UpdateStep struct {
	Plugin            string             `json:"plugin"`
	Events            []string           `json:"events"`
	UpdateEntryPoints []UpdateEntryPoint `json:"entry_point"`
}

type EntryPoint struct {
	Request  map[string]interface{} `json:"request"`
	Response []string               `json:"response"`
}
type UpdateEntryPoint struct {
	Request  map[string]interface{} `json:"request"`
	Response []string               `json:"response"`
}

type DeleteRequest struct {
	WorkflowID string `json:"workflow_id"`
}

func (uw *UpdateWorkflowRequest) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            uw.ID,
		"workflow_name": uw.WorkflowName,
		"bu":            uw.BU,
	}
}
