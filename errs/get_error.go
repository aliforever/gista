package errs

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/aliforever/gista/responses"
)

var exceptionMap = map[string][]string{
	"LoginRequired":       {"login_required"},
	"CheckpointRequired":  {"checkpoint_required", "checkpoint_challenge_required"},
	"ChallengeRequired":   {"challenge_required"},
	"FeedbackRequired":    {"feedback_required"},
	"ConsentRequired":     {"consent_required"},
	"IncorrectPassword":   {"/password(.*?)incorrect/", "bad_password"},
	"InvalidSmsCode":      {"/check(.*?)security(.*?)code/", "sms_code_validation_code_invalid"},
	"AccountDisabled":     {"/account(.*?)disabled(.*?)violating/"},
	"SentryBlock":         {"sentry_block"},
	"InvalidUser":         {"/username(.*?)doesn't(.*?)belong/", "invalid_user"},
	"ForcedPasswordReset": {"/reset(.*?)password/"},
}

var errorMap = map[string]func(t, m *string, resp *http.Response) error{
	"LoginRequired": func(t, m *string, resp *http.Response) error {
		err := LoginRequired{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
	"CheckpointRequired": func(t, m *string, resp *http.Response) error {
		err := CheckpointRequired{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
	"ChallengeRequired": func(t, m *string, resp *http.Response) error {
		err := ChallengeRequired{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
	"FeedbackRequired": func(t, m *string, resp *http.Response) error {
		err := FeedbackRequired{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
	"ConsentRequired": func(t, m *string, resp *http.Response) error {
		err := ConsentRequired{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
	"IncorrectPassword": func(t, m *string, resp *http.Response) error {
		err := IncorrectPassword{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
	"InvalidSmsCode": func(t, m *string, resp *http.Response) error {
		err := InvalidSmsCode{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
	"AccountDisabled": func(t, m *string, resp *http.Response) error {
		err := AccountDisabled{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
	"SentryBlock": func(t, m *string, resp *http.Response) error {
		err := SentryBlock{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
	"InvalidUser": func(t, m *string, resp *http.Response) error {
		err := InvalidUser{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
	"ForcedPasswordReset": func(t, m *string, resp *http.Response) error {
		err := ForcedPasswordReset{Type: t, Message: m, HTTPResponse: resp}
		return err
	},
}

func GetError(prefix *string, serverMessage *string, response responses.ResponseInterface, httpResponse *http.Response) (err error) {
	var messages []string
	if serverMessage != nil {
		messages = append(messages, *serverMessage)
	}
	var serverErrorType *string = nil
	if response != nil {
		if t := (response).GetErrorType(); t != nil {
			serverErrorType = t
			messages = append(messages, *t)
		}
	}
	var exceptionClass *string = nil
Loop:
	for _, message := range messages {
		for className, patterns := range exceptionMap {
			for _, pattern := range patterns {
				if pattern[0] == '/' {
					r := regexp.MustCompile(strings.ReplaceAll(pattern, "/", ""))
					if r.MatchString(message) {
						exceptionClass = &className
						break Loop
					}
				} else {
					if strings.Contains(message, pattern) {
						exceptionClass = &className
						break Loop
					}
				}
			}
		}
	}

	if exceptionClass == nil {
		var httpStatusCode = 0
		if httpResponse != nil {
			httpStatusCode = httpResponse.StatusCode
		}
		switch httpStatusCode {
		case 400:
			s := "BadRequest"
			exceptionClass = &s
		case 404:
			s := "NotFound"
			exceptionClass = &s
		default:
			s := "Endpoint"
			exceptionClass = &s
		}
	}
	var displayMessage *string = nil
	if serverMessage != nil && len(*serverMessage) > 0 {
		displayMessage = serverMessage
	} else {
		displayMessage = serverErrorType
	}
	if displayMessage != nil {
		prettified := prettifyMessage(*displayMessage)
		displayMessage = &prettified
		if prefix != nil {
			m := fmt.Sprintf("%s: %s", *prefix, *displayMessage)
			displayMessage = &m
		}
	}
	e := errorMap[*exceptionClass](serverErrorType, displayMessage, httpResponse)
	err = e
	return
}

func prettifyMessage(message string) string {
	lastChar := message[len(message)-1]
	if lastChar != '.' && lastChar != '!' && lastChar != '?' {
		message += "."
	}
	message = strings.Title(message)
	message = strings.ReplaceAll(message, "_", " ")
	return message
}
