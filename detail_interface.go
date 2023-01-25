package berr

type Details []Detail

type Detail interface {
	Key() string
	Value() any
	Type() string
}
