import {
    MatreshkaBeAPI,
    ListConfigsRequest,
    GetConfigNodeRequest,
    Node, CreateConfigRequest, PatchConfigRequest, Config, ListConfigsResponse, ConfigTypePrefix
} from "@vervstack/matreshka";

import {parseVervConfigFromEnv} from "@/processes/api/model_mapping.ts";

import {AppInfoClass, Change, ServiceListClass} from "@/models/configs/verv/info/AppInfo.ts";
import {ConfigValueClass} from "@/models/shared/common.ts";
import {Config as ConfigWithContent} from "@/models/configs/config.ts";
import {KeyValueConfigContent} from "@/models/configs/keyvalue/config.ts";


const apiPrefix = {pathPrefix: ''};

export function setBackendAddress(url: string) {
    apiPrefix.pathPrefix = url
}

const fallbackErrorConverting = 'error during conversion'

export async function ListServices(req: ListConfigsRequest): Promise<ServiceListClass> {
    return MatreshkaBeAPI
        .ListConfigs(req, apiPrefix)
        .then((r: ListConfigsResponse) => {
                const servicesInfo: AppInfoClass[] = []
                if (!r.configs) {
                    throw {message: "invalid contract"}
                }

                r.configs
                    .map((v: Config) => {
                        const name = new ConfigValueClass(
                            "Service name",
                            v.name || fallbackErrorConverting,
                        )

                        const version = new ConfigValueClass(
                            "Version",
                            v.version || fallbackErrorConverting,
                        )

                        const appInfo = new AppInfoClass(name, version)
                        if (v.updatedAtUtcTimestamp) {
                            appInfo.updated_at = new Date(Number(v.updatedAtUtcTimestamp) * 1000)
                        }

                        appInfo.versions = v.versions || []

                        servicesInfo.push(appInfo)
                    })

                return new ServiceListClass(servicesInfo, r.totalRecords || servicesInfo.length)
            }
        )
}

export async function GetConfigNodes(configName: string, version: string): Promise<ConfigWithContent> {
    const req = {
        configName: configName,
        version: version,
    } as GetConfigNodeRequest;

    return MatreshkaBeAPI.GetConfigNodes(req, apiPrefix)
        .then((res) => {
            if (!res.root) {
                throw {message: "Empty env config root"}
            }

            const cfg = new ConfigWithContent(configName);

            switch (cfg.type) {
                case ConfigTypePrefix.verv:
                    cfg.content = parseVervConfigFromEnv(res.root)
                    break;
                default:
                    // TODO
                    cfg.content = new KeyValueConfigContent()
            }

            return cfg
        })
}

export async function PatchConfig(serviceName: string, version: string, changeList: Change[]): Promise<ConfigWithContent> {
    const req: PatchConfigRequest = {
        configName: serviceName,
        version: version,
        changes: changeList.map((n) => {
            return {
                name: n.envName,
                value: n.newValue,
            } as Node
        }),
    } as PatchConfigRequest;


    return MatreshkaBeAPI.PatchConfig(req, apiPrefix)
        .then(() => {
            changeList.map((n) => {
                if (n.envName.includes('APP-INFO_NAME')) {
                    serviceName = n.newValue
                }
            })
            return GetConfigNodes(serviceName, version)
        })
}

export async function CreateConfig(name: string) {
    const newCfg: CreateConfigRequest = {
        configName: encodeURIComponent(name)
    }

    return MatreshkaBeAPI.CreateConfig(newCfg, apiPrefix)
}
