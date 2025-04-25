import {Node} from "@vervstack/matreshka";

import {extractDataSources} from "@/models/configs/verv/Resources/mapping.ts";
import {mapAppInfo} from "@/models/configs/verv/info/Mapping.ts";
import {mapServer} from "@/models/configs/verv/Servers/Mapping.ts";

import {VervConfig} from "@/models/configs/verv/VervConfig.ts";
import {AppInfoClass} from "@/models/configs/verv/info/VervConfig.ts";
import {ServerClass} from "@/models/configs/verv/Servers/Servers.ts";
import {DataSourceClass} from "@/models/configs/verv/Resources/Resource.ts";

export function parseVervConfigFromEnv(root: Node): VervConfig {
    let appInfo: AppInfoClass | undefined;
    let dataSources: DataSourceClass[] = []
    let servers: ServerClass[] = []

    root.innerNodes?.map((node: Node) => {
        switch (node.name) {
            case 'APP-INFO':
                appInfo = mapAppInfo(node)
                break
            case 'DATA-SOURCES':
                dataSources = extractDataSources(node)
                break
            case 'SERVERS':
                servers = mapServer(node)
        }
    })

    if (!appInfo) {
        throw {message: "No app info found in env"}
    }

    return new VervConfig(appInfo, dataSources, servers)
}
