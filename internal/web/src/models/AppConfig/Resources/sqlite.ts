import {Node} from "matreshka-api/api/grpc/matreshka-be_api.pb.ts";
import {DataSourceClass, Sqlite} from "@/models/AppConfig/Resources/Resource.ts";
import {extractStringValue} from "@/models/shared/common.ts";


export function mapSqlite(root: Node): DataSourceClass {
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
