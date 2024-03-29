package berr

const AttachmentErrorType = "error"

type errorAttachment struct {
	value error
	key   string
}

func (d errorAttachment) Key() string {
	return d.key
}

func (d errorAttachment) Value() any {
	return d.value
}

func (d errorAttachment) Type() string {
	return AttachmentErrorType
}

func (d errorAttachment) Sensitive() bool {
	return true
}
