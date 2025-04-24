import {ConfigValueClass} from "@/models/shared/common.ts";

export class ServiceListClass {
    servicesInfo: AppInfoClass[]
    total: number

    constructor(servicesInfo: AppInfoClass[], total: number) {
        this.servicesInfo = servicesInfo;
        this.total = total;
    }
}

export class AppInfoClass {
    name: ConfigValueClass<string>
    serviceVersion: ConfigValueClass<string>
    public updated_at?: Date

    versions: string[] = []

    constructor(name: ConfigValueClass<string>, serviceVersion: ConfigValueClass<string>) {
        this.name = name;
        this.serviceVersion = serviceVersion;
    }

    getChanges(): Change[] {
        const changes: Change[] = []
        changes.push(...this.name.getChanges())
        changes.push(...this.serviceVersion.getChanges())
        return changes
    }

    isChanged() : boolean{
        return this.getChanges().length != 0
    }

    rollback() {
        this.name.rollback()
        this.serviceVersion.rollback()
    }
}

//  TODO move to models/configs
export type Change = {
    envName: string
    newValue: string
}

export enum SourceCodeSystem {
    unknown = 0,
    github = 1,
}

export function ExtractSourceCodeSystemFromServiceName(name: string): SourceCodeSystem | undefined {
    if (name.includes("github")) {
        return SourceCodeSystem.github
    }
}

const scsToPiIcon = new Map<SourceCodeSystem, string>()
scsToPiIcon.set(SourceCodeSystem.github, 'pi-github')

export function PiIconFromSourceCodeSystem(sourceCodeSystem: SourceCodeSystem | undefined): string | undefined {
    if (!sourceCodeSystem) {
        return
    }

    return scsToPiIcon.get(sourceCodeSystem)
}

