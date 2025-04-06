import {Node} from "@godverv/matreshka";

import {extractResourceType} from "@/models/shared/common.ts";

import {mapPostgres} from "@/models/configs/verv/Resources/postgres.ts";
import {mapRedis} from "@/models/configs/verv/Resources/redis.ts";
import {mapSqlite} from "@/models/configs/verv/Resources/sqlite.ts";
import {mapTelegram} from "@/models/configs/verv/Resources/telegram.ts";
import {mapGrpc} from "@/models/configs/verv/Resources/grpc.ts";

import {ResourceType} from "@/models/configs/verv/Resources/ResourceTypes.ts";
import {DataSourceClass} from "@/models/configs/verv/Resources/Resource.ts";

const resourceMapping = new Map<string, (node: Node) => DataSourceClass>()
resourceMapping.set(ResourceType.Postgres, mapPostgres)
resourceMapping.set(ResourceType.Sqlite, mapSqlite)

resourceMapping.set(ResourceType.Redis, mapRedis)
resourceMapping.set(ResourceType.Telegram, mapTelegram)
resourceMapping.set(ResourceType.Grpc, mapGrpc)

export function extractDataSources(root: Node): DataSourceClass[] {

    const dataSources: DataSourceClass[] = []
    root.innerNodes?.map((n) => {
        const resType = extractResourceType(n, root)

        if (!resType) {
            return;
        }

        const mapper = resourceMapping.get(resType)

        if (mapper) {
            dataSources.push(mapper(n))
        }
    })

    return dataSources
}
