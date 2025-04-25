import {Component} from "vue";

import {Node} from "@vervstack/matreshka";

import AppInfoClass from "@/models/configs/verv/info/VervConfig.ts";
import DataSourceClass from "@/models/configs/verv/resources/Resource.ts";
import ServerClass from "@/models/configs/verv/servers/Servers.ts";

import ConfigContent from "@/models/configs/ConfigContent.ts";
import VervConfigView from "@/components/config/verv/VervConfigView.vue";

import {extractDataSources} from "@/models/configs/verv/resources/mapping.ts";
import {mapServer} from "@/models/configs/verv/servers/Mapping.ts";
import {Change} from "@/models/configs/Change.ts";

export default class VervConfig implements ConfigContent {
    appInfo: AppInfoClass
    dataSources: DataSourceClass[]
    servers: ServerClass[]

    constructor(root: Node) {
        let appInfo: AppInfoClass | undefined;
        let dataSources: DataSourceClass[] = []
        let servers: ServerClass[] = []

        root.innerNodes?.map((node: Node) => {
            switch (node.name) {
                case 'APP-INFO':
                    appInfo = new AppInfoClass(root)
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

    getComponent(): Component {
        return VervConfigView;
    }
}
