package user

type ValidationError struct {
	InvalidField string
	Err          error
}

func (m *ValidationError) Error() string {
	return "Invalid Group Buy Item"
}
