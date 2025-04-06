import unknown from "@/assets/svg/errorconfig.svg";
import matreshka from "@/assets/svg/matreshka.svg";

import {ConfigTypePrefix} from "@godverv/matreshka";

const configIconMap = new Map<ConfigTypePrefix, string>([
    [ConfigTypePrefix.verv, matreshka]
])

export function getConfigIcon(configType: ConfigTypePrefix) {
    const icon = configIconMap.get(configType)
    if (icon) {
        console.log(icon)
        return icon
    }

    return unknown
}
