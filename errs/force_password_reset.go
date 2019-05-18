package errs

type ForcedPasswordReset GistaError

func (fpr ForcedPasswordReset) Error() string {
	m := "unknown"
	if fpr.Message != nil {
		m = *fpr.Message
	}
	return m
}
