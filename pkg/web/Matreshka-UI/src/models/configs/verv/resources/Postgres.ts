import { Node } from "@vervstack/matreshka";

import { Change } from "@/models/configs/Change.ts";
import DataSource from "@/models/configs/verv/resources/Resource.ts";
import { ResourceType } from "@/models/configs/verv/resources/ResourceTypes.ts";
import { ConfigValue, extractNumberValue, extractStringValue } from "@/models/shared/Values.ts";

export class Postgres extends DataSource {
  host: ConfigValue<string> = new ConfigValue("", "");
  name: ConfigValue<string> = new ConfigValue("", "");
  port: ConfigValue<number> = new ConfigValue("", 0);
  user: ConfigValue<string> = new ConfigValue("", "");
  pwd: ConfigValue<string> = new ConfigValue("", "");
  ssl_mode: ConfigValue<string> = new ConfigValue("", "");

  constructor(resourceName: string) {
    super(resourceName, ResourceType.Postgres);
  }

  rollback(): void {
    this.host.rollback();
    this.name.rollback();
    this.port.rollback();
    this.user.rollback();
    this.pwd.rollback();
    this.ssl_mode.rollback();
  }

  getChanges(): Change[] {
    const changes: Change[] = [];

    changes.push(...this.host.getChanges());
    changes.push(...this.name.getChanges());
    changes.push(...this.port.getChanges());
    changes.push(...this.user.getChanges());
    changes.push(...this.pwd.getChanges());
    changes.push(...this.ssl_mode.getChanges());

    return changes;
  }
}

export function newPostgres(root: Node): DataSource {
  if (!root.name) {
    throw { message: "No data for postgres to map" };
  }

  const pg = new Postgres(root.name.slice(root.name.indexOf("POSTGRES")).toLowerCase());

  root.innerNodes?.map((n) => {
    if (!n.name || !root.name) {
      return;
    }

    const fieldName = n.name.slice(root.name.length + 1);
    switch (fieldName) {
      case "HOST":
        pg.host = extractStringValue(n);
        break;
      case "PORT":
        pg.port = extractNumberValue(n);
        break;
      case "USER":
        pg.user = extractStringValue(n);
        break;
      case "PWD":
        pg.pwd = extractStringValue(n);
        break;
      case "DB-NAME":
        pg.name = extractStringValue(n);
        break;
      case "SSL-MODE":
        pg.ssl_mode = extractStringValue(n);
        break;
    }
  });

  return pg;
}
