import {Component} from "vue";
import {ConfigTypePrefix} from "@vervstack/matreshka";
import {Config_content} from "@/models/configs/config_content.ts";
import {Change} from "@/models/configs/verv/info/VervConfig.ts";
import {KeyValueConfigContent} from "@/models/configs/keyvalue/config.ts";

const defaultVersion = 'master'

export class ConfigBase {
    type: ConfigTypePrefix;
    name: string;

    updated_at?: Date

    versions: string[] = [defaultVersion];
    selectedVersion: string = defaultVersion;

    constructor(name: string) {
        this.type = extractType(name)
        if (this.type !== ConfigTypePrefix.unknown) {
            name = name.substring(this.type.length + 1)
        }

        this.name = name
    }

    getMatreshkaName(): string {
        return this.type + '_' + this.name
    }
}

export class Config extends ConfigBase {
    content: Config_content;

    constructor(name: string) {
        super(name);

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

    getComponent(): Component {
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
