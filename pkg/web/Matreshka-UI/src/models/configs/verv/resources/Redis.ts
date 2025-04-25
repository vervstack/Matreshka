import {Node} from "@vervstack/matreshka";
import DataSource from "@/models/configs/verv/resources/Resource.ts";
import {ConfigValue, extractNumberValue, extractStringValue} from "@/models/shared/Common.ts";
import {ResourceType} from "@/models/configs/verv/resources/ResourceTypes.ts";
import {Change} from "@/models/configs/Change.ts";

export default class Redis extends DataSource {
    host: ConfigValue<string> = new ConfigValue<string>("", "")
    port: ConfigValue<number> = new ConfigValue<number>("", 0)
    user: ConfigValue<string> = new ConfigValue<string>("", "")
    pwd: ConfigValue<string> = new ConfigValue<string>("", "")
    db: ConfigValue<number> = new ConfigValue<number>("", 0)

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

export function mapRedis(root: Node) : DataSource {
    if (!root.name) {
        throw  {message: "Can't parse redis config"}
    }

    const rds = new Redis(root.name.slice(root.name.indexOf('REDIS')).toLowerCase())

    root.innerNodes?.map(
        (n) => {
            if (!n.name || !root.name) {
                return
            }

            const fieldName = n.name.slice(root.name.length + 1)
            switch (fieldName) {
                case "HOST":
                    rds.host = extractStringValue(n)
                    break
                case "PORT":
                    rds.port = extractNumberValue(n)
                    break
                case "USER":
                    rds.user = extractStringValue(n)
                    break
                case "PWD":
                    rds.pwd = extractStringValue(n)
                    break
                case "DB":
                    rds.db = extractNumberValue(n)
            }
        }
    )

    return rds;
}
