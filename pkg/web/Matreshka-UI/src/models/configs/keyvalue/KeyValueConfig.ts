import { Component } from "vue";

import EnvNode from "@/models/shared/Node";

import KeyValueConfigView from "@/components/config/keyvalue/KeyValueConfigView.vue";
import { Change } from "@/models/configs/Change.ts";
import ConfigContent from "@/models/configs/ConfigContent.ts";
import { ConfigValue } from "@/models/shared/Values.ts";

// TODO Implement
export default class KeyValueConfig implements ConfigContent {
  configValue: ConfigValue<string>;
  children: KeyValueConfig[] = [];

  constructor(root: EnvNode) {
    this.configValue = new ConfigValue<string>(root.name, root.value, root.children.length > 0);

    root.children.map((n: EnvNode) => {
      this.children.push(new KeyValueConfig(n));
    });
  }

  getChanges(): Change[] {
    return [];
  }

  rollback(): void {
  }

  isChanged(): boolean {
    return this.configValue?.isChanged() || false
      || this.children.find(v => v.isChanged()) !== undefined;
  }

  getComponent(): Component {
    return KeyValueConfigView;
  }
}
