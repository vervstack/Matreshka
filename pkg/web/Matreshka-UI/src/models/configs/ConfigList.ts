import ConfigBase from "@/models/configs/ConfigBase.ts";

export default class ConfigList {
  configInfo: ConfigBase[];
  total: number;

  constructor(configInfo: ConfigBase[], total: number) {
    this.configInfo = configInfo;
    this.total = total;
  }
}
