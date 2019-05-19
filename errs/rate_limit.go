package errs

type RateLimit GistaError

func (rl RateLimit) Error() string {
	m := "unknown"
	if rl.Message != nil {
		m = *rl.Message
	}
	return m
}
