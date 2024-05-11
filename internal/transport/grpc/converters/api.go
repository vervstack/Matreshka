package converters

import (
	matreshkaServers "github.com/godverv/matreshka/api"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

var (
	toProtoServerMap = map[string]func(matreshkaServers.Api) *matreshka_api.Server_Config{
		matreshkaServers.GRPSServerType: ToProtoGrpcApi,
		matreshkaServers.RestServerType: ToProtoRestApi,
	}
	toProtoServerTypeMap = map[string]matreshka_api.Server_Type{
		matreshkaServers.GRPSServerType: matreshka_api.Server_GrpcServerType,
		matreshkaServers.RestServerType: matreshka_api.Server_RestServerType,
	}
)

func ToProtoApi(apis []matreshkaServers.Api) []*matreshka_api.Server {
	out := make([]*matreshka_api.Server, 0, len(apis))

	for _, r := range apis {
		mapper := toProtoServerMap[""]
		if mapper == nil {
			mapper = ToProtoUnknownServer
		}

		server := &matreshka_api.Server{
			SwaggerLink: "",

			MakoshName: r.GetName(),

			Server: mapper(r),
			Type:   toProtoServerTypeMap[r.GetType()],
		}

		out = append(out, server)
	}

	return out
}
func ToProtoUnknownServer(in matreshkaServers.Api) *matreshka_api.Server_Config {
	res, _ := in.(*matreshkaServers.Unknown)
	return &matreshka_api.Server_Config{
		Server: &matreshka_api.Server_Config_Unknown{
			Unknown: &matreshka_api.Server_Unknown{
				Values: res.Values,
			},
		},
	}
}
func ToProtoGrpcApi(in matreshkaServers.Api) *matreshka_api.Server_Config {
	return &matreshka_api.Server_Config{
		Server: &matreshka_api.Server_Config_Grpc{
			Grpc: &matreshka_api.Server_Grpc{
				Port: uint32(in.GetPort()),
			},
		},
	}
}
func ToProtoRestApi(in matreshkaServers.Api) *matreshka_api.Server_Config {
	return &matreshka_api.Server_Config{
		Server: &matreshka_api.Server_Config_Rest{
			Rest: &matreshka_api.Server_Rest{
				Port: uint32(in.GetPort()),
			},
		},
	}
}

var fromProtoApiMap = map[matreshka_api.Server_Type]func(config *matreshka_api.Server) matreshkaServers.Api{
	matreshka_api.Server_GrpcServerType: FromProtoGrpcApi,
	matreshka_api.Server_RestServerType: FromProtoRestApi,
}

func FromProtoApi(apis []*matreshka_api.Server) []matreshkaServers.Api {
	out := make([]matreshkaServers.Api, 0, len(apis))

	for _, r := range apis {
		mapper := fromProtoApiMap[r.GetType()]
		if mapper == nil {
			mapper = FromProtoUnknownApi
		}

		out = append(out)
	}

	return out
}
func FromProtoUnknownApi(in *matreshka_api.Server) matreshkaServers.Api {
	r := in.GetServer().GetUnknown()
	return &matreshkaServers.Unknown{
		Name:   matreshkaServers.Name(in.MakoshName),
		Values: r.GetValues(),
	}
}
func FromProtoGrpcApi(in *matreshka_api.Server) matreshkaServers.Api {
	r := in.GetServer().GetGrpc()
	return &matreshkaServers.GRPC{
		Name: matreshkaServers.Name(in.GetMakoshName()),
		Port: uint16(r.GetPort()),
	}
}
func FromProtoRestApi(in *matreshka_api.Server) matreshkaServers.Api {
	r := in.GetServer().GetRest()
	return &matreshkaServers.Rest{
		Name: matreshkaServers.Name(in.GetMakoshName()),
		Port: uint16(r.GetPort()),
	}
}
