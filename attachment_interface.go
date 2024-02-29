package berr

type Attachments []Attachment

type Attachment interface {
	Key() string
	Value() any
	Type() string
	Sensitive() bool
}
