import {Node} from "@vervstack/matreshka";
import {DataSourceClass, Redis} from "@/models/configs/verv/Resources/Resource.ts";
import {extractNumberValue, extractStringValue} from "@/models/shared/common.ts";

export function mapRedis(root: Node) : DataSourceClass {
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
