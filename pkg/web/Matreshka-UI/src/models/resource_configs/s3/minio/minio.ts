import {ConfigValueClass} from "@/models/shared/common.ts";

export type IamConfig = {
    version: ConfigValueClass<string>;
    statements: Statement[]
}


export type Statement = {
    allow: ConfigValueClass<boolean>
    action: ConfigValueClass<S3Action[]>
    resources: ConfigValueClass<string[]>
}

export enum S3Action {
    ListBucket = "s3:ListBucket",
    GetObject = "s3:GetObject",
    PutObject = "s3:PutObject",
    DeleteObject = "s3:DeleteObject"
}


export type IamMinioExportConfig = {
    Version: string
    Statement: IamMinioStatementExport[]
}

export type IamMinioStatementExport = {
    Effect: string
    Action: string[]
    Resource: string[]
}


export function exportMinioStatement(s: Statement): IamMinioStatementExport {
    return {
        Effect: s.allow.value ? 'Allow' : 'Deny',
        Action: s.action.value.map(v => v.toString()),
        Resource: s.resources.value.map((s) => 'arn:aws:s3:::' + s),
    } as IamMinioStatementExport
}
