import { Component } from "vue";

import { Change } from "@/models/configs/Change.ts";
import ConfigBase from "@/models/configs/ConfigBase.ts";
import ConfigContent from "@/models/configs/ConfigContent.ts";
import KeyValueConfigContent from "@/models/configs/keyvalue/KeyValueConfig.ts";

export default class Config extends ConfigBase {
  content: ConfigContent;

  constructor(name: string) {
    super(name);

    this.content = new KeyValueConfigContent();
  }

  rollback() {
    this.content.rollback();
  }

  getChanges(): Change[] {
    return this.content.getChanges();
  }

  isChanged(): boolean {
    return this.content.isChanged();
  }

  getComponent(): Component {
    const com = this.content.getComponent();
    console.log(com);
    return com;
  }
}
