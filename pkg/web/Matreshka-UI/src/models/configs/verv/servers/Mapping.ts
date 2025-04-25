import { Node } from "@vervstack/matreshka";

import Server, { FsHandler, GrpcHandler } from "@/models/configs/verv/servers/Servers.ts";
import { ConfigValue, extractStringValue } from "@/models/shared/Common.ts";

export function mapServer(root: Node): Server[] {
  if (!root.innerNodes) {
    throw { message: "Empty server node" };
  }

  const servers: Server[] = [];

  root.innerNodes.map((n) => {
    if (!n.name) {
      return;
    }

    const parts = n.name.split("_");

    const server: Server = new Server(parts[1]);

    if (n.innerNodes) {
      n.innerNodes.map((subNod) => {
        if (n.name) extractServerInfo(server, subNod, n.name);
      });
    }

    servers.push(server);
  });

  return servers;
}

function extractServerInfo(trg: Server, node: Node, rootPrefix: string) {
  if (!node.innerNodes || !node.name) {
    return;
  }

  const path = node.name.substring(rootPrefix.length + 1);
  switch (path) {
    case "/{GRPC}":
      trg.grpc.push(extractGrpcHandler(node.innerNodes, node.name));
      break;
    case "/{FS}":
      trg.fs.push(extractFsHandler(node.innerNodes, node.name));

      break;
    case "PORT":
      trg.port = new ConfigValue<number>(node.name, Number(node.value));
      break;
    default:
    // TODO http сервер
  }
}

function extractGrpcHandler(nodes: Node[], rootPrefix: string): GrpcHandler {
  const gh = new GrpcHandler();
  nodes.map((n) => {
    if (!n.name) {
      return;
    }

    const part = n.name.substring(rootPrefix.length + 1);
    switch (part) {
      case "GATEWAY":
        gh.gateway = extractStringValue(n);
        break;
      case "MODULE":
        gh.module = extractStringValue(n);
        break;
    }
  });

  return gh;
}

function extractFsHandler(nodes: Node[], rootPrefix: string): FsHandler {
  const fsH = new FsHandler();
  nodes.map((n) => {
    if (!n.name) {
      return;
    }

    const part = n.name.substring(rootPrefix.length + 1);
    switch (part) {
      case "DIST":
        fsH.dist = extractStringValue(n);
    }
  });

  return fsH;
}
