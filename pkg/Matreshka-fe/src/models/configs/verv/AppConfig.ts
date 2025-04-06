import {AppInfoClass, Change} from "@/models/configs/verv/info/AppInfo.ts";
import {DataSourceClass} from "@/models/configs/verv/Resources/Resource.ts";
import {ServerClass} from "@/models/configs/verv/Servers/Servers.ts";

export class AppConfigClass {
    appInfo: AppInfoClass
    dataSources: DataSourceClass[]
    servers: ServerClass[]

    constructor(appInfo: AppInfoClass, dataSources: DataSourceClass[], servers: ServerClass[]) {
        this.appInfo = appInfo;
        this.dataSources = dataSources;
        this.servers = servers;
    }

    public isChanged(): boolean {
        return this.getChanges().length != 0
    }

    public getChanges(): Change[] {
        const changes: Change[] = []
        changes.push(...this.appInfo.getChanges())

        this.dataSources.map(ds => changes.push(...ds.getChanges()))

        this.servers.map(s => changes.push(...s.getChanges()))

        return changes
    }

    public getChangedDataSourcesNames(): string[] {
        const changedDataSourceNames: string[] = []
        this.dataSources.map(ds => {
            if (ds.isChanged()) {
                changedDataSourceNames.push(ds.resourceName)
            }
        })

        return changedDataSourceNames
    }

    public getChangedServersNames(): string[] {
        const changedServerNames: string[] = []
        this.servers.map(serv => {
            if (serv.isChanged()) {
                changedServerNames.push(serv.name)
            }
        })
        return changedServerNames
    }

    public rollback() {
        this.appInfo.rollback()
        this.dataSources.map(ds => ds.rollback())
        this.servers.map(s => s.rollback())
    }
}
