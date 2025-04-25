export function isConfigNameValid(prefix: string, cfgName: string): boolean {
    if (cfgName.length === 0 || prefix.length === 0) {
        return false
    }

    return /^[a-zA-Z0-9_]+$/.test(cfgName)
}
