import {Node} from "@vervstack/matreshka";

import {ConfigValueClass, extractStringValue} from "@/models/shared/common.ts";
import { AppInfoClass } from "@/models/configs/verv/info/AppInfo.ts";

export function mapAppInfo(root: Node): AppInfoClass {

    let appName : ConfigValueClass<string>| undefined;
    let appVersion : ConfigValueClass<string> | undefined;

    root.innerNodes?.map((n)=> {
        if (!n.name || !root.name) {
            return;
        }

        const name = n.name.slice(root.name.length+1)
        switch (name) {
            case "NAME":
                const name = extractStringValue(n);
                appName = new ConfigValueClass<string>(name.envName, name.value)
                break
            case "VERSION":
                const version = extractStringValue(n);
                appVersion = new ConfigValueClass<string>(version.envName, version.value)
                break
        }

        return
    })

    if (!appName) {
        throw {message: "no app name provided"}
    }

    if (!appVersion) {
        throw {message: "no app version provided"}
    }

    return new AppInfoClass(appName, appVersion)
}
