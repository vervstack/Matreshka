import {Node} from "@vervstack/matreshka";
import {DataSourceClass, Postgres} from "@/models/configs/verv/Resources/Resource.ts";
import {extractNumberValue, extractStringValue} from "@/models/shared/common.ts";

export function mapPostgres(root: Node): DataSourceClass {
    if (!root.name) {
        throw {message: "No data for postgres to map"}
    }

    const pg = new Postgres(root.name.slice(root.name.indexOf('POSTGRES')).toLowerCase())

    root.innerNodes?.map(
        (n) => {
            if (!n.name || !root.name) {
                return
            }

            const fieldName = n.name.slice(root.name.length + 1)
            switch (fieldName) {
                case "HOST":
                    pg.host = extractStringValue(n)
                    break
                case "PORT":
                    pg.port = extractNumberValue(n)
                    break
                case "USER":
                    pg.user = extractStringValue(n)
                    break
                case "PWD":
                    pg.pwd = extractStringValue(n)
                    break
                case "DB-NAME":
                    pg.name = extractStringValue(n)
                    break
                case "SSL-MODE":
                    pg.ssl_mode = extractStringValue(n)
                    break
            }
        }
    )

    return pg
}
