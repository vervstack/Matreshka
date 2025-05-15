import { Component } from "vue";

import { PatchConfigPatch } from "@vervstack/matreshka";


export default interface ConfigContent {
  getChanges(): PatchConfigPatch[];
  rollback(): void;

  isChanged(): boolean;

  getComponent(): Component;
}
