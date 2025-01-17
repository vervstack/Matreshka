package domain

type PatchConfigRequest struct {
	ServiceName string
	Batch       []PatchConfig
}

type PatchConfig struct {
	FieldName  string
	FieldValue *string
}

type RenameRequest struct {
	OldName string
	NewName string
}
