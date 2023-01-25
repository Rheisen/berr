package berr

type detail struct {
	key   string
	value any
}

func (d detail) Key() string {
	return d.key
}

func (d detail) Value() any {
	return d.value
}

func (d detail) Type() string {
	return "berr_detail"
}
