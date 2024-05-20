package berr

const AttachmentDetailType = "detail"
const AttachmentMetadataType = "metadata"
const AttachmentErrorType = "error"

type Attachment struct {
	attachmentType string
	key            string
	value          any
	sensitive      bool
}

func (a Attachment) Key() string {
	return a.key
}

func (a Attachment) Value() any {
	return a.value
}

func (a Attachment) Type() string {
	return a.attachmentType
}

func (a Attachment) Sensitive() bool {
	return a.sensitive
}
