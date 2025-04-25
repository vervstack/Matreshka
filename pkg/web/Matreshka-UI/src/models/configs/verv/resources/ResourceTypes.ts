import {Component} from "vue";

import ResourcePostgres from "@/components/config/verv/resource/ResourcePostgres.vue";
import ResourceGrpc from "@/components/config/verv/resource/ResourceGrpc.vue";
import ResourceRedis from "@/components/config/verv/resource/ResourceRedis.vue";
import ResourceSqlite from "@/components/config/verv/resource/ResourceSqlite.vue";
import ResourceTelegram from "@/components/config/verv/resource/ResourceTelegram.vue";
import KeyMapComponent from "@/components/base/config/fields/KeyMap.vue";

export enum ResourceType {
    Postgres = "postgres",
    Redis = "redis",
    Sqlite = "sqlite",
    Grpc = "grpc",
    Telegram = "telegram"
}

export namespace ResourceType {
    const typeToDefinition = new Map<ResourceType, Component>()
    typeToDefinition.set(ResourceType.Postgres, ResourcePostgres)
    typeToDefinition.set(ResourceType.Grpc, ResourceGrpc)
    typeToDefinition.set(ResourceType.Redis, ResourceRedis)
    typeToDefinition.set(ResourceType.Sqlite, ResourceSqlite)
    typeToDefinition.set(ResourceType.Telegram, ResourceTelegram)

    export function GetComponent(rt: ResourceType): Component {
        return typeToDefinition.get(rt) || KeyMapComponent
    }
}
