import {ConfigValueClass} from "@/models/shared/common.ts";
import {Change} from "@/models/configs/verv/info/AppInfo.ts";


export class ServerClass {
    port: ConfigValueClass<number> = new ConfigValueClass<number>("", 0)
    name: string
    grpc: GrpcHandler[] = []
    fs: FsHandler[] = []

    constructor(name: string) {
        this.name = name
    }

    public isChanged(): boolean {
        let grpcChanged: boolean = false
        this.grpc.map(s => grpcChanged = grpcChanged || s.isChanged())

        let fsChanged: boolean = false
        this.fs.map(s => fsChanged = fsChanged || s.isChanged())
        return this.port.isChanged() || grpcChanged || fsChanged
    }

    rollback(): void {
        this.port.rollback()
        this.grpc.map(g => g.rollback())
        this.fs.map((f => f.rollback()))
    }

    getChanges() : Change[]{
        const changes: Change[] = []
        changes.push(...this.port.getChanges())

        this.grpc.map(g => g.getChanges())
        this.fs.map((f => f.getChanges()))

        return changes
    }
}

export class GrpcHandler {
    module: ConfigValueClass<string> = new ConfigValueClass("", "")
    gateway: ConfigValueClass<string> = new ConfigValueClass("", "")

    isChanged(): boolean {
        return this.module.isChanged() || this.gateway.isChanged()
    }

    rollback(): void {
        this.module.rollback()
        this.gateway.rollback()
    }

    getChanges(): Change[] {
        const changes : Change[] = []

        changes.push(...this.module.getChanges())
        changes.push(...this.gateway.getChanges())

        return changes
    }
}

export class FsHandler {
    dist: ConfigValueClass<string> = new ConfigValueClass<string>("", "")

    isChanged(): boolean {
        return this.dist.isChanged()
    }

    rollback(): void {
        this.dist.rollback()
    }

    getChanges() : Change[]{
        const changes: Change[] = []

        changes.push(...this.dist.getChanges())

        return changes
    }
}
