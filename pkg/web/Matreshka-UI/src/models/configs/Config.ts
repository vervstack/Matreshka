import { Component } from "vue";

import ConfigBase from "@/models/configs/ConfigBase.ts";
import ConfigContent from "@/models/configs/ConfigContent.ts";
import KeyValueConfig from "@/models/configs/keyvalue/KeyValueConfig.ts";
import EnvNode from "@/models/shared/Node.ts";
import { PatchConfigPatch } from "@vervstack/matreshka";

export default class Config extends ConfigBase {
  content: ConfigContent;

  constructor(name: string) {
    super(name);

    this.content = new KeyValueConfig(new EnvNode("", ""));
  }

  rollback() {
    this.content.rollback();
  }

  getPatches(): PatchConfigPatch[] {
    return this.content.getChanges();
  }

  isChanged(): boolean {
    return this.content.isChanged();
  }

  getComponent(): Component {
    return this.content.getComponent();
  }
}
