package converters

import (
	"github.com/godverv/matreshka/resources"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

var toProtoResourceMap = map[string]func(resources.Resource) *matreshka_api.Resource_Config{
	resources.SqliteResourceName:   ToProtoSqliteResource,
	resources.GrpcResourceName:     ToProtoGrpcResource,
	resources.PostgresResourceName: ToProtoPostgresResource,
	resources.RedisResourceName:    ToProtoRedisResource,
	resources.TelegramResourceName: ToProtoTelegramResource,
}

var fromProtoMatreshkaResourceTypeMap = map[matreshka_api.Resource_Type]string{
	matreshka_api.Resource_UnknownResourceType:  "unknown",
	matreshka_api.Resource_PostgresResourceType: resources.PostgresResourceName,
	matreshka_api.Resource_RedisResourceType:    resources.RedisResourceName,
	matreshka_api.Resource_SqliteResourceType:   resources.SqliteResourceName,
	matreshka_api.Resource_GrpcResourceType:     resources.GrpcResourceName,
	matreshka_api.Resource_TelegramResourceType: resources.TelegramResourceName,
}

var toProtoMatreshkaResourceTypeMap = map[string]matreshka_api.Resource_Type{
	"unknown":                      matreshka_api.Resource_UnknownResourceType,
	resources.PostgresResourceName: matreshka_api.Resource_PostgresResourceType,
	resources.RedisResourceName:    matreshka_api.Resource_RedisResourceType,
	resources.SqliteResourceName:   matreshka_api.Resource_SqliteResourceType,
	resources.GrpcResourceName:     matreshka_api.Resource_GrpcResourceType,
	resources.TelegramResourceName: matreshka_api.Resource_TelegramResourceType,
}

func ToProtoResources(rsc []resources.Resource) []*matreshka_api.Resource {
	out := make([]*matreshka_api.Resource, 0, len(rsc))

	for _, r := range rsc {
		mapper := toProtoResourceMap[r.GetType()]
		if mapper == nil {
			mapper = ToProtoUnknown
		}

		mapper(r)

		res := &matreshka_api.Resource{
			Name:           r.GetName(),
			ResourceConfig: mapper(r),
		}

		res.ResourceType = toProtoMatreshkaResourceTypeMap[r.GetType()]

		out = append(out, res)
	}

	return out
}
func ToProtoUnknown(in resources.Resource) *matreshka_api.Resource_Config {
	res, _ := in.(*resources.Unknown)

	return &matreshka_api.Resource_Config{
		Resource: &matreshka_api.Resource_Config_Unknown{
			Unknown: &matreshka_api.Resource_Unknown{
				Environment: res.Content,
			},
		},
	}
}
func ToProtoPostgresResource(in resources.Resource) *matreshka_api.Resource_Config {
	res, _ := in.(*resources.Postgres)

	return &matreshka_api.Resource_Config{
		Resource: &matreshka_api.Resource_Config_Postgres{
			Postgres: &matreshka_api.Resource_Postgres{
				Host: res.Host,
				Port: uint32(res.Port),

				DbName:   res.DbName,
				UserName: res.User,

				Pwd: res.Pwd,
			},
		},
	}
}
func ToProtoRedisResource(in resources.Resource) *matreshka_api.Resource_Config {
	res, _ := in.(*resources.Redis)

	return &matreshka_api.Resource_Config{
		Resource: &matreshka_api.Resource_Config_Redis{
			Redis: &matreshka_api.Resource_Redis{
				Host: res.Host,
				Port: uint32(res.Port),

				User: res.User,
				Pwd:  res.Pwd,

				Db: int32(res.Db),
			},
		},
	}
}
func ToProtoSqliteResource(in resources.Resource) *matreshka_api.Resource_Config {
	res, _ := in.(*resources.Sqlite)

	return &matreshka_api.Resource_Config{
		Resource: &matreshka_api.Resource_Config_Sqlite{
			Sqlite: &matreshka_api.Resource_Sqlite{
				Path: res.Path,
			},
		},
	}
}
func ToProtoGrpcResource(in resources.Resource) *matreshka_api.Resource_Config {
	res, _ := in.(*resources.GRPC)

	return &matreshka_api.Resource_Config{
		Resource: &matreshka_api.Resource_Config_Grpc{
			Grpc: &matreshka_api.Resource_Grpc{
				ConnectionString: res.ConnectionString,
				Module:           res.Module,
			},
		},
	}
}
func ToProtoTelegramResource(in resources.Resource) *matreshka_api.Resource_Config {
	res, _ := in.(*resources.Telegram)

	return &matreshka_api.Resource_Config{
		Resource: &matreshka_api.Resource_Config_Telegram{
			Telegram: &matreshka_api.Resource_Telegram{
				ApiKey: res.ApiKey,
			},
		},
	}
}

var fromProtoResourcesMap = map[matreshka_api.Resource_Type]func(config *matreshka_api.Resource) resources.Resource{
	matreshka_api.Resource_PostgresResourceType: FromProtoPostgresResource,
	matreshka_api.Resource_RedisResourceType:    FromProtoRedisResource,
	matreshka_api.Resource_SqliteResourceType:   FromProtoSqliteResource,
	matreshka_api.Resource_GrpcResourceType:     FromProtoGrpcResource,
	matreshka_api.Resource_TelegramResourceType: FromProtoTelegramResource,
}

func FromProtoResources(rsc []*matreshka_api.Resource) []resources.Resource {
	out := make([]resources.Resource, 0, len(rsc))

	for _, r := range rsc {
		mapper := fromProtoResourcesMap[r.GetResourceType()]
		if mapper == nil {
			mapper = FromProtoUnknownResource
		}

		out = append(out, mapper(r))
	}

	return out
}

func FromProtoUnknownResource(in *matreshka_api.Resource) resources.Resource {
	res := in.GetResourceConfig().GetUnknown()

	return &resources.Unknown{
		Name:    resources.Name(in.GetName()),
		Content: res.GetEnvironment(),
	}
}
func FromProtoPostgresResource(in *matreshka_api.Resource) resources.Resource {
	res := in.GetResourceConfig().GetPostgres()

	return &resources.Postgres{
		Name:   resources.Name(in.Name),
		Host:   res.GetHost(),
		Port:   uint64(res.Port),
		User:   res.GetUserName(),
		Pwd:    res.GetPwd(),
		DbName: res.GetDbName(),
	}
}
func FromProtoRedisResource(in *matreshka_api.Resource) resources.Resource {
	res := in.GetResourceConfig().GetRedis()

	return &resources.Redis{
		Name: resources.Name(in.Name),
		Host: res.GetHost(),
		Port: uint16(res.GetPort()),
		User: res.GetUser(),
		Pwd:  res.GetPwd(),
		Db:   int(res.GetDb()),
	}
}
func FromProtoSqliteResource(in *matreshka_api.Resource) resources.Resource {
	res := in.GetResourceConfig().GetSqlite()

	return &resources.Sqlite{
		Name: resources.Name(in.GetName()),
		Path: res.GetPath(),
	}
}
func FromProtoGrpcResource(in *matreshka_api.Resource) resources.Resource {
	res := in.GetResourceConfig().GetGrpc()

	return &resources.GRPC{
		Name:             resources.Name(in.Name),
		ConnectionString: res.GetConnectionString(),
		Module:           res.GetModule(),
	}
}
func FromProtoTelegramResource(in *matreshka_api.Resource) resources.Resource {
	res := in.GetResourceConfig().GetTelegram()

	return &resources.Telegram{
		Name:   resources.Name(in.GetName()),
		ApiKey: res.GetApiKey(),
	}
}
