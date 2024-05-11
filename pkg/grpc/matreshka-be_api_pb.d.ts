import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from '../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"


export class Resource extends jspb.Message {
  getConn(): Resource.Connection | undefined;
  setConn(value?: Resource.Connection): Resource;
  hasConn(): boolean;
  clearConn(): Resource;

  getName(): string;
  setName(value: string): Resource;

  getResourceType(): Resource.Type;
  setResourceType(value: Resource.Type): Resource;

  getResourceConfig(): Resource.Config | undefined;
  setResourceConfig(value?: Resource.Config): Resource;
  hasResourceConfig(): boolean;
  clearResourceConfig(): Resource;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Resource.AsObject;
  static toObject(includeInstance: boolean, msg: Resource): Resource.AsObject;
  static serializeBinaryToWriter(message: Resource, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Resource;
  static deserializeBinaryFromReader(message: Resource, reader: jspb.BinaryReader): Resource;
}

export namespace Resource {
  export type AsObject = {
    conn?: Resource.Connection.AsObject,
    name: string,
    resourceType: Resource.Type,
    resourceConfig?: Resource.Config.AsObject,
  }

  export class Unknown extends jspb.Message {
    getEnvironmentMap(): jspb.Map<string, string>;
    clearEnvironmentMap(): Unknown;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Unknown.AsObject;
    static toObject(includeInstance: boolean, msg: Unknown): Unknown.AsObject;
    static serializeBinaryToWriter(message: Unknown, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Unknown;
    static deserializeBinaryFromReader(message: Unknown, reader: jspb.BinaryReader): Unknown;
  }

  export namespace Unknown {
    export type AsObject = {
      environmentMap: Array<[string, string]>,
    }
  }


  export class Postgres extends jspb.Message {
    getAddress(): string;
    setAddress(value: string): Postgres;

    getPort(): number;
    setPort(value: number): Postgres;

    getDbName(): string;
    setDbName(value: string): Postgres;

    getUserName(): string;
    setUserName(value: string): Postgres;

    getPwd(): string;
    setPwd(value: string): Postgres;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Postgres.AsObject;
    static toObject(includeInstance: boolean, msg: Postgres): Postgres.AsObject;
    static serializeBinaryToWriter(message: Postgres, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Postgres;
    static deserializeBinaryFromReader(message: Postgres, reader: jspb.BinaryReader): Postgres;
  }

  export namespace Postgres {
    export type AsObject = {
      address: string,
      port: number,
      dbName: string,
      userName: string,
      pwd: string,
    }
  }


  export class Redis extends jspb.Message {
    getHost(): string;
    setHost(value: string): Redis;

    getPort(): number;
    setPort(value: number): Redis;

    getUser(): string;
    setUser(value: string): Redis;

    getPwd(): string;
    setPwd(value: string): Redis;

    getDb(): number;
    setDb(value: number): Redis;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Redis.AsObject;
    static toObject(includeInstance: boolean, msg: Redis): Redis.AsObject;
    static serializeBinaryToWriter(message: Redis, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Redis;
    static deserializeBinaryFromReader(message: Redis, reader: jspb.BinaryReader): Redis;
  }

  export namespace Redis {
    export type AsObject = {
      host: string,
      port: number,
      user: string,
      pwd: string,
      db: number,
    }
  }


  export class Sqlite extends jspb.Message {
    getPath(): string;
    setPath(value: string): Sqlite;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Sqlite.AsObject;
    static toObject(includeInstance: boolean, msg: Sqlite): Sqlite.AsObject;
    static serializeBinaryToWriter(message: Sqlite, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Sqlite;
    static deserializeBinaryFromReader(message: Sqlite, reader: jspb.BinaryReader): Sqlite;
  }

  export namespace Sqlite {
    export type AsObject = {
      path: string,
    }
  }


  export class Grpc extends jspb.Message {
    getConnectionString(): string;
    setConnectionString(value: string): Grpc;

    getModule(): string;
    setModule(value: string): Grpc;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Grpc.AsObject;
    static toObject(includeInstance: boolean, msg: Grpc): Grpc.AsObject;
    static serializeBinaryToWriter(message: Grpc, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Grpc;
    static deserializeBinaryFromReader(message: Grpc, reader: jspb.BinaryReader): Grpc;
  }

