import { PatchConfigPatch } from "@vervstack/matreshka";

export function newRenamePatch(fieldName: string, newName: string): PatchConfigPatch {
  return {
    fieldName: fieldName,
    rename: newName,
  } as PatchConfigPatch;
}

export function newUpdatePatch(fieldName: string, newValue: string): PatchConfigPatch {
  return {
    fieldName: fieldName,
    updateValue: newValue,
  } as PatchConfigPatch;
}


export function newDeletePatch(fieldName: string): PatchConfigPatch {
  return {
    fieldName: fieldName,
    delete: true,
  } as PatchConfigPatch;
}

