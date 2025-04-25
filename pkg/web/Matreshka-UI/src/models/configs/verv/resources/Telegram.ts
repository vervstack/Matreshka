import { Node } from "@vervstack/matreshka";

import { Change } from "@/models/configs/Change.ts";
import DataSource from "@/models/configs/verv/resources/Resource.ts";
import { ResourceType } from "@/models/configs/verv/resources/ResourceTypes.ts";
import { ConfigValue, extractStringValue } from "@/models/shared/Common.ts";

export class Telegram extends DataSource {
  apiKey: ConfigValue<string> = new ConfigValue("", "");

  constructor(resourceName: string) {
    super(resourceName, ResourceType.Telegram);
  }

  rollback(): void {
    this.apiKey.rollback();
  }

  getChanges(): Change[] {
    const changes: Change[] = [];

    changes.push(...this.apiKey.getChanges());

    return changes;
  }
}

export function mapTelegram(root: Node): Telegram {
  if (!root.name) {
    throw { message: "No data to parse telegram" };
  }
  const tg = new Telegram(root.name.slice(root.name.indexOf("TELEGRAM")).toLowerCase());

  root.innerNodes?.map((n) => {
    if (!n.name || !root.name) {
      return;
    }

    const fieldName = n.name.slice(root.name.length + 1);
    switch (fieldName) {
      case "API-KEY":
        tg.apiKey = extractStringValue(n);
    }
  });

  return tg;
}
