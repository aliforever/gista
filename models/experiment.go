package models

type Experiment struct {
	Name             *string  `json:"name"`
	Group            string   `json:"group,omitempty"`
	AdditionalParams []string `json:"additional_params"`
	Params           *[]Param `json:"params,omitempty"`
	LoggingId        string   `json:"logging_id"`
	Expired          bool     `json:"expired"`
}
