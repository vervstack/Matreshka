import { Node } from "@vervstack/matreshka";

import { Change } from "@/models/configs/Change.ts";
import DataSource from "@/models/configs/verv/resources/Resource.ts";
import { ResourceType } from "@/models/configs/verv/resources/ResourceTypes.ts";
import { ConfigValue, extractStringValue } from "@/models/shared/Common.ts";

export class GrpcClient extends DataSource {
  connectionString: ConfigValue<string> = new ConfigValue("", "");
  module: ConfigValue<string> = new ConfigValue("", "");

  constructor(resourceName: string) {
    super(resourceName, ResourceType.Grpc);
  }

  rollback(): void {
    this.connectionString.rollback();
    this.module.rollback();
  }

  getChanges(): Change[] {
    const changes: Change[] = [];

    changes.push(...this.connectionString.getChanges());
    changes.push(...this.module.getChanges());

    return changes;
  }
}

export function mapGrpc(root: Node): GrpcClient {
  if (!root.name) {
    throw { message: "Can't parse grpc client config" };
  }

  const grpcClient = new GrpcClient(root.name.slice(root.name.indexOf("GRPC")).toLowerCase());

  root.innerNodes?.map((n) => {
    if (!n.name || !root.name) {
      return;
    }

    const fieldName = n.name.slice(root.name.length + 1);

    switch (fieldName) {
      case "CONNECTION-STRING":
        grpcClient.connectionString = extractStringValue(n);
        break;
      case "MODULE":
        grpcClient.module = extractStringValue(n);
        break;
    }
  });

  return grpcClient;
}
