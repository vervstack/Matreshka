import { Component } from "vue";

import EnvNode from "@/models/shared/Node";

import KeyValueConfigView from "@/components/config/keyvalue/KeyValueConfigView.vue";
import { Change } from "@/models/configs/Change.ts";
import ConfigContent from "@/models/configs/ConfigContent.ts";
import { ConfigValue } from "@/models/shared/Values.ts";

const objectSeparator = "_";

export default class KeyValueConfig implements ConfigContent {
  configValue: ConfigValue<string>;
  children: KeyValueConfig[] = [];

  constructor(root: EnvNode) {
    this.configValue = new ConfigValue<string>(root.name, root.value);

    root.children.map((n: EnvNode) => {
      this.children.push(new KeyValueConfig(n));
    });
  }

  getChanges(): Change[] {
    const changes = this.configValue.getChanges();

    this.children.map((c: KeyValueConfig) => {
        const childChanges = c.getChanges();

        childChanges.map((c: Change, idx: number) => {
          if (!this.isRoot()) {
            childChanges[idx].envName = this.configValue.envName + objectSeparator + c.envName;
          }
        });

        changes.push(...childChanges);
      },
    );

    return changes;
  }

  rollback(): void {
    this.configValue.rollback();

    if (this.children.find(v => v.configValue.isNew)) {
      this.children = [];
    }
  }

  isChanged(): boolean {
    return this.configValue.isChanged() ||
      this.children.find(v => v.isChanged()) !== undefined ||
      this.configValue.isNew;
  }

  getComponent(): Component {
    return KeyValueConfigView;
  }

  isRoot(): boolean {
    return this.configValue.envName == "";
  }

  countChildren(): number {
    let basicLength = this.children.length;
    this.children.forEach((c: KeyValueConfig) => {
      basicLength += c.countChildren();
    });

    return basicLength;
  }

  mute() {
    this.configValue.isMuted = true;
    this.children.forEach((c: KeyValueConfig) => {
      c.mute();
    });
  }

  unmute() {
    this.configValue.isMuted = false;
    this.children.forEach((c: KeyValueConfig) => {
      c.unmute();
    });
  }
}
