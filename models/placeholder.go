package models

type Placeholder struct {
	IsLinked bool   `json:"is_linked"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}
