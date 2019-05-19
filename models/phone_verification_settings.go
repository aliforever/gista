package models

type PhoneVerificationSettings struct {
	MaxSmsCount              int  `json:"max_sms_count"`
	ResendSmsDelaySec        int  `json:"resend_sms_delay_sec"`
	RobocallCountDownTimeSec int  `json:"robocall_count_down_time_sec"`
	RobocallAfterMaxSms      bool `json:"robocall_after_max_sms"`
}
