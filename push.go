package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/errs"
	"github.com/aliforever/gista/responses"
)

type push struct {
	ig *Instagram
}

func newPush(i *Instagram) *push {
	return &push{ig: i}
}

func (p *push) Register(pushChannel, token string) (res *responses.PushRegister, err error) {
	if pushChannel != "mqtt" && pushChannel != "gcm" {
		err = errs.BadPushChannel(pushChannel)
		return
	}
	res = &responses.PushRegister{}
	dt := "android_gcm"
	if pushChannel == "mqtt" {
		dt = "android_mqtt"
	}
	mainPushChannel := "false"
	if pushChannel == "mqtt" {
		mainPushChannel = "true"
	}
	err = p.ig.client.Request(constants.PushRegister).
		SetSignedPost(false).
		AddPost("device_type", dt).
		AddPost("is_main_push_channel", mainPushChannel).
		AddPhoneIdPost().
		AddPost("device_token", token).
		AddCSRFPost().
		AddGuIdPost().
		AddUuIdPost().
		AddPost("users", *p.ig.AccountId).
		GetResponse(res)
	return
}
