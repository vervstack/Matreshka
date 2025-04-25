const s3Endpoint = 'https://s3-api.redsock.ru/verv/matreshka/'

export function getResourceLink(resource: string): string {
    return `${s3Endpoint}/${resource}`
}
