import {Node} from "@godverv/matreshka";
import {AppConfigClass} from "@/models/AppConfig/AppConfig.ts";
import {extractDataSources} from "@/models/AppConfig/Resources/mapping.ts";

import {mapAppInfo} from "@/models/AppConfig/Info/Mapping.ts";
import {mapServer} from "@/models/AppConfig/Servers/Mapping.ts";
import {AppInfoClass} from "@/models/AppConfig/Info/AppInfo.ts";
import {DataSourceClass} from "@/models/AppConfig/Resources/Resource.ts";
import {ServerClass} from "@/models/AppConfig/Servers/Servers.ts";

export function parseAppConfigFromEnv(root: Node): AppConfigClass {
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

    return new AppConfigClass(appInfo, dataSources, servers)
}
