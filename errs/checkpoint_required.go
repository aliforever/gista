package errs

import (
	"github.com/aliforever/gista/responses"
)

type CheckpointRequired struct {
	Type         *string
	Message      *string
	HTTPResponse responses.ResponseInterface
}

func (cp CheckpointRequired) Error() string {
	m := "unknown"
	if cp.Message != nil {
		m = *cp.Message
	}
	return m
}

func (cp CheckpointRequired) GetCheckpointUrl() string {
	obj := cp.HTTPResponse.(interface{})
	checkpointUrl := obj.(*responses.Login).CheckPointUrl
	return checkpointUrl
}
