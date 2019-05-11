package responses

import "github.com/aliforever/gista/models"

type Sync struct {
	Response
	//Experiments []map[string]interface{} `json:"experiments"`
	Experiments []models.Experiment `json:"experiments"`
}
