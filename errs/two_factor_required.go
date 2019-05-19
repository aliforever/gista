package errs

import (
	"github.com/aliforever/gista/models"
	"github.com/aliforever/gista/responses"
)

type TwoFactorRequired struct {
	Type         *string
	Message      *string
	HTTPResponse responses.ResponseInterface
}

func (tfr TwoFactorRequired) Error() string {
	m := "unknown"
	if tfr.Message != nil {
		m = *tfr.Message
	}
	return m
}

func (tfr TwoFactorRequired) GetTwoFactorInfo() *models.TwoFactorInfo {
	obj := tfr.HTTPResponse.(interface{})
	twoFactorInfo := obj.(*responses.Login).TwoFactorInfo
	return twoFactorInfo
}
