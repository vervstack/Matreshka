import { Node } from "@vervstack/matreshka";

import { Change } from "@/models/configs/Change.ts";

export type KeyMap = {
  [key: string]: any;
};

export class ConfigValue<T> {
  envName: string;
  value: T;

  private readonly originalName: string;
  private readonly originalValue: T;

  constructor(envName: string, value: T) {
    this.originalName = envName;
    this.originalValue = value;

    this.envName = envName;
    this.value = value;
  }

  getOriginalValue(): T {
    return this.originalValue;
  }

  getOriginalName(): string {
    return this.originalName;
  }

  isChanged(): boolean {
    return this.value != this.originalValue ||
      this.envName != this.originalName;
  }

  isNameChanged(): boolean {
    return this.envName != this.originalName;
  }

  isValueChanged(): boolean {
    return this.value != this.originalValue;
  }

  getChanges(): Change[] {
    const changes: Change[] = [];
    if (this.value != this.originalValue) {
      changes.push({
        envName: this.envName,
        newValue: this.value,
      } as Change);
    }

    if (this.envName != this.originalName) {
      changes.push({
        envName: this.originalName,
        newValue: "",
      });
    }

    return changes;
  }

  rollback() {
    this.value = this.originalValue;
  }

  rollbackName() {
    this.envName = this.originalName;
  }
}

export function extractStringValue(n: Node): ConfigValue<string> {
  return new ConfigValue<string>(n.name ?? "", n.value ?? "");
}

export function extractNumberValue(n: Node): ConfigValue<number> {
  return new ConfigValue(n.name ?? "", Number(n.value) ?? 0);
}

export function extractResourceType(node: Node, root: Node): string | undefined {
  if (!node.name || !root.name) {
    return;
  }

  let name = node.name.slice(root.name.length + 1);

  const resourceNameEndIdx = name.indexOf("_");
  if (resourceNameEndIdx > 0) {
    name = name.slice(resourceNameEndIdx);
  }
  name = name.toLowerCase();

  let resType = name;
  const resourceTypeNameEndIdx = resType.indexOf("-");
  if (resourceTypeNameEndIdx > 0) {
    resType = name.slice(0, resourceTypeNameEndIdx);
  }

  return resType;
}
