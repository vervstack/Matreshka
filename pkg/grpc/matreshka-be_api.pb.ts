/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../fetch.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);

export enum ResourceType {
  UnknownResourceType = "UnknownResourceType",
  PostgresResourceType = "PostgresResourceType",
  RedisResourceType = "RedisResourceType",
  SqliteResourceType = "SqliteResourceType",
  GrpcResourceType = "GrpcResourceType",
  TelegramResourceType = "TelegramResourceType",
}

export enum ServerType {
  UnknownServerType = "UnknownServerType",
  GrpcServerType = "GrpcServerType",
  RestServerType = "RestServerType",
}

export type ResourceUnknown = {
  environment?: {[key: string]: string}
}

export type ResourcePostgres = {
  host?: string
  port?: number
  dbName?: string
  userName?: string
  pwd?: string
}

export type ResourceRedis = {
  host?: string
  port?: number
  user?: string
  pwd?: string
  db?: number
}

export type ResourceSqlite = {
  path?: string
}

export type ResourceGrpc = {
  connectionString?: string
  module?: string
}

export type ResourceTelegram = {
  apiKey?: string
}

export type ResourceConnection = {
  connectionString?: string
}


type BaseResourceConfig = {
}

export type ResourceConfig = BaseResourceConfig
  & OneOf<{ unknown: ResourceUnknown; postgres: ResourcePostgres; redis: ResourceRedis; sqlite: ResourceSqlite; grpc: ResourceGrpc; telegram: ResourceTelegram }>

export type Resource = {
  conn?: ResourceConnection
  name?: string
  resourceType?: ResourceType
  resourceConfig?: ResourceConfig
}

export type ServerUnknown = {
  environment?: {[key: string]: string}
}

export type ServerGrpc = {
  port?: number
}

export type ServerRest = {
  port?: number
}


type BaseServerConfig = {
}

export type ServerConfig = BaseServerConfig
  & OneOf<{ unknown: ServerUnknown; grpc: ServerGrpc; rest: ServerRest }>

export type Server = {
  swaggerLink?: string
  makoshName?: string
  config?: ServerConfig
  type?: ServerType
}

export type ConfigAppConfig = {
  name?: string
  version?: string
  startupDurationSec?: number
}

export type Config = {
  appConfig?: ConfigAppConfig
  resources?: Resource[]
  servers?: Server[]
  environment?: {[key: string]: string}
}

export type ApiVersionRequest = {
}

export type ApiVersionResponse = {
  version?: string
}

export type ApiVersion = {
}

export type PatchConfigRequest = {
  config?: Config
}

export type PatchConfigResponse = {
}

export type PatchConfig = {
}

export type GetConfigRequest = {
  serviceName?: string
}

export type GetConfigResponse = {
  config?: Config
}

export type GetConfig = {
}

export type GetConfigRawRequest = {
  serviceName?: string
}

export type GetConfigRawResponse = {
  config?: Uint8Array
}

export type GetConfigRaw = {
}

export type PatchConfigRawRequest = {
  raw?: Uint8Array
  serviceName?: string
}

export type PatchConfigRawResponse = {
}

export type PatchConfigRaw = {
}

export type ListRequest = {
  limit?: number
  offset?: number
}

export type ListConfigsRequest = {
  listRequest?: ListRequest
  serviceName?: string
}

export type ListConfigsResponse = {
  services?: ConfigAppConfig[]
}

export type ListConfigs = {
}

export class MatreshkaBeAPI {
  static ApiVersion(req: ApiVersionRequest, initReq?: fm.InitReq): Promise<ApiVersionResponse> {
    return fm.fetchReq<ApiVersionRequest, ApiVersionResponse>(`/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static UpsertConfig(req: PatchConfigRequest, initReq?: fm.InitReq): Promise<PatchConfigResponse> {
    return fm.fetchReq<PatchConfigRequest, PatchConfigResponse>(`/config`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetConfig(req: GetConfigRequest, initReq?: fm.InitReq): Promise<GetConfigResponse> {
    return fm.fetchReq<GetConfigRequest, GetConfigResponse>(`/config/${req["serviceName"]}?${fm.renderURLSearchParams(req, ["serviceName"])}`, {...initReq, method: "GET"})
  }
  static GetConfigRaw(req: GetConfigRawRequest, initReq?: fm.InitReq): Promise<GetConfigRawResponse> {
    return fm.fetchReq<GetConfigRawRequest, GetConfigRawResponse>(`/config/raw/${req["serviceName"]}?${fm.renderURLSearchParams(req, ["serviceName"])}`, {...initReq, method: "GET"})
  }
  static PatchConfigRaw(req: PatchConfigRawRequest, initReq?: fm.InitReq): Promise<PatchConfigRawResponse> {
    return fm.fetchReq<PatchConfigRawRequest, PatchConfigRawResponse>(`/config/raw/${req["serviceName"]}`, {...initReq, method: "PATCH", body: JSON.stringify(req, fm.replacer)})
  }
  static ListConfigs(req: ListConfigsRequest, initReq?: fm.InitReq): Promise<ListConfigsResponse> {
    return fm.fetchReq<ListConfigsRequest, ListConfigsResponse>(`/config/list`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}