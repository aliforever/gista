package models

type ImageVersions2 struct {
	Candidates []ImageCandidate `json:"candidates"`
	TraceToken interface{}      `json:"trace_token"`
}
