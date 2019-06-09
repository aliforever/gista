package errs

type UserNotFound GistaError

func (unf UserNotFound) Error() string {
	m := "unknown"
	if unf.Message != nil {
		m = *unf.Message
	}
	return m
}
