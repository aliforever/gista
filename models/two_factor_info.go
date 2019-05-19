package models

type TwoFactorInfo struct {
	Username                  string                     `json:"username,omitempty"`
	SmsTwoFactorOn            bool                       `json:"sms_two_factor_on,omitempty"`
	TotpTwoFactorOn           bool                       `json:"totp_two_factor_on,omitempty"`
	ObfuscatedPhoneNumber     string                     `json:"obfuscated_phone_number,omitempty"`
	TwoFactorIdentifier       string                     `json:"two_factor_identifier,omitempty"`
	ShowMessengerCodeOption   bool                       `json:"show_messenger_code_option,omitempty"`
	ShowNewLoginScreen        bool                       `json:"show_new_login_screen,omitempty"`
	ShowTrustedDeviceOption   bool                       `json:"show_trusted_device_option,omitempty"`
	PhoneVerificationSettings *PhoneVerificationSettings `json:"phone_verification_settings,omitempty"`
}
