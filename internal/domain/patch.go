package domain

type PatchConfigRequest struct {
	ServiceName string
	Batch       []PatchConfig
}

type PatchConfig struct {
	FieldPath  string
	FieldValue any
}
