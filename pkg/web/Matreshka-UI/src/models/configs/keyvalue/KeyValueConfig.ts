import { Component } from "vue";

import EnvNode from "@/models/shared/Node";

import KeyValueConfigView from "@/components/config/keyvalue/KeyValueConfigView.vue";
import ConfigContent from "@/models/configs/ConfigContent.ts";
import { ConfigValue } from "@/models/shared/Values.ts";
import { PatchConfigPatch } from "@vervstack/matreshka";

const objectSeparator = "_";

export default class KeyValueConfig implements ConfigContent {
  configValue: ConfigValue<string>;
  children: KeyValueConfig[] = [];
  isFolded: boolean = false;

  constructor(root: EnvNode) {
    this.configValue = new ConfigValue<string>(root.name, root.value);

    root.children.map((n: EnvNode) => {
      this.children.push(new KeyValueConfig(n));
    });
  }

  getChanges(): PatchConfigPatch[] {
    const changes = this.configValue.getChanges();

    this.children.map((c: KeyValueConfig) => {
        const childChanges = c.getChanges();

        childChanges.map((_: PatchConfigPatch, idx: number) => {
          if (!this.isRoot()) {
            childChanges[idx].fieldName =
              this.configValue.getOriginalName() +
              objectSeparator +
              childChanges[idx].fieldName;

            if (childChanges[idx].rename) {
              childChanges[idx].rename =
                this.configValue.envName +
                objectSeparator +
                childChanges[idx].rename;
            }
          }
        });

        changes.push(...childChanges);
      },
    );

    return changes;
  }

  rollback(): void {
    this.configValue.rollback();

    this.children = this.children.filter(v => !v.configValue.isNew);
    this.children.map(v => {
      v.rollback();
    });
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
    if (this.isFolded) {
      return 0;
    }

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
