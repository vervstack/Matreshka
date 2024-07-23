package tests

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/godverv/matreshka/environment"
	"github.com/godverv/matreshka/resources"
	"github.com/godverv/matreshka/servers"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

type PatchConfigSuite struct {
	suite.Suite

	ctx context.Context
}

func (s *PatchConfigSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *PatchConfigSuite) Test_PatchConfigEnvironment() {
	serviceName := s.T().Name()
	testEnv.create(s.T(), serviceName, fullConfigBytes)

	newConfig := getFullConfig(s.T())

	patchReq := &matreshka_api.PatchConfig_Request{
		ServiceName: serviceName,
	}

	// Change old environment variable
	{
		newConfig.Environment[0].Value = []int{50051}

		patchStr := fmt.Sprint(newConfig.Environment[0].Value)
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name:  "ENVIRONMENT_AVAILABLE-PORTS",
				Value: &patchStr,
			})
	}
	// Delete environment variable
	{
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name: "ENVIRONMENT_WELCOME-STRING",
			})
		newConfig.Environment = newConfig.Environment[:len(newConfig.Environment)-1]
	}
	// Add new environment variable
	{
		someValue := "rand val"
		valueType := string(environment.VariableTypeStr)

		newEnvVar := &environment.Variable{
			Name:  "new value",
			Value: someValue,
			Type:  environment.VariableTypeStr,
		}

		newConfig.Environment = append(newConfig.Environment, newEnvVar)

		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name:  "ENVIRONMENT_NEW-VALUE",
				Value: &someValue,
				InnerNodes: []*matreshka_api.Node{
					{
						Name:  "TYPE",
						Value: &valueType,
					},
				},
			},
		)
	}

	patchResp, err := testEnv.grpcClient.PatchConfig(s.ctx, patchReq)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), patchResp)

	patchedConfig := testEnv.get(s.T(), serviceName)

	sort.Slice(newConfig.Environment, func(i, j int) bool {
		return newConfig.Environment[i].Name < newConfig.Environment[j].Name
	})

	require.Equal(s.T(), patchedConfig, newConfig)
}

func (s *PatchConfigSuite) Test_PatchConfigServers() {
	serviceName := s.T().Name()
	testEnv.create(s.T(), serviceName, fullConfigBytes)

	newConfig := getFullConfig(s.T())

	patchReq := &matreshka_api.PatchConfig_Request{
		ServiceName: serviceName,
	}

	// Change old rpc server port
	{
		newConfig.Servers[0] = &servers.GRPC{
			Name: "grpc",
			Port: 5432,
		}

		portStr := newConfig.Servers[0].GetPortStr()
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name:  "SERVERS_GRPC_PORT",
				Value: &portStr,
			})
	}

	// Delete old rest server
	{
		newConfig.Servers = newConfig.Servers[:1]

		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name: "SERVERS_REST_PORT",
			})
	}

	// Add new rest server
	{
		newConfig.Servers = append(newConfig.Servers,
			&servers.Rest{
				Name: "rest_api_gateway",
				Port: 5678,
			})
		portStr := newConfig.Servers[1].GetPortStr()
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name:  "SERVERS_REST-API-GATEWAY_PORT",
				Value: &portStr,
			})
	}

	patchResp, err := testEnv.grpcClient.PatchConfig(s.ctx, patchReq)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), patchResp)

	patchedConfig := testEnv.get(s.T(), serviceName)

	sort.Slice(newConfig.Environment, func(i, j int) bool {
		return newConfig.Environment[i].Name < newConfig.Environment[j].Name
	})

	require.Equal(s.T(), patchedConfig, newConfig)
}

func (s *PatchConfigSuite) Test_PatchConfigDataSources() {
	serviceName := s.T().Name()
	testEnv.create(s.T(), serviceName, fullConfigBytes)

	newConfig := getFullConfig(s.T())

	patchReq := &matreshka_api.PatchConfig_Request{
		ServiceName: serviceName,
	}

	// Change old resource (pg) port
	{
		pg := newConfig.DataSources[1].(*resources.Postgres)
		pg.Port = 5433
		portStr := strconv.Itoa(int(pg.Port))
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name:  "DATA-SOURCES_POSTGRES_PORT",
				Value: &portStr,
			})
	}

	// Delete old resource (telegram) data source
	{
		newConfig.DataSources = newConfig.DataSources[:3]
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name: "DATA-SOURCES_TELEGRAM",
			})
	}

	// Add new telegram data source
	{

		tg := &resources.Telegram{
			Name:   "telegram_bot",
			ApiKey: "jjggwwkk",
		}
		newConfig.DataSources = append(newConfig.DataSources, tg)
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name:  "DATA-SOURCES_TELEGRAM-BOT_API-KEY",
				Value: &tg.ApiKey,
			})
	}

	patchResp, err := testEnv.grpcClient.PatchConfig(s.ctx, patchReq)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), patchResp)

	patchedConfig := testEnv.get(s.T(), serviceName)

	sort.Slice(newConfig.Environment, func(i, j int) bool {
		return newConfig.Environment[i].Name < newConfig.Environment[j].Name
	})

	require.Equal(s.T(), patchedConfig, newConfig)
}

func Test_PatchConfig(t *testing.T) {
	suite.Run(t, new(PatchConfigSuite))
}
