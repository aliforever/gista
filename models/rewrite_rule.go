package models

type RewriteRule struct {
	Matcher  string `json:"matcher"`
	Replacer string `json:"replacer"`
}
