import {
  MatreshkaBeAPI,
  ListConfigsRequest,
  GetConfigNodeRequest,
  Node,
  CreateConfigRequest,
  PatchConfigRequest,
  Config,
  ListConfigsResponse,
  ConfigTypePrefix,
} from "@vervstack/matreshka";

import ConfigWithContent from "@/models/configs/Config.ts";
import ConfigBase from "@/models/configs/ConfigBase.ts";
import ConfigList from "@/models/configs/ConfigList.ts";
import KeyValueConfigContent from "@/models/configs/keyvalue/KeyValueConfig.ts";
import VervConfig from "@/models/configs/verv/VervConfig.ts";

const apiPrefix = { pathPrefix: "" };

export function setBackendAddress(url: string) {
  apiPrefix.pathPrefix = url;
}

const fallbackErrorConverting = "error during conversion";

export async function ListServices(req: ListConfigsRequest): Promise<ConfigList> {
  return MatreshkaBeAPI.ListConfigs(req, apiPrefix).then((r: ListConfigsResponse) => {
    const servicesInfo: ConfigBase[] = [];
    if (!r.configs) {
      throw { message: "invalid contract" };
    }

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
  version: string
): Promise<ConfigWithContent> {
  const req = {
    configName: configName,
    version: version,
  } as GetConfigNodeRequest;

  return MatreshkaBeAPI.GetConfigNodes(req, apiPrefix).then((res) => {
    if (!res.root) {
      throw { message: "Empty env config root" };
    }

    const cfg = new ConfigWithContent(configName);

    switch (cfg.type) {
      case ConfigTypePrefix.verv:
        cfg.content = new VervConfig(res.root);
        break;
      default:
        // TODO
        cfg.content = new KeyValueConfigContent();
    }

    return cfg;
  });
}

export async function PatchConfig(cfg: ConfigWithContent): Promise<ConfigWithContent> {
  if (!cfg.isChanged()) return cfg;

  const changeList = cfg.getChanges();

  const req: PatchConfigRequest = {
    configName: cfg.name,
    version: cfg.selectedVersion,
    changes: changeList.map((n) => {
      return {
        name: n.envName,
        value: n.newValue,
      } as Node;
    }),
  } as PatchConfigRequest;

  return MatreshkaBeAPI.PatchConfig(req, apiPrefix).then(() => {
    return GetConfigNodes(cfg.getMatreshkaName(), cfg.selectedVersion);
  });
}

export async function CreateConfig(name: string) {
  const newCfg: CreateConfigRequest = {
    configName: encodeURIComponent(name),
  };

  return MatreshkaBeAPI.CreateConfig(newCfg, apiPrefix);
}
