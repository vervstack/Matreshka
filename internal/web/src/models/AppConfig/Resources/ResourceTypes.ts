import ResourcePostgres from "@/components/config/Resource/types/ResourcePostgres.vue";
import ResourceGrpc from "@/components/config/Resource/types/ResourceGrpc.vue";
import ResourceRedis from "@/components/config/Resource/types/ResourceRedis.vue";
import ResourceSqlite from "@/components/config/Resource/types/ResourceSqlite.vue";
import ResourceTelegram from "@/components/config/Resource/types/ResourceTelegram.vue";
import KeyMapComponent from "@/components/base/KeyMapComponent.vue";
import {Component} from "vue";

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
