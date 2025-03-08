import {
    MatreshkaBeAPI,
    ListConfigsRequest,
    GetConfigNodeRequest,
    Node, CreateConfigRequest, PatchConfigRequest, AppInfo
} from "@godverv/matreshka";

import {parseAppConfigFromEnv} from "@/processes/Api/ModelMapping.ts";
import {AppInfoClass, Change, ServiceListClass} from "@/models/AppConfig/Info/AppInfo.ts";
import {AppConfigClass} from "@/models/AppConfig/AppConfig.ts";
import {getBackendUrl} from "@/app/store/settings.ts";
import {ConfigValueClass} from "@/models/shared/common.ts";

const prefix = {pathPrefix: getBackendUrl()};

export function setBackendUrl(url: string) {
    prefix.pathPrefix = url
}

const fallbackErrorConverting = 'error during convertion'

export async function ListServices(req: ListConfigsRequest): Promise<ServiceListClass> {
    return MatreshkaBeAPI
        .ListConfigs(req, prefix)
        .then((r) => {
                const servicesInfo: AppInfoClass[] = []
                if (!r.services) {
                    throw {message: "invalid contract"}
                }

                r.services
                    .map((v: AppInfo) => {
                        const name = new ConfigValueClass(
                            "Service name",
                            v.name || fallbackErrorConverting,
                        )

                        const version = new ConfigValueClass(
                            "Version",
                            v.serviceVersion || fallbackErrorConverting,
                        )

                        const appInfo = new AppInfoClass(name, version)
                        if (v.updatedAtUtcTimestamp) {
                            appInfo.updated_at = new Date(Number(v.updatedAtUtcTimestamp) * 1000)
                        }

                        appInfo.versions = v.configVersions || []

                        servicesInfo.push(appInfo)
                    })

                return new ServiceListClass(servicesInfo, r.totalRecords || servicesInfo.length)
            }
        )
}

export async function GetConfigNodes(serviceName: string, version: string): Promise<AppConfigClass> {
    const req = {
        serviceName: serviceName,
        version: version,
    } as GetConfigNodeRequest;

    return MatreshkaBeAPI.GetConfigNodes(req, prefix)
        .then((res) => {
            if (!res.root) {
                throw {message: "Empty env config root"}
            }

            return parseAppConfigFromEnv(res.root);
        })
}

export async function PatchConfig(serviceName: string, version: string, changeList: Change[]) {
    const req: PatchConfigRequest = {
        serviceName: serviceName,
        version: version,
        changes: changeList.map((n) => {
            return {
                name: n.envName,
                value: n.newValue,
            } as Node
        }),
    } as PatchConfigRequest;

    return MatreshkaBeAPI.PatchConfig(req, prefix)
        .then(() => GetConfigNodes(serviceName, version))
}

export async function CreateConfig(name: string) {
    const newCfg = {
        serviceName: encodeURIComponent(name)
    } as CreateConfigRequest

    return MatreshkaBeAPI.CreateConfig(newCfg, prefix)
}
