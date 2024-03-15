package berr

const AttachmentDetailType = "detail"

type detailAttachment struct {
	value any
	key   string
}

func (d detailAttachment) Key() string {
	return d.key
}

func (d detailAttachment) Value() any {
	return d.value
}

func (d detailAttachment) Type() string {
	return AttachmentDetailType
}

func (d detailAttachment) Sensitive() bool {
	return false
}