  export namespace Grpc {
    export type AsObject = {
      connectionString: string,
      module: string,
    }
  }


  export class Telegram extends jspb.Message {
    getApiKey(): string;
    setApiKey(value: string): Telegram;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Telegram.AsObject;
    static toObject(includeInstance: boolean, msg: Telegram): Telegram.AsObject;
    static serializeBinaryToWriter(message: Telegram, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Telegram;
    static deserializeBinaryFromReader(message: Telegram, reader: jspb.BinaryReader): Telegram;
  }

  export namespace Telegram {
    export type AsObject = {
      apiKey: string,
    }
  }


  export class Connection extends jspb.Message {
    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Connection.AsObject;
    static toObject(includeInstance: boolean, msg: Connection): Connection.AsObject;
    static serializeBinaryToWriter(message: Connection, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Connection;
    static deserializeBinaryFromReader(message: Connection, reader: jspb.BinaryReader): Connection;
  }

  export namespace Connection {
    export type AsObject = {
    }
  }


  export class Config extends jspb.Message {
    getUnknown(): Resource.Unknown | undefined;
    setUnknown(value?: Resource.Unknown): Config;
    hasUnknown(): boolean;
    clearUnknown(): Config;

    getPostgres(): Resource.Postgres | undefined;
    setPostgres(value?: Resource.Postgres): Config;
    hasPostgres(): boolean;
    clearPostgres(): Config;

    getRedis(): Resource.Redis | undefined;
    setRedis(value?: Resource.Redis): Config;
    hasRedis(): boolean;
    clearRedis(): Config;

    getSqlite(): Resource.Sqlite | undefined;
    setSqlite(value?: Resource.Sqlite): Config;
    hasSqlite(): boolean;
    clearSqlite(): Config;

    getGrpc(): Resource.Grpc | undefined;
    setGrpc(value?: Resource.Grpc): Config;
    hasGrpc(): boolean;
    clearGrpc(): Config;

    getTelegram(): Resource.Telegram | undefined;
    setTelegram(value?: Resource.Telegram): Config;
    hasTelegram(): boolean;
    clearTelegram(): Config;

    getResourceCase(): Config.ResourceCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Config.AsObject;
    static toObject(includeInstance: boolean, msg: Config): Config.AsObject;
    static serializeBinaryToWriter(message: Config, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Config;
    static deserializeBinaryFromReader(message: Config, reader: jspb.BinaryReader): Config;
  }

  export namespace Config {
    export type AsObject = {
      unknown?: Resource.Unknown.AsObject,
      postgres?: Resource.Postgres.AsObject,
      redis?: Resource.Redis.AsObject,
      sqlite?: Resource.Sqlite.AsObject,
      grpc?: Resource.Grpc.AsObject,
      telegram?: Resource.Telegram.AsObject,
    }

    export enum ResourceCase { 
      RESOURCE_NOT_SET = 0,
      UNKNOWN = 4,
      POSTGRES = 5,
      REDIS = 6,
      SQLITE = 7,
      GRPC = 8,
      TELEGRAM = 9,
    }
  }


  export enum Type { 
    UNKNOWNRESOURCETYPE = 0,
    POSTGRESRESOURCETYPE = 1,
    REDISRESOURCETYPE = 2,
    SQLITERESOURCETYPE = 3,
    GRPCRESOURCETYPE = 4,
    TELEGRAMRESOURCETYPE = 5,
  }
}

export class Server extends jspb.Message {
  getSwaggerLink(): string;
  setSwaggerLink(value: string): Server;

  getMakoshName(): string;
  setMakoshName(value: string): Server;

  getServer(): Server.Config | undefined;
  setServer(value?: Server.Config): Server;
  hasServer(): boolean;
  clearServer(): Server;

