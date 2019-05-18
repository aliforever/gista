package errs

type CheckpointRequired GistaError

func (ck CheckpointRequired) Error() string {
	m := "unknown"
	if ck.Message != nil {
		m = *ck.Message
	}
	return m
}
