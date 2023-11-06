package item

type InvalidItem struct {
	InvalidField string
	Err          error
}

func (e *InvalidItem) Error() string {
	return "Invalid Group Buy Item field: " + e.InvalidField
}
