package responses

type LauncherSync struct {
	Response
	Configs map[string]interface{} `json:"configs"`
}
