package berr

type errorDetail struct {
	key   string
	value error
}

func (d errorDetail) Key() string {
	return d.key
}

func (d errorDetail) Value() any {
	return d.value
}

func (d errorDetail) Type() string {
	return "berr_error_detail"
}

func (d errorDetail) Sensitive() bool {
	return true
}
