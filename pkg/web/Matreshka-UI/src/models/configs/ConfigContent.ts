import { Component } from "vue";

import { PatchConfigPatch } from "@vervstack/matreshka/matreshka-be_api.pb.ts";


export default interface ConfigContent {
  getChanges(): PatchConfigPatch[];
  rollback(): void;

  isChanged(): boolean;

  getComponent(): Component;
}
