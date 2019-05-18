package errs

type SentryBlock GistaError

func (sb SentryBlock) Error() string {
	m := "unknown"
	if sb.Message != nil {
		m = *sb.Message
	}
	return m
}
