package berr

const AttachmentMetadataType = "metadata"

type metadataAttachment struct {
	value any
	key   string
}

func (d metadataAttachment) Key() string {
	return d.key
}

func (d metadataAttachment) Value() any {
	return d.value
}

func (d metadataAttachment) Type() string {
	return AttachmentMetadataType
}

func (d metadataAttachment) Sensitive() bool {
	return true
}
