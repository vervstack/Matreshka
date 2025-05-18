package domain

import (
	"go.redsock.ru/evon"
)

type PatchConfigRequest struct {
	ConfigName    ConfigName
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

type ReplaceConfig struct {
	Name    string
	Version string
	Config  []*evon.Node
}
