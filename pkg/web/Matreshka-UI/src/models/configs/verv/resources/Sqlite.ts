import {Node} from "@vervstack/matreshka";
import DataSource from "@/models/configs/verv/resources/Resource.ts";
import {ConfigValue, extractStringValue} from "@/models/shared/Common.ts";
import {ResourceType} from "@/models/configs/verv/resources/ResourceTypes.ts";
import {Change} from "@/models/configs/Change.ts";

export default class Sqlite extends DataSource {
    path: ConfigValue<string> = new ConfigValue("", "")

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

export function mapSqlite(root: Node): DataSource {
    if (!root.name) {
        throw {message: "no data to parse sqlite"}
    }


    const sqlite = new Sqlite(root.name.slice(root.name.indexOf('SQLITE')).toLowerCase())

    root.innerNodes?.map(
        (n) => {
            if (!n.name || !root.name) {
                return
            }

            const fieldName = n.name.slice(root.name.length + 1)
            switch (fieldName) {
                case 'PATH':
                    sqlite.path = extractStringValue(n)
            }
        }
    )


    return sqlite
}
