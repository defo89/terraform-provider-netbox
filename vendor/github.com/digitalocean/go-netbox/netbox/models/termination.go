package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Termination connected_endpoint
// swagger:model ConnectedEndpoint
type Termination struct {
	ID      uint64   `json:"id,omitempty"`
	URL     string   `json:"url,omitempty"`
	Circuit *Circuit `json:"circuit,omitempty"`
	Name    string   `json:"name,omitempty"`
	Cable   uint64   `json:"cable,omitempty"`
}
