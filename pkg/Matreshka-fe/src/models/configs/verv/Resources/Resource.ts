import {Component} from "vue";

import {ConfigValueClass} from "@/models/shared/common.ts";
import {getResourceLink} from "@/models/shared/s3.ts";

import {Change} from "@/models/configs/verv/info/AppInfo.ts";
import {ResourceType} from "@/models/configs/verv/Resources/ResourceTypes.ts";

type NamedResource = {
    resource_name: string
    type: ResourceType
}

export type ResourceSqlite = NamedResource & {
    path: ConfigValueClass<string>
}

export type ResourceTelegram = NamedResource & {
    api_key: ConfigValueClass<string>
}

export type ResourceGrpc = NamedResource & {
    connection_string: ConfigValueClass<string>
    module: ConfigValueClass<string>
}

export function NormalizeName(res: NamedResource): string {
    if (res.resource_name === res.type) {
        return res.resource_name
    }


    let out = res.resource_name.slice(res.type.length)
    if (out[0] === '-') {
        out = out.slice(1)
    }

    return out
}


export abstract class DataSourceClass {
    resourceName: string
    readonly type: ResourceType

    private readonly resourceTypeToImagePath = new Map<string, string>()
        .set(ResourceType.Postgres, getResourceLink('pg.png'))
        .set(ResourceType.Redis, getResourceLink('redis.png'))
        .set(ResourceType.Sqlite, getResourceLink('sqlite.png'))
        .set(ResourceType.Grpc, getResourceLink('grpc.png'))
        .set(ResourceType.Telegram, getResourceLink('telegram.png'))

    constructor(resourceName: string, resType: ResourceType) {
        this.resourceName = resourceName;
        this.type = resType;
    }

    normalizeName() {
        if (this.resourceName === this.type) {
            return this.resourceName
        }


        let out = this.resourceName.slice(this.type.length)
        if (out[0] === '-') {
            out = out.slice(1)
        }

        return out
    }

    getComponent(): Component {
        return ResourceType.GetComponent(this.type)
    }

    getIcon(): string {
        const imagePath = this.resourceTypeToImagePath.get(this.type)
        if (imagePath) {
            return imagePath
        }

        return getResourceLink('unknown.png');
    }

    isChanged(): boolean {
        return this.getChanges().length != 0
    }

    abstract rollback(): void

    abstract getChanges(): Change[]
}

export class Postgres extends DataSourceClass {
    host: ConfigValueClass<string> = new ConfigValueClass("", "")
    name: ConfigValueClass<string> = new ConfigValueClass("", "")
    port: ConfigValueClass<number> = new ConfigValueClass("", 0)
    user: ConfigValueClass<string> = new ConfigValueClass("", "")
    pwd: ConfigValueClass<string> = new ConfigValueClass("", "")
    ssl_mode: ConfigValueClass<string> = new ConfigValueClass("", "")

    constructor(resourceName: string) {
        super(resourceName, ResourceType.Postgres);
    }

    rollback(): void {
        this.host.rollback()
        this.name.rollback()
        this.port.rollback()
        this.user.rollback()
        this.pwd.rollback()
        this.ssl_mode.rollback()
    }

    getChanges(): Change[] {
        const changes: Change[] = []

        changes.push(...this.host.getChanges())
        changes.push(...this.name.getChanges())
        changes.push(...this.port.getChanges())
        changes.push(...this.user.getChanges())
        changes.push(...this.pwd.getChanges())
        changes.push(...this.ssl_mode.getChanges())

        return changes
    }
}

export class Sqlite extends DataSourceClass {
    path: ConfigValueClass<string> = new ConfigValueClass("", "")

    constructor(resourceName: string) {
        super(resourceName, ResourceType.Sqlite);
    }

    rollback(): void {
        this.path.rollback()
    }

    getChanges(): Change[] {
        const changes: Change[] = []

        changes.push(...this.path.getChanges())

        return changes
    }
}

export class Redis extends DataSourceClass {
    host: ConfigValueClass<string> = new ConfigValueClass<string>("", "")
    port: ConfigValueClass<number> = new ConfigValueClass<number>("", 0)
    user: ConfigValueClass<string> = new ConfigValueClass<string>("", "")
    pwd: ConfigValueClass<string> = new ConfigValueClass<string>("", "")
    db: ConfigValueClass<number> = new ConfigValueClass<number>("", 0)

    constructor(resourceName: string) {
        super(resourceName, ResourceType.Redis);
    }

    rollback(): void {
        this.host.rollback()
        this.port.rollback()
        this.user.rollback()
        this.pwd.rollback()
        this.db.rollback()
    }

    getChanges(): Change[] {
        const changes: Change[] = []

        changes.push(...this.host.getChanges())
        changes.push(...this.port.getChanges())
        changes.push(...this.user.getChanges())
        changes.push(...this.pwd.getChanges())
        changes.push(...this.db.getChanges())

        return changes
    }
}

export class Telegram extends DataSourceClass {
    apiKey: ConfigValueClass<string> = new ConfigValueClass("", "")

    constructor(resourceName: string) {
        super(resourceName, ResourceType.Telegram);
    }

    rollback(): void {
        this.apiKey.rollback()
    }

    getChanges(): Change[] {
        const changes: Change[] = []

        changes.push(...this.apiKey.getChanges())

        return changes
    }

}

export class GrpcClient extends DataSourceClass {
    connectionString: ConfigValueClass<string> = new ConfigValueClass("", "")
    module: ConfigValueClass<string> = new ConfigValueClass("", "")

    constructor(resourceName: string) {
        super(resourceName, ResourceType.Grpc);
    }

    rollback(): void {
        this.connectionString.rollback()
        this.module.rollback()
    }

    getChanges(): Change[] {
        const changes: Change[] = []

        changes.push(...this.connectionString.getChanges())
        changes.push(...this.module.getChanges())

        return changes
    }
}


