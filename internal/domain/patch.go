package domain

type PatchConfigRequest struct {
	ConfigName    string
	ConfigVersion string
	Update        []PatchUpdate
	RenameTo      []PatchRename
	Delete        []string
}

type PatchUpdate struct {
	FieldName  string
	FieldValue string
}

type PatchRename struct {
	OldName string
	NewName string
}
