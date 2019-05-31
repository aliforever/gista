package errs

import (
	"github.com/aliforever/gista/models"
	"github.com/aliforever/gista/responses"
)

type ChallengeRequired struct {
	Type         *string
	Message      *string
	HTTPResponse responses.ResponseInterface
}

func (cr ChallengeRequired) Error() string {
	m := "unknown"
	if cr.Message != nil {
		m = *cr.Message
	}
	return m
}

func (cr ChallengeRequired) GetChallenge() *models.Challenge {
	obj := cr.HTTPResponse.(interface{})
	checkpointUrl := obj.(*responses.Login).Challenge
	return checkpointUrl
}
