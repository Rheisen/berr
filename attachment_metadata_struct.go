package berr

type metadataDetail struct {
	key   string
	value any
}

func (d metadataDetail) Key() string {
	return d.key
}

func (d metadataDetail) Value() any {
	return d.value
}

func (d metadataDetail) Type() string {
	return "berr_metadata_detail"
}

func (d metadataDetail) Sensitive() bool {
	return true
}