  getType(): Server.Type;
  setType(value: Server.Type): Server;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Server.AsObject;
  static toObject(includeInstance: boolean, msg: Server): Server.AsObject;
  static serializeBinaryToWriter(message: Server, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Server;
  static deserializeBinaryFromReader(message: Server, reader: jspb.BinaryReader): Server;
}

export namespace Server {
  export type AsObject = {
    swaggerLink: string,
    makoshName: string,
    server?: Server.Config.AsObject,
    type: Server.Type,
  }

  export class Unknown extends jspb.Message {
    getValuesMap(): jspb.Map<string, string>;
    clearValuesMap(): Unknown;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Unknown.AsObject;
    static toObject(includeInstance: boolean, msg: Unknown): Unknown.AsObject;
    static serializeBinaryToWriter(message: Unknown, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Unknown;
    static deserializeBinaryFromReader(message: Unknown, reader: jspb.BinaryReader): Unknown;
  }

  export namespace Unknown {
    export type AsObject = {
      valuesMap: Array<[string, string]>,
    }
  }


  export class Grpc extends jspb.Message {
    getPort(): number;
    setPort(value: number): Grpc;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Grpc.AsObject;
    static toObject(includeInstance: boolean, msg: Grpc): Grpc.AsObject;
    static serializeBinaryToWriter(message: Grpc, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Grpc;
    static deserializeBinaryFromReader(message: Grpc, reader: jspb.BinaryReader): Grpc;
  }

  export namespace Grpc {
    export type AsObject = {
      port: number,
    }
  }


  export class Rest extends jspb.Message {
    getPort(): number;
    setPort(value: number): Rest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Rest.AsObject;
    static toObject(includeInstance: boolean, msg: Rest): Rest.AsObject;
    static serializeBinaryToWriter(message: Rest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Rest;
    static deserializeBinaryFromReader(message: Rest, reader: jspb.BinaryReader): Rest;
  }

  export namespace Rest {
    export type AsObject = {
      port: number,
    }
  }


  export class Config extends jspb.Message {
    getUnknown(): Server.Unknown | undefined;
    setUnknown(value?: Server.Unknown): Config;
    hasUnknown(): boolean;
    clearUnknown(): Config;

    getGrpc(): Server.Grpc | undefined;
    setGrpc(value?: Server.Grpc): Config;
    hasGrpc(): boolean;
    clearGrpc(): Config;

    getRest(): Server.Rest | undefined;
    setRest(value?: Server.Rest): Config;
    hasRest(): boolean;
    clearRest(): Config;

    getServerCase(): Config.ServerCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Config.AsObject;
    static toObject(includeInstance: boolean, msg: Config): Config.AsObject;
    static serializeBinaryToWriter(message: Config, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Config;
    static deserializeBinaryFromReader(message: Config, reader: jspb.BinaryReader): Config;
  }

  export namespace Config {
    export type AsObject = {
      unknown?: Server.Unknown.AsObject,
      grpc?: Server.Grpc.AsObject,
      rest?: Server.Rest.AsObject,
    }

    export enum ServerCase { 
      SERVER_NOT_SET = 0,
      UNKNOWN = 3,
      GRPC = 4,
      REST = 5,
    }
  }


  export enum Type { 
    UNKNOWNSERVERTYPE = 0,
    GRPCSERVERTYPE = 1,
    RESTSERVERTYPE = 2,
  }
}

export class Config extends jspb.Message {
  getAppConfig(): Config.AppConfig | undefined;
  setAppConfig(value?: Config.AppConfig): Config;
  hasAppConfig(): boolean;
  clearAppConfig(): Config;

  getResourcesList(): Array<Resource>;
  setResourcesList(value: Array<Resource>): Config;
  clearResourcesList(): Config;
  addResources(value?: Resource, index?: number): Resource;

  getApiList(): Array<Server>;
  setApiList(value: Array<Server>): Config;
  clearApiList(): Config;
  addApi(value?: Server, index?: number): Server;

  getEnvironmentMap(): jspb.Map<string, string>;
  clearEnvironmentMap(): Config;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Config.AsObject;
  static toObject(includeInstance: boolean, msg: Config): Config.AsObject;
  static serializeBinaryToWriter(message: Config, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Config;
  static deserializeBinaryFromReader(message: Config, reader: jspb.BinaryReader): Config;
}

export namespace Config {
  export type AsObject = {
    appConfig?: Config.AppConfig.AsObject,
    resourcesList: Array<Resource.AsObject>,
    apiList: Array<Server.AsObject>,
    environmentMap: Array<[string, string]>,
  }

