import {Node} from "matreshka-api/api/grpc/matreshka-be_api.pb.ts";

import {mapPostgres} from "@/models/AppConfig/Resources/postgres.ts";
import {mapRedis} from "@/models/AppConfig/Resources/redis.ts";
import {mapSqlite} from "@/models/AppConfig/Resources/sqlite.ts";
import {mapTelegram} from "@/models/AppConfig/Resources/telegram.ts";
import {mapGrpc} from "@/models/AppConfig/Resources/grpc.ts";
import {extractResourceType} from "@/models/shared/common.ts";
import {ResourceType} from "@/models/AppConfig/Resources/ResourceTypes.ts";
import {DataSourceClass} from "@/models/AppConfig/Resources/Resource.ts";

const resourceMapping = new Map<string, (node: Node) => DataSourceClass>()
resourceMapping.set(ResourceType.Postgres, mapPostgres)
resourceMapping.set(ResourceType.Sqlite, mapSqlite)

resourceMapping.set(ResourceType.Redis, mapRedis)
resourceMapping.set(ResourceType.Telegram, mapTelegram)
resourceMapping.set(ResourceType.Grpc, mapGrpc)

export function extractDataSources(root: Node): DataSourceClass[] {

    const dataSources : DataSourceClass[] = []
    root.innerNodes?.map((n)=> {
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
