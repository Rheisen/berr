package berr

const AttachmentMetadataType = "metadata"

type metadataAttachment struct {
	key   string
	value any
}

func (d metadataAttachment) Key() string {
	return d.key
}

func (d metadataAttachment) Value() any {
	return d.value
}

func (d metadataAttachment) Type() string {
	return "berr_metadata_detail"
}

func (d metadataAttachment) Sensitive() bool {
	return true
}