  export class AppConfig extends jspb.Message {
    getName(): string;
    setName(value: string): AppConfig;

    getVersion(): string;
    setVersion(value: string): AppConfig;

    getStartupDurationSec(): number;
    setStartupDurationSec(value: number): AppConfig;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AppConfig.AsObject;
    static toObject(includeInstance: boolean, msg: AppConfig): AppConfig.AsObject;
    static serializeBinaryToWriter(message: AppConfig, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AppConfig;
    static deserializeBinaryFromReader(message: AppConfig, reader: jspb.BinaryReader): AppConfig;
  }

  export namespace AppConfig {
    export type AsObject = {
      name: string,
      version: string,
      startupDurationSec: number,
    }
  }

}

export class ApiVersion extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ApiVersion.AsObject;
  static toObject(includeInstance: boolean, msg: ApiVersion): ApiVersion.AsObject;
  static serializeBinaryToWriter(message: ApiVersion, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ApiVersion;
  static deserializeBinaryFromReader(message: ApiVersion, reader: jspb.BinaryReader): ApiVersion;
}

export namespace ApiVersion {
  export type AsObject = {
  }

  export class Request extends jspb.Message {
    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
  }

  export namespace Request {
    export type AsObject = {
    }
  }


  export class Response extends jspb.Message {
    getVersion(): string;
    setVersion(value: string): Response;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
  }

  export namespace Response {
    export type AsObject = {
      version: string,
    }
  }

}

export class PatchConfig extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PatchConfig.AsObject;
  static toObject(includeInstance: boolean, msg: PatchConfig): PatchConfig.AsObject;
  static serializeBinaryToWriter(message: PatchConfig, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PatchConfig;
  static deserializeBinaryFromReader(message: PatchConfig, reader: jspb.BinaryReader): PatchConfig;
}

export namespace PatchConfig {
  export type AsObject = {
  }

  export class Request extends jspb.Message {
    getConfig(): Config | undefined;
    setConfig(value?: Config): Request;
    hasConfig(): boolean;
    clearConfig(): Request;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
  }

  export namespace Request {
    export type AsObject = {
      config?: Config.AsObject,
    }
  }


  export class Response extends jspb.Message {
    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
  }

  export namespace Response {
    export type AsObject = {
    }
  }

}

export class GetConfig extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetConfig.AsObject;
  static toObject(includeInstance: boolean, msg: GetConfig): GetConfig.AsObject;
  static serializeBinaryToWriter(message: GetConfig, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetConfig;
  static deserializeBinaryFromReader(message: GetConfig, reader: jspb.BinaryReader): GetConfig;
}

export namespace GetConfig {
  export type AsObject = {
  }

  export class Request extends jspb.Message {
    getServiceName(): string;
    setServiceName(value: string): Request;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
  }

  export namespace Request {
    export type AsObject = {
      serviceName: string,
    }
  }


  export class Response extends jspb.Message {
    getConfig(): Config | undefined;
    setConfig(value?: Config): Response;
    hasConfig(): boolean;
    clearConfig(): Response;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
  }

  export namespace Response {
    export type AsObject = {
      config?: Config.AsObject,
    }
  }

}

export class GetConfigRaw extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetConfigRaw.AsObject;
  static toObject(includeInstance: boolean, msg: GetConfigRaw): GetConfigRaw.AsObject;
  static serializeBinaryToWriter(message: GetConfigRaw, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetConfigRaw;
  static deserializeBinaryFromReader(message: GetConfigRaw, reader: jspb.BinaryReader): GetConfigRaw;
}

export namespace GetConfigRaw {
  export type AsObject = {
  }

  export class Request extends jspb.Message {
    getServiceName(): string;
    setServiceName(value: string): Request;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
  }

  export namespace Request {
    export type AsObject = {
      serviceName: string,
    }
  }


