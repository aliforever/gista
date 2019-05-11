package models

type Aymf struct {
	Items         []AymfItem  `json:"items"`
	MoreAvailable interface{} `json:"more_available"`
}
