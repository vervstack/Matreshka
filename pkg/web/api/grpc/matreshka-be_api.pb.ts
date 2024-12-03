/* eslint-disable */
// @ts-nocheck

/**
 * This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
 */

import * as fm from "../../fetch.pb";


export enum SortType {
  default = "default",
  by_name = "by_name",
  by_updated_at = "by_updated_at",
}

export type AppInfo = {
  name?: string;
  version?: string;
  updatedAtUtcTimestamp?: string;
};

export type Paging = {
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

export type PatchConfigRequest = {
  serviceName?: string;
  changes?: Node[];
};

export type PatchConfigResponse = Record<string, never>;

export type PatchConfig = Record<string, never>;

export type ListConfigsRequest = {
  paging?: Paging;
  searchPattern?: string;
  sort?: Sort;
};

export type ListConfigsResponse = {
  services?: AppInfo[];
  totalRecords?: number;
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

export type CreateConfigRequest = {
  serviceName?: string;
};

export type CreateConfigResponse = {
  id?: string;
};

export type CreateConfig = Record<string, never>;

export type RenameConfigRequest = {
  serviceName?: string;
  newName?: string;
};

export type RenameConfigResponse = {
  newName?: string;
};

export type RenameConfig = Record<string, never>;

export type Sort = {
  type?: SortType;
  desc?: boolean;
};

export class MatreshkaBeAPI {
  static ApiVersion(this:void, req: ApiVersionRequest, initReq?: fm.InitReq): Promise<ApiVersionResponse> {
    return fm.fetchRequest<ApiVersionResponse>(`/api/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"});
  }
  static GetConfig(this:void, req: GetConfigRequest, initReq?: fm.InitReq): Promise<GetConfigResponse> {
    return fm.fetchRequest<GetConfigResponse>(`/api/config/${req.serviceName}?${fm.renderURLSearchParams(req, ["serviceName"])}`, {...initReq, method: "GET"});
  }
  static GetConfigNodes(this:void, req: GetConfigNodeRequest, initReq?: fm.InitReq): Promise<GetConfigNodeResponse> {
    return fm.fetchRequest<GetConfigNodeResponse>(`/api/config/nodes/${req.serviceName}?${fm.renderURLSearchParams(req, ["serviceName"])}`, {...initReq, method: "GET"});
  }
  static ListConfigs(this:void, req: ListConfigsRequest, initReq?: fm.InitReq): Promise<ListConfigsResponse> {
    return fm.fetchRequest<ListConfigsResponse>(`/api/config/list`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static CreateConfig(this:void, req: CreateConfigRequest, initReq?: fm.InitReq): Promise<CreateConfigResponse> {
    return fm.fetchRequest<CreateConfigResponse>(`/api/config/${req.serviceName}/new`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static PatchConfig(this:void, req: PatchConfigRequest, initReq?: fm.InitReq): Promise<PatchConfigResponse> {
    return fm.fetchRequest<PatchConfigResponse>(`/api/config/${req.serviceName}/patch`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static RenameConfig(this:void, req: RenameConfigRequest, initReq?: fm.InitReq): Promise<RenameConfigResponse> {
    return fm.fetchRequest<RenameConfigResponse>(`/api/config/${req.serviceName}/rename/${req.newName}`, {...initReq, method: "POST"});
  }
}