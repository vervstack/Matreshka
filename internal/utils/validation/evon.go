package validation

import (
	"strings"

	"go.redsock.ru/evon"
	"go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/domain"
)

func (v Validator) AsEvon(original evon.NodeStorage, patch *domain.PatchConfigRequest) (err error) {
	for i := range patch.Update {
		patch.Update[i].FieldName, err = v.normalizeAndValidateEnvName(patch.Update[i].FieldName)
		if err != nil {
			return rerrors.Wrap(err)
		}
	}

	for i := range patch.RenameTo {
		originalNode := original[patch.RenameTo[i].OldName]
		if originalNode == nil {
			continue
		}

		patch.RenameTo[i].OldName, err = v.normalizeAndValidateEnvName(patch.RenameTo[i].OldName)
		if err != nil {
			return rerrors.Wrap(err)
		}

		patch.RenameTo[i].NewName, err = v.normalizeAndValidateEnvName(patch.RenameTo[i].NewName)
		if err != nil {
			return rerrors.Wrap(err)
		}

		patch.RenameTo = append(patch.RenameTo,
			walkAndRename(originalNode.InnerNodes, patch.RenameTo[i].OldName, patch.RenameTo[i].NewName)...)
	}

	for i := range patch.Delete {
		originalNode := original[patch.Delete[i]]
		if originalNode == nil {
			continue
		}

		patch.Delete[i], err = v.normalizeAndValidateEnvName(patch.Delete[i])
		if err != nil {
			return rerrors.Wrap(err)
		}

		deletedChildren := childrenAsPlainSlice(originalNode)
		for _, c := range deletedChildren {
			patch.Delete = append(patch.Delete, c.Name)
		}
	}
	return nil
}

func walkAndRename(children []*evon.Node, oldName, newName string) []domain.PatchRename {
	out := make([]domain.PatchRename, 0, len(children))
	for _, child := range children {
		out = append(out, domain.PatchRename{
			OldName: child.Name,
			NewName: replaceAtStart(child.Name, oldName, newName),
		})
		out = append(out, walkAndRename(child.InnerNodes, oldName, newName)...)
	}

	return out
}

func childrenAsPlainSlice(root *evon.Node) []*evon.Node {
	out := make([]*evon.Node, 0, len(root.InnerNodes))
	for _, child := range root.InnerNodes {
		out = append(out, child)
		out = append(out, childrenAsPlainSlice(child)...)
	}

	return out
}

func replaceAtStart(str, old, new string) string {
	if !strings.HasPrefix(str, old) {
		return str
	}

	return new + str[len(old):]

}
