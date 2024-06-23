/* eslint-disable */
// @ts-nocheck

/**
 * This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
 */

import * as fm from "../fetch.pb";


export type AppInfo = {
  name?: string;
  version?: string;
};

export type ListRequest = {
  limit?: number;
  offset?: number;
};

export type ApiVersionRequest = Record<string, never>;

export type ApiVersionResponse = {
  version?: string;
};

export type ApiVersion = Record<string, never>;

export type GetConfigRequest = {
  serviceName?: string;
};

export type GetConfigResponse = {
  config?: Uint8Array;
};

export type GetConfig = Record<string, never>;

export type PostConfigRequest = {
  content?: Uint8Array;
  serviceName?: string;
};

export type PostConfigResponse = Record<string, never>;

export type PostConfig = Record<string, never>;

export type PatchConfigRequest = {
  serviceName?: string;
  changes?: Node[];
};

export type PatchConfigResponse = Record<string, never>;

export type PatchConfig = Record<string, never>;

export type ListConfigsRequest = {
  listRequest?: ListRequest;
  searchPattern?: string;
};

export type ListConfigsResponse = {
  services?: AppInfo[];
};

export type ListConfigs = Record<string, never>;

export type Node = {
  name?: string;
  value?: string;
  innerNodes?: Node[];
};

export type GetConfigNodeRequest = {
  serviceName?: string;
};

export type GetConfigNodeResponse = {
  root?: Node;
};

export type GetConfigNode = Record<string, never>;

export class MatreshkaBeAPI {
  static ApiVersion(this:void, req: ApiVersionRequest, initReq?: fm.InitReq): Promise<ApiVersionResponse> {
    return fm.fetchRequest<ApiVersionResponse>(`/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"});
  }
  static GetConfig(this:void, req: GetConfigRequest, initReq?: fm.InitReq): Promise<GetConfigResponse> {
    return fm.fetchRequest<GetConfigResponse>(`/config/${req.serviceName}?${fm.renderURLSearchParams(req, ["serviceName"])}`, {...initReq, method: "GET"});
  }
  static GetConfigNodes(this:void, req: GetConfigNodeRequest, initReq?: fm.InitReq): Promise<GetConfigNodeResponse> {
    return fm.fetchRequest<GetConfigNodeResponse>(`/config/nodes/${req.serviceName}?${fm.renderURLSearchParams(req, ["serviceName"])}`, {...initReq, method: "GET"});
  }
  static ListConfigs(this:void, req: ListConfigsRequest, initReq?: fm.InitReq): Promise<ListConfigsResponse> {
    return fm.fetchRequest<ListConfigsResponse>(`/config/list`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static PostConfig(this:void, req: PostConfigRequest, initReq?: fm.InitReq): Promise<PostConfigResponse> {
    return fm.fetchRequest<PostConfigResponse>(`/config/create/${req.serviceName}`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static PatchConfig(this:void, req: PatchConfigRequest, initReq?: fm.InitReq): Promise<PatchConfigResponse> {
    return fm.fetchRequest<PatchConfigResponse>(`/config/patch/${req.serviceName}`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
}