/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../fetch.pb"
export type AppInfo = {
  name?: string
  version?: string
}

export type ListRequest = {
  limit?: number
  offset?: number
}

export type ApiVersionRequest = {
}

export type ApiVersionResponse = {
  version?: string
}

export type ApiVersion = {
}

export type GetConfigRequest = {
  serviceName?: string
}

export type GetConfigResponse = {
  config?: Uint8Array
}

export type GetConfig = {
}

export type PostConfigRequest = {
  content?: Uint8Array
  serviceName?: string
}

export type PostConfigResponse = {
}

export type PostConfig = {
}

export type PatchConfigRequest = {
  serviceName?: string
  pathToValue?: {[key: string]: string}
}

export type PatchConfigResponse = {
}

export type PatchConfig = {
}

export type ListConfigsRequest = {
  listRequest?: ListRequest
  searchPattern?: string
}

export type ListConfigsResponse = {
  services?: AppInfo[]
}

export type ListConfigs = {
}

export class MatreshkaBeAPI {
  static ApiVersion(req: ApiVersionRequest, initReq?: fm.InitReq): Promise<ApiVersionResponse> {
    return fm.fetchReq<ApiVersionRequest, ApiVersionResponse>(`/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetConfig(req: GetConfigRequest, initReq?: fm.InitReq): Promise<GetConfigResponse> {
    return fm.fetchReq<GetConfigRequest, GetConfigResponse>(`/config/${req["serviceName"]}?${fm.renderURLSearchParams(req, ["serviceName"])}`, {...initReq, method: "GET"})
  }
  static ListConfigs(req: ListConfigsRequest, initReq?: fm.InitReq): Promise<ListConfigsResponse> {
    return fm.fetchReq<ListConfigsRequest, ListConfigsResponse>(`/config/list`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static PostConfig(req: PostConfigRequest, initReq?: fm.InitReq): Promise<PostConfigResponse> {
    return fm.fetchReq<PostConfigRequest, PostConfigResponse>(`/config/${req["serviceName"]}`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static PatchConfig(req: PatchConfigRequest, initReq?: fm.InitReq): Promise<PatchConfigResponse> {
    return fm.fetchReq<PatchConfigRequest, PatchConfigResponse>(`/config/patch/${req["serviceName"]}`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}