import {Component} from "vue";

import {getResourceLink} from "@/models/shared/S3.ts";

import {Change} from "@/models/configs/Change.ts";
import {ResourceType} from "@/models/configs/verv/resources/ResourceTypes.ts";

export default abstract class DataSource {
    resourceName: string
    readonly type: ResourceType

    private readonly resourceTypeToImagePath = new Map<string, string>()
        .set(ResourceType.Postgres, getResourceLink('pg.png'))
        .set(ResourceType.Redis, getResourceLink('redis.png'))
        .set(ResourceType.Sqlite, getResourceLink('sqlite.png'))
        .set(ResourceType.Grpc, getResourceLink('grpc.png'))
        .set(ResourceType.Telegram, getResourceLink('telegram.png'))

    protected constructor(resourceName: string, resType: ResourceType) {
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
