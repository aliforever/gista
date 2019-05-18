package errs

type FeedbackRequired GistaError

func (fr FeedbackRequired) Error() string {
	m := "unknown"
	if fr.Message != nil {
		m = *fr.Message
	}
	return m
}
