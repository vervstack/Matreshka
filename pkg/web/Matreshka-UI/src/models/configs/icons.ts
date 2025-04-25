import unknown from "@/assets/svg/errorconfig.svg";
import matreshka from "@/assets/svg/matreshka.svg";
import pg from "@/assets/svg/pg.svg";
import nginx from "@/assets/svg/nginx.svg";
import minio from "@/assets/svg/minio.svg";

import {ConfigTypePrefix} from "@vervstack/matreshka";

const configIconMap = new Map<ConfigTypePrefix, string>([
    [ConfigTypePrefix.verv, matreshka],
    [ConfigTypePrefix.pg, pg],
    [ConfigTypePrefix.nginx, nginx],
    [ConfigTypePrefix.minio, minio]
])

export function getConfigIcon(configType: ConfigTypePrefix) {
    const icon = configIconMap.get(configType)
    if (icon) {
        return icon
    }

    return unknown
}
