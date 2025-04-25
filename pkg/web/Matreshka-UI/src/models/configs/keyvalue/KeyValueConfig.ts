import { Component } from "vue";

import KeyValueConfigView from "@/components/config/keyvalue/KeyValueConfigView.vue";
import { Change } from "@/models/configs/Change.ts";
import ConfigContent from "@/models/configs/ConfigContent.ts";

// TODO Implement
export default class KeyValueConfigContent implements ConfigContent {
  getChanges(): Change[] {
    return [];
  }
  rollback(): void {}
  isChanged(): boolean {
    return false;
  }

  getComponent(): Component {
    return KeyValueConfigView;
  }
}
