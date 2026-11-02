package permissions

import "tholian-core/protocol/id"

const (
	PurposeInvestigation = "investigation"
	PurposeIncident      = "incident"
	PurposeResponse      = "response"
	PurposeMitigation    = "mitigation"
	PurposeTask          = "task"
)

type Purpose struct {
	ID          string `json:"id"`
	Kind        string `json:"kind"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Incident    string `json:"incident,omitempty"`
}

func NewPurpose(kind string, title string) Purpose {

	return Purpose{
		ID:          id.NewID(),
		Kind:        kind,
		Title:       title,
		Description: "",
		Incident:    "",
	}

}

func (purpose *Purpose) IsValid() bool {

	if purpose == nil {
		return false
	}

	if purpose.ID != "" && purpose.Kind != "" && purpose.Title != "" {

		switch purpose.Kind {
		case PurposeInvestigation,
			PurposeIncident,
			PurposeResponse,
			PurposeMitigation,
			PurposeTask:
			return true
		}

	}

	return false

}
