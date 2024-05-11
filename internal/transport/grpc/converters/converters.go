package converters

//import (
//	"fmt"
//	"time"
//
//	"github.com/godverv/matreshka"
//	"github.com/godverv/matreshka/api"
//	"github.com/godverv/matreshka/resources"
//
//	"github.com/godverv/matreshka-be/pkg/matreshka_api"
//)
//
//func cfgToProto(cfg *matreshka.AppConfig) *matreshka_api.Config {
//	if cfg == nil {
//		return nil
//	}
//
//	out := &matreshka_api.Config{
//		Resources:   make([]*matreshka_api.Config_Resource, len(cfg.Resources)),
//		Environment: make([]*matreshka_api.Config_Environment, 0, len(cfg.Environment)),
//	}
//
//	out.AppConfig = &matreshka_api.Config_AppConfig{
//		Name:            cfg.Name,
//		Version:         cfg.Version,
//		StartupDuration: cfg.StartupDuration.String(),
//	}
//
//	for i, r := range cfg.Resources {
//		out.Resources[i] = &matreshka_api.Config_Resource{
//			ResourceType:     resourceToProtoEnum[r.GetType()],
//			ConnectionString: "TODO", // TODO
//		}
//	}
//
//	for k, v := range cfg.Environment {
//		out.Environment = append(out.Environment, &matreshka_api.Config_Environment{
//			Key:   k,
//			Value: fmt.Sprintf("%s", v),
//		})
//	}
//
//	return out
//}
//
//var resourceToProtoEnum = map[string]matreshka_api.Config_Resource_Type{
//	resources.PostgresResourceName: matreshka_api.Config_Resource_postgres,
//	resources.RedisResourceName:    matreshka_api.Config_Resource_redis,
//	resources.TelegramResourceName: matreshka_api.Config_Resource_telegram,
//	resources.GrpcResourceName:     matreshka_api.Config_Resource_grpc,
//}
//
//func ()  {
//	{
//		appConfig := req.GetConfig().GetAppConfig()
//
//		c.cfg.AppInfo = matreshka.AppInfo{
//			Name:    appConfig.Name,
//			Version: appConfig.Version,
//		}
//
//		c.cfg.StartupDuration, _ = time.ParseDuration(appConfig.StartupDuration)
//	}
//
//	{
//		servers := req.GetConfig().GetApi()
//		c.cfg.Servers = make(matreshka.Servers, len(servers))
//
//		for i, item := range servers {
//			switch item.ApiType {
//			case matreshka_api.Config_Api_grpc:
//				c.cfg.Servers[i] = &api.GRPC{
//					Name: api.Name(item.MakoshName),
//				}
//			case matreshka_api.Config_Api_rest:
//				c.cfg.Servers[i] = &api.Rest{
//					Name: api.Name(item.MakoshName),
//				}
//			default:
//				c.cfg.Servers[i] = &api.Unknown{
//					Name: api.Name(item.MakoshName),
//				}
//			}
//		}
//	}
//
//	{
//		env := req.GetConfig().GetEnvironment()
//		c.cfg.Environment = make(map[string]interface{}, len(env))
//		for _, item := range env {
//			c.cfg.Environment[item.Key] = item.Value
//		}
//	}
//}
