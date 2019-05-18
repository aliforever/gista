package errs

type ChallengeRequired GistaError

func (cr ChallengeRequired) Error() string {
	m := "unknown"
	if cr.Message != nil {
		m = *cr.Message
	}
	return m
}