  export class Response extends jspb.Message {
    getConfig(): Uint8Array | string;
    getConfig_asU8(): Uint8Array;
    getConfig_asB64(): string;
    setConfig(value: Uint8Array | string): Response;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
  }

  export namespace Response {
    export type AsObject = {
      config: Uint8Array | string,
    }
  }

}

export class PatchConfigRaw extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PatchConfigRaw.AsObject;
  static toObject(includeInstance: boolean, msg: PatchConfigRaw): PatchConfigRaw.AsObject;
  static serializeBinaryToWriter(message: PatchConfigRaw, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PatchConfigRaw;
  static deserializeBinaryFromReader(message: PatchConfigRaw, reader: jspb.BinaryReader): PatchConfigRaw;
}

export namespace PatchConfigRaw {
  export type AsObject = {
  }

  export class Request extends jspb.Message {
    getRaw(): Uint8Array | string;
    getRaw_asU8(): Uint8Array;
    getRaw_asB64(): string;
    setRaw(value: Uint8Array | string): Request;

    getServiceName(): string;
    setServiceName(value: string): Request;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
  }

  export namespace Request {
    export type AsObject = {
      raw: Uint8Array | string,
      serviceName: string,
    }
  }


  export class Response extends jspb.Message {
    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
  }

  export namespace Response {
    export type AsObject = {
    }
  }

}

export class ListRequest extends jspb.Message {
  getLimit(): number;
  setLimit(value: number): ListRequest;

  getOffset(): number;
  setOffset(value: number): ListRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRequest): ListRequest.AsObject;
  static serializeBinaryToWriter(message: ListRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRequest;
  static deserializeBinaryFromReader(message: ListRequest, reader: jspb.BinaryReader): ListRequest;
}

export namespace ListRequest {
  export type AsObject = {
    limit: number,
    offset: number,
  }
}

export class ListConfigs extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListConfigs.AsObject;
  static toObject(includeInstance: boolean, msg: ListConfigs): ListConfigs.AsObject;
  static serializeBinaryToWriter(message: ListConfigs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListConfigs;
  static deserializeBinaryFromReader(message: ListConfigs, reader: jspb.BinaryReader): ListConfigs;
}

export namespace ListConfigs {
  export type AsObject = {
  }

  export class Request extends jspb.Message {
    getListrequest(): ListRequest | undefined;
    setListrequest(value?: ListRequest): Request;
    hasListrequest(): boolean;
    clearListrequest(): Request;

    getServicename(): string;
    setServicename(value: string): Request;
    hasServicename(): boolean;
    clearServicename(): Request;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
  }

  export namespace Request {
    export type AsObject = {
      listrequest?: ListRequest.AsObject,
      servicename?: string,
    }

    export enum ServicenameCase { 
      _SERVICENAME_NOT_SET = 0,
      SERVICENAME = 2,
    }
  }


  export class Response extends jspb.Message {
    getServicesList(): Array<ListConfigs.Response.ServiceInfo>;
    setServicesList(value: Array<ListConfigs.Response.ServiceInfo>): Response;
    clearServicesList(): Response;
    addServices(value?: ListConfigs.Response.ServiceInfo, index?: number): ListConfigs.Response.ServiceInfo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
  }

  export namespace Response {
    export type AsObject = {
      servicesList: Array<ListConfigs.Response.ServiceInfo.AsObject>,
    }

    export class ServiceInfo extends jspb.Message {
      getName(): string;
      setName(value: string): ServiceInfo;

      serializeBinary(): Uint8Array;
      toObject(includeInstance?: boolean): ServiceInfo.AsObject;
      static toObject(includeInstance: boolean, msg: ServiceInfo): ServiceInfo.AsObject;
      static serializeBinaryToWriter(message: ServiceInfo, writer: jspb.BinaryWriter): void;
      static deserializeBinary(bytes: Uint8Array): ServiceInfo;
      static deserializeBinaryFromReader(message: ServiceInfo, reader: jspb.BinaryReader): ServiceInfo;
    }

    export namespace ServiceInfo {
      export type AsObject = {
        name: string,
      }
    }

  }

}

