package domain

type PatchConfigRequest struct {
	ServiceName string
	Batch       []PatchConfig
}

type PatchConfig struct {
	FieldName  string
	FieldValue any
}
