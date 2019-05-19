package gista

import (
	"encoding/json"
	"regexp"

	"github.com/aliforever/gista/errs"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type account struct {
	ig *Instagram
}

func newAccount(ig *Instagram) (a *account) {
	a = &account{ig: ig}
	return
}

func (a *account) login(username, password string) (r *responses.Login, err error) {
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

func (a *account) finishTwoFactorLogin(username, password, twoFactorIdentifier, verificationCode, verificationMethod string, appRefreshInterval int, usernameHandler *string) (r *responses.Login, err error) {
	if username == "" || password == "" {
		err = errs.UsernameAndPasswordRequiredForTwoFactor
		return
	}
	methodGiven := false
	for _, method := range []string{"1", "2", "3"} {
		if method == verificationMethod {
			methodGiven = true
		}
	}
	if !methodGiven {
		err = errs.InvalidMethodForTwoFactor(verificationMethod)
		return
	}
	if verificationCode == "" {
		err = errs.VerificationCodeRequiredForTwoFactor
		return
	}
	if a.ig.Username != username || a.ig.Password != password {
		a.ig.setUser(username, password)
	}
	if usernameHandler != nil {
		username = *usernameHandler
	}
	rgx := regexp.MustCompile(`\s+`)
	verificationCode = rgx.ReplaceAllString(verificationCode, "")
	//pretty.Println(username, verificationMethod, verificationCode, twoFactorIdentifier)
	r = &responses.Login{}
	err = a.ig.client.Request(constants.TwoFactorLogin).
		SetNeedsAuth(false).
		AddPost("verification_method", verificationMethod).
		AddPost("verification_code", verificationCode).
		AddPost("two_factor_identifier", twoFactorIdentifier).
		AddPost("username", username).
		AddCSRFPost().
		AddGuIdPost().
		AddDeviceIdPost().
		GetResponse(r)
	err = a.ig.updateLoginState(r)
	if err != nil {
		return
	}
	err = a.ig.sendLoginFlow(true, appRefreshInterval)
	if err != nil {
		return
	}
	return
}

func (a *account) GetCurrentUser() (r *responses.UserInfo, err error) {
	r = &responses.UserInfo{}
	err = a.ig.client.Request(constants.CurrentUser).
		AddParam("edit", "true").
		GetResponse(r)
	return
}

func (a *account) ChangeProfilePicture(photoFileName string) (r *responses.UserInfo, err error) {
	r = &responses.UserInfo{}

	req := a.ig.client.Request(constants.ChangeProfilePicture).
		AddUuIdPost().
		AddUIdPost().
		AddCSRFPost()
	fileName := "profile_pic"
	req, err = req.AddFile("profile_pic", photoFileName, &fileName, nil)
	if err != nil {
		return
	}
	err = req.GetResponse(r)
	return
}

func (a *account) RemoveProfilePicture() (r *responses.UserInfo, err error) {
	r = &responses.UserInfo{}
	err = a.ig.client.Request(constants.RemoveProfilePicture).
		AddUuIdPost().
		AddUIdPost().
		AddCSRFPost().
		GetResponse(r)
	return
}

func (a *account) SetPublic() (r *responses.UserInfo, err error) {
	r = &responses.UserInfo{}
	err = a.ig.client.Request(constants.SetPublic).
		AddUuIdPost().
		AddUIdPost().
		AddCSRFPost().
		GetResponse(r)
	return
}

func (a *account) SetPrivate() (r *responses.UserInfo, err error) {
	r = &responses.UserInfo{}
	err = a.ig.client.Request(constants.SetPrivate).
		AddUuIdPost().
		AddUIdPost().
		AddCSRFPost().
		GetResponse(r)
	return
}

func (a *account) SetBusinessInfo(phoneNumber, email, categoryId string) (r *responses.UserInfo, err error) {
	r = &responses.UserInfo{}
	ppc, _ := json.Marshal(map[string]string{
		"public_phone_number":     phoneNumber,
		"business_contact_method": "CALL",
	})
	err = a.ig.client.Request(constants.CreateBusinessInfo).
		AddPost("set_public", "true").
		AddPost("entry_point", "setting").
		AddPost("public_phone_contact", string(ppc)).
		AddPost("public_email", email).
		AddPost("category_id", categoryId).
		AddUuIdPost().
		AddUIdPost().
		AddCSRFPost().
		GetResponse(r)
	return
}

func (a *account) SwitchToBusinessProfile() (r *responses.UserInfo, err error) {
	r = &responses.UserInfo{}
	err = a.ig.client.Request(constants.SwitchToBusinessProfile).
		GetResponse(r)
	return
}

func (a *account) CheckUsername(username string) (r *responses.CheckUsername, err error) {
	r = &responses.CheckUsername{}
	err = a.ig.client.Request(constants.CheckUsername).
		AddUIdPost().
		AddUuIdPost().
		AddCSRFPost().
		AddPost("username", username).
		GetResponse(r)
	return
}

func (a *account) SwitchToPersonalProfile() (r *responses.SwitchPersonalProfile, err error) {
	r = &responses.SwitchPersonalProfile{}
	err = a.ig.client.Request(constants.SwitchToPersonalProfile).
		AddUuIdPost().
		AddUIdPost().
		AddCSRFPost().
		GetResponse(r)
	return
}

func (a *account) SetBiography(biography string) (r *responses.UserInfo, err error) {
	if len(biography) > 150 {
		err = errs.InvalidBiography(biography)
		return
	}
	r = &responses.UserInfo{}
	err = a.ig.client.Request(constants.SetBiography).
		AddPost("raw_text", biography).
		AddUuIdPost().
		AddUIdPost().
		AddDeviceIdPost().
		AddCSRFPost().
		GetResponse(r)
	return
}

func (a *account) EditProfile(url, phone, name, biography, email, gender string, newUsername *string) (r *responses.UserInfo, err error) {
	var currentUser *responses.UserInfo
	currentUser, err = a.GetCurrentUser()
	if err != nil {
		return
	}
	username := currentUser.User.Username
	if newUsername != nil && len(*newUsername) > 0 {
		username = *newUsername
	}
	r = &responses.UserInfo{}
	err = a.ig.client.Request(constants.EditProfile).
		AddPost("external_url", url).
		AddPost("phone_number", phone).
		AddPost("username", username).
		AddPost("first_name", name).
		AddPost("biography", biography).
		AddPost("email", email).
		AddPost("gender", gender).
		AddUuIdPost().
		AddUIdPost().
		AddDeviceIdPost().
		AddCSRFPost().
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
