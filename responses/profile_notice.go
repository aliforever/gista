package responses

type ProfileNotice struct {
	Response
	HasChangePasswordMegaphone bool `json:"has_change_password_megaphone"`
}
