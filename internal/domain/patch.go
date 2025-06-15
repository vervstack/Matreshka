package domain

import (
	"go.redsock.ru/evon"
)

type PatchConfigRequest struct {
	ConfigName    ConfigName
	ConfigVersion string
	Upsert        []PatchUpdate
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

type ReplaceConfigReq struct {
	Name    ConfigName
	Version string
	Config  *evon.Node
}
