package server_message_thrower

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/aliforever/gista/responses"
)

var exceptionMap = map[string][]string{
	"LoginRequiredException":       {"login_required"},
	"CheckpointRequiredException":  {"checkpoint_required", "checkpoint_challenge_required"},
	"ChallengeRequiredException":   {"challenge_required"},
	"FeedbackRequiredException":    {"feedback_required"},
	"ConsentRequiredException":     {"consent_required"},
	"IncorrectPasswordException":   {"/password(.*?)incorrect/", "bad_password"},
	"InvalidSmsCodeException":      {"/check(.*?)security(.*?)code/", "sms_code_validation_code_invalid"},
	"AccountDisabledException":     {"/account(.*?)disabled(.*?)violating/"},
	"SentryBlockException":         {"sentry_block"},
	"InvalidUserException":         {"/username(.*?)doesn't(.*?)belong/", "invalid_user"},
	"ForcedPasswordResetException": {"/reset(.*?)password/"},
}

type ServerMessageThrower struct {
}

func AutoError(prefix *string, serverMessage *string, response *responses.ResponseInterface, httpResponse *http.Response) (err error) {
	var messages []string
	if serverMessage != nil {
		messages = append(messages, *serverMessage)
	}
	var serverErrorType *string = nil
	if response != nil {
		if t := (*response).GetErrorType(); t != nil {
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
					r := regexp.MustCompile(pattern)
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
		var httpStatusCode int = 0
		if httpResponse != nil {
			httpStatusCode = httpResponse.StatusCode
		}
		switch httpStatusCode {
		case 400:
			s := "BadRequestException"
			exceptionClass = &s
		case 404:
			s := "NotFoundException"
			exceptionClass = &s
		default:
			s := "EndpointException"
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
		prettified := PrettifyMessage(*displayMessage)
		displayMessage = &prettified
	}
	//if serverRe
	return
}

func PrettifyMessage(message string) string {
	lastChar := message[len(message)-1]
	if lastChar != '.' && lastChar != '!' && lastChar != '?' {
		message += "."
	}
	message = strings.Title(message)
	message = strings.ReplaceAll(message, "_", " ")
	return message
}
