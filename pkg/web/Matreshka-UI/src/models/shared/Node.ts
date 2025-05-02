import { Node } from "@vervstack/matreshka";

export default class EnvNode {
  name: string;
  value: string;
  children: EnvNode[] = [];

  constructor(name: string, value: string) {
    this.name = name;
    this.value = value;
  }
}

export function fromPbEnvNode(root: Node, parentPrefix = ""): EnvNode {
  let name = root.name || "";

  if (name.startsWith(parentPrefix)) {
    name = name.substring(parentPrefix.length);
  }

  while (name.startsWith("_")) {
    name = name.substring(1);
  }

  const node = new EnvNode(name, root.value || "");

  if (parentPrefix != "") {
    parentPrefix += "_" + name;
  } else {
    parentPrefix = name;
  }
  if (root.innerNodes) {
    root.innerNodes.map((child: Node) => {
      node.children.push(fromPbEnvNode(child, parentPrefix));
    });
  }
  return node;
}
