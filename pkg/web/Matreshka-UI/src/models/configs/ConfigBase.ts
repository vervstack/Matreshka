import { ConfigTypePrefix } from "@vervstack/matreshka";

export const defaultVersion = "master";

export default class ConfigBase {
  type: ConfigTypePrefix;
  name: string;

  updated_at?: Date;

  versions: string[] = [defaultVersion];
  selectedVersion: string = defaultVersion;

  constructor(name: string) {
    this.type = extractType(name);

    if (this.type != ConfigTypePrefix.plain) {
      name = name.substring(this.type.length + 1);
    }
    this.name = name;
  }

  getMatreshkaName(): string {
    let name = this.name;
    if (this.type != ConfigTypePrefix.plain) {
      name = this.type + "_" + this.name;
    }

    return name;
  }
}

const supportedTypes: ConfigTypePrefix[] = [
  ConfigTypePrefix.verv,
  ConfigTypePrefix.minio,
  ConfigTypePrefix.pg,
  ConfigTypePrefix.nginx,
];

function extractType(configName: string): ConfigTypePrefix {
  const foundType = supportedTypes.find((typePrefix: ConfigTypePrefix) =>
    configName.startsWith(typePrefix),
  );

  if (foundType) {
    return foundType;
  }

  return ConfigTypePrefix.plain;
}
