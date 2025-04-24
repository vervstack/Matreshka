import {ConfigTypePrefix} from "@vervstack/matreshka";
import {ConfigContent} from "@/models/configs/configContent.ts";
import {Change} from "@/models/configs/verv/info/AppInfo.ts";
import {KeyValueConfigContent} from "@/models/configs/keyvalue/config.ts";
import {Component} from "vue";

const defaultVersion = 'master'

export class Config {
    type: ConfigTypePrefix;
    name: string;

    versions: string[] = [defaultVersion];
    selectedVersion: string = defaultVersion

    content: ConfigContent;

    constructor(name: string) {
        this.type = extractType(name)
        if (this.type !== ConfigTypePrefix.unknown) {
            name = name.substring(this.type.length + 1)
        }

        this.name = name
        this.content = new KeyValueConfigContent()
    }

    rollback() {
        this.content.rollback()
    }

    getChanges(): Change[] {
        return this.content.getChanges()
    }

    isChanged(): boolean {
        return this.content.isChanged()
    }

    getComponent() : Component {
        const com = this.content.getComponent()
        console.log(com)
        return com
    }

}

const supportedTypes: ConfigTypePrefix[] = [
    ConfigTypePrefix.verv,
    ConfigTypePrefix.minio,
    ConfigTypePrefix.pg,
    ConfigTypePrefix.nginx,
]

function extractType(configName: string): ConfigTypePrefix {
    const foundType = supportedTypes
        .find(
            (typePrefix: ConfigTypePrefix) => configName.startsWith(typePrefix))

    if (foundType) {
        return foundType;
    }

    return ConfigTypePrefix.unknown
}
