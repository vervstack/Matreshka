import {
  Config,
  ConfigTypePrefix,
  CreateConfigRequest,
  GetConfigNodeRequest, GetConfigNodeResponse,
  ListConfigsRequest,
  ListConfigsResponse,
  MatreshkaBeAPI,
  PatchConfigRequest,
  Format,
} from "@vervstack/matreshka";

import ConfigWithContent from "@/models/configs/Config.ts";
import ConfigBase, { defaultVersion } from "@/models/configs/ConfigBase.ts";
import ConfigList from "@/models/configs/ConfigList.ts";
import KeyValueConfig from "@/models/configs/keyvalue/KeyValueConfig.ts";
import VervConfig from "@/models/configs/verv/VervConfig.ts";
import { fromPbEnvNode } from "@/models/shared/Node.ts";

const apiPrefix = {
  pathPrefix: "",
  headers: {
    'Grpc-Metadata-R-Auth': ''
  }
};

export function setBackendAddress(url: string) {
  apiPrefix.pathPrefix = url;
}

export function setPass(pass: string ) {
  apiPrefix.headers['Grpc-Metadata-R-Auth'] = 'Pass '+ pass
}

const fallbackErrorConverting = "error during conversion";

export async function ListServices(req: ListConfigsRequest): Promise<ConfigList> {
  return MatreshkaBeAPI.ListConfigs(req, apiPrefix).then((r: ListConfigsResponse) => {
    if (!r.configs) {
      throw { message: "invalid contract" };
    }

    const servicesInfo: ConfigBase[] = [];

    r.configs.map((v: Config) => {
      const cfgInfo = new ConfigBase(v.name || fallbackErrorConverting);
      cfgInfo.selectedVersion = v.version || cfgInfo.selectedVersion;

      if (v.updatedAtUtcTimestamp) {
        cfgInfo.updated_at = new Date(Number(v.updatedAtUtcTimestamp) * 1000);
      }

      cfgInfo.versions = v.versions || [];

      servicesInfo.push(cfgInfo);
    });

    return new ConfigList(servicesInfo, r.totalRecords || servicesInfo.length);
  });
}

export async function GetConfigNodes(
  configName: string,
  version: string,
): Promise<ConfigWithContent> {
  const req = {
    configName: configName,
    version: version,
  } as GetConfigNodeRequest;

  return MatreshkaBeAPI.GetConfigNodes(req, apiPrefix).then((resp: GetConfigNodeResponse) => {
    if (!resp.root) {
      throw { message: "Empty env config root" };
    }

    const cfg = new ConfigWithContent(configName);

    switch (cfg.type) {
      case ConfigTypePrefix.verv:
        cfg.content = new VervConfig(resp.root);
        break;
      default:
        // TODO
        cfg.content = new KeyValueConfig(fromPbEnvNode(resp.root));
    }

    cfg.versions = resp.versions || [];
    if (req.version) {
      cfg.selectedVersion = req.version;
    }

    cfg.versions.sort();

    const masterIdx = cfg.versions.findIndex(v => v == defaultVersion);
    if (masterIdx != -1) {
      cfg.versions[masterIdx] = cfg.versions[0];
      cfg.versions[0] = defaultVersion;
    }

    return cfg;
  });
}

export async function PatchConfig(cfg: ConfigWithContent): Promise<ConfigWithContent> {
  if (!cfg.isChanged()) return cfg;

  const req: PatchConfigRequest = {
    configName: cfg.getMatreshkaName(),
    version: cfg.selectedVersion,
    patches: cfg.getPatches(),
  } as PatchConfigRequest;

  return MatreshkaBeAPI.PatchConfig(req, apiPrefix).then(() => {
    return GetConfigNodes(cfg.getMatreshkaName(), cfg.selectedVersion);
  });
}

export async function CreateConfig(name: string, confType: ConfigTypePrefix) {
  const newCfg: CreateConfigRequest = {
    configName: encodeURIComponent(name),
    type: confType,
  };

  return MatreshkaBeAPI.CreateConfig(newCfg, apiPrefix);
}

export function linkToConfigSource(configName: string, format: Format, version?: string): string {
  let base = `${apiPrefix.pathPrefix}/download/${configName}`;

  const params: string[] = [];
  if (version) {
    params.push(`version=${version}`);
  }

  if (format != Format.yaml) {
    params.push(`format=${format.toString()}`);
  }

  if (params.length > 0) {
    base += `?${params.join("&")}`;
  }

  return base;
}
