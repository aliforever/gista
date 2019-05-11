package zero_rating

var DefaultRewrite = map[string]string{
	`^(https?:\/\/)(i)(\.instagram\.com/.*)$`: "$1b.$2$3",
}
