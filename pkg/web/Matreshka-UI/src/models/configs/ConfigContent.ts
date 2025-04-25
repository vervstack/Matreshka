import { Component } from "vue";

import { Change } from "@/models/configs/Change.ts";

export default interface ConfigContent {
  getChanges(): Change[];
  rollback(): void;

  isChanged(): boolean;

  getComponent(): Component;
}
