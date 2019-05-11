package responses

import "github.com/aliforever/gista/models"

type Login struct {
	Response
	/*Pk                         int              `json:"pk"`
	Username                   string           `json:"username"`
	HasAnonymousProfilePicture bool             `json:"has_anonymous_profile_picture"`
	ProfilePicUrl              string           `json:"profile_pic_url"`
	ProfilePicId               string           `json:"profile_pic_id"`
	FullName                   string           `json:"full_name"`
	IsPrivate                  bool             `json:"is_private"`
	IsVerified                 bool             `json:"is_verified"`
	AllowedCommenterType       string           `json:"allowed_commenter_type"`
	ReelAutoArchive            string           `json:"reel_auto_archive"`
	AllowContactsSync          bool             `json:"allow_contacts_sync"`
	PhoneNumber                string           `json:"phone_number"`
	CountryCode                int              `json:"country_code"`
	NationalNumber             int              `json:"national_number"`*/
	ErrorTitle         interface{}      `json:"error_title"`
	ErrorType          interface{}      `json:"error_type"`
	Buttons            interface{}      `json:"buttons"`
	InvalidCredentials interface{}      `json:"invalid_credentials"`
	LoggedInUser       models.User      `json:"logged_in_user"`
	TwoFactorRequired  interface{}      `json:"two_factor_required"`
	CheckPointUrl      string           `json:"check_point_url"`
	Lock               interface{}      `json:"lock"`
	HelpUrl            string           `json:"help_url"`
	Challenge          models.Challenge `json:"challenge"`
}
