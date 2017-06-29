package exception

type NullReference struct {
	message string
}

func (e *NullReference) Error() string {
	return e.message
}
