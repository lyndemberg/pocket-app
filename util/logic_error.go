package util

//LogicError means that there was a problem processing due to some invalid input
type LogicError struct {
	Msg string
}

func (m LogicError) Error() string {
	return m.Msg
}
