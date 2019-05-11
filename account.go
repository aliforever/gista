package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type account struct {
	ig *instagram
}

func newAccount(ig *instagram) (a *account) {
	a = &account{ig: ig}
	return
}

func (a *account) Login(username, password string) (r *responses.Login, err error) {
	r = &responses.Login{}
	err = a.ig.client.Request(constants.Login).
		SetNeedsAuth(false).
		AddPost("country_codes", `[{"country_code":"1","source":["default","sim"]}]`).
		AddPhoneIdPost().
		AddCSRFPost().
		AddPost("username", username).
		AddPost("password", password).
		AddAdIdPost().
		AddGuIdPost().
		AddDeviceIdPost().
		AddPost("google_tokens", "[]").
		AddPost("login_attempt_count", "0").
		GetResponse(r)
	return
}

func (a *account) SetContactPointPreFill(usage string) (r *responses.Generic, err error) {
	r = &responses.Generic{}
	err = a.ig.client.Request(constants.ContactPointPreFill).
		SetNeedsAuth(false).
		AddPhoneIdPost().
		AddCSRFPost().
		AddPost("usage", usage).GetResponse(r)
	return
}
