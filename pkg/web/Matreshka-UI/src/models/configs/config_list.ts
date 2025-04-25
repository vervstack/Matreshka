import {ConfigBase} from "@/models/configs/config.ts";

export class CfgList {
    configInfo: ConfigBase[]
    total: number

    constructor(configInfo: ConfigBase[], total: number) {
        this.configInfo = configInfo;
        this.total = total;
    }
}
