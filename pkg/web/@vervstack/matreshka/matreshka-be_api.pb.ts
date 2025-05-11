/* eslint-disable */
// @ts-nocheck

/**
 * This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
 */

import * as fm from "./fetch.pb";

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };

type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (keyof T extends infer K
      ? K extends string & keyof T
        ? { [k in K]: T[K] } & Absent<T, K>
        : never
      : never);

export enum ConfigTypePrefix {
  unknown = "unknown",
  verv = "verv",
  minio = "minio",
  pg = "pg",
  nginx = "nginx",
}

export enum Format {
  yaml = "yaml",
  env = "env",
}

export enum SortType {
  default = "default",
  by_name = "by_name",
  by_updated_at = "by_updated_at",
}

export type Config = {
  name?: string;
  version?: string;
  updatedAtUtcTimestamp?: string;
  versions?: string[];
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
  configName?: string;
  version?: string;
  format?: Format;
};

export type GetConfigResponse = {
  config?: Uint8Array;
};

export type GetConfig = Record<string, never>;

export type PatchConfigRequest = {
  configName?: string;
  version?: string;
  patches?: PatchConfigPatch[];
};

export type PatchConfigResponse = Record<string, never>;

type BasePatchConfigPatch = {
  fieldName?: string;
};

export type PatchConfigPatch = BasePatchConfigPatch &
  OneOf<{
    rename: string;
    updateValue: string;
    delete: boolean;
  }>;

export type PatchConfig = Record<string, never>;

export type ListConfigsRequest = {
  paging?: Paging;
  searchPattern?: string;
  sort?: Sort;
};

export type ListConfigsResponse = {
  configs?: Config[];
  totalRecords?: number;
};

export type ListConfigs = Record<string, never>;

export type Node = {
  name?: string;
  value?: string;
  innerNodes?: Node[];
};

export type GetConfigNodeRequest = {
  configName?: string;
  version?: string;
};

export type GetConfigNodeResponse = {
  root?: Node;
  versions?: string[];
};

export type GetConfigNode = Record<string, never>;

export type CreateConfigRequest = {
  configName?: string;
};

export type CreateConfigResponse = {
  id?: string;
};

export type CreateConfig = Record<string, never>;

export type RenameConfigRequest = {
  configName?: string;
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

export type SubscribeOnChangesRequest = {
  subscribeConfigNames?: string[];
  unsubscribeConfigNames?: string[];
};

export type SubscribeOnChangesResponse = {
  configName?: string;
  timestamp?: number;
  patches?: PatchConfigPatch[];
};

export type SubscribeOnChanges = Record<string, never>;

export class MatreshkaBeAPI {
  static ApiVersion(this:void, req: ApiVersionRequest, initReq?: fm.InitReq): Promise<ApiVersionResponse> {
    return fm.fetchRequest<ApiVersionResponse>(`/api/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"});
  }
  static GetConfig(this:void, req: GetConfigRequest, initReq?: fm.InitReq): Promise<GetConfigResponse> {
    return fm.fetchRequest<GetConfigResponse>(`/api/config/${req.configName}?${fm.renderURLSearchParams(req, ["configName"])}`, {...initReq, method: "GET"});
  }
  static GetConfigNodes(this:void, req: GetConfigNodeRequest, initReq?: fm.InitReq): Promise<GetConfigNodeResponse> {
    return fm.fetchRequest<GetConfigNodeResponse>(`/api/config/nodes`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static ListConfigs(this:void, req: ListConfigsRequest, initReq?: fm.InitReq): Promise<ListConfigsResponse> {
    return fm.fetchRequest<ListConfigsResponse>(`/api/config/list`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static CreateConfig(this:void, req: CreateConfigRequest, initReq?: fm.InitReq): Promise<CreateConfigResponse> {
    return fm.fetchRequest<CreateConfigResponse>(`/api/config/${req.configName}/new`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static PatchConfig(this:void, req: PatchConfigRequest, initReq?: fm.InitReq): Promise<PatchConfigResponse> {
    return fm.fetchRequest<PatchConfigResponse>(`/api/config/${req.configName}/patch`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static RenameConfig(this:void, req: RenameConfigRequest, initReq?: fm.InitReq): Promise<RenameConfigResponse> {
    return fm.fetchRequest<RenameConfigResponse>(`/api/config/${req.configName}/rename/${req.newName}`, {...initReq, method: "POST"});
  }
}