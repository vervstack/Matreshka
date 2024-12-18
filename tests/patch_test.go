package tests

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.verv.tech/matreshka/environment"
	"go.verv.tech/matreshka/resources"

	"go.verv.tech/matreshka-be/pkg/matreshka_be_api"
)

type PatchConfigSuite struct {
	suite.Suite

	ctx context.Context
}

func (s *PatchConfigSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *PatchConfigSuite) Test_PatchConfigEnvironment() {
	s.T().Skip()
	serviceName := s.T().Name()
	testEnv.create(s.T(), serviceName)

	newConfig := getFullConfig(s.T())

	patchReq := &matreshka_be_api.PatchConfig_Request{
		ServiceName: serviceName,
	}

	// Change old environment variable
	{
		newConfig.Environment[0].Value = []int{50051}

		patchStr := fmt.Sprint(newConfig.Environment[0].Value)
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_be_api.Node{
				Name:  "ENVIRONMENT_AVAILABLE-PORTS",
				Value: &patchStr,
			})
	}
	// Delete environment variable
	{
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_be_api.Node{
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
			&matreshka_be_api.Node{
				Name:  "ENVIRONMENT_NEW-VALUE",
				Value: &someValue,
				InnerNodes: []*matreshka_be_api.Node{
					{
						Name:  "TYPE",
						Value: &valueType,
					},
				},
			},
		)
	}

	patchResp, err := testEnv.grpcApi.PatchConfig(s.ctx, patchReq)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), patchResp)

	patchedConfig := testEnv.get(s.T(), serviceName)

	sort.Slice(newConfig.Environment, func(i, j int) bool {
		return newConfig.Environment[i].Name < newConfig.Environment[j].Name
	})

	require.Equal(s.T(), patchedConfig, newConfig)
}

func (s *PatchConfigSuite) Test_PatchConfigServers() {
	s.T().Skip()
	serviceName := s.T().Name()
	testEnv.create(s.T(), serviceName)

	newConfig := getFullConfig(s.T())

	patchReq := &matreshka_be_api.PatchConfig_Request{
		ServiceName: serviceName,
	}

	patchResp, err := testEnv.grpcApi.PatchConfig(s.ctx, patchReq)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), patchResp)

	patchedConfig := testEnv.get(s.T(), serviceName)

	sort.Slice(newConfig.Environment, func(i, j int) bool {
		return newConfig.Environment[i].Name < newConfig.Environment[j].Name
	})

	require.Equal(s.T(), patchedConfig, newConfig)
}

func (s *PatchConfigSuite) Test_PatchConfigDataSources() {
	s.T().Skip()
	serviceName := s.T().Name()
	testEnv.create(s.T(), serviceName)

	newConfig := getFullConfig(s.T())

	patchReq := &matreshka_be_api.PatchConfig_Request{
		ServiceName: serviceName,
	}

	// Change old resource (pg) port
	{
		pg := newConfig.DataSources[0].(*resources.Postgres)
		pg.Port = 5433
		portStr := strconv.Itoa(int(pg.Port))
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_be_api.Node{
				Name:  "DATA-SOURCES_POSTGRES_PORT",
				Value: &portStr,
			})
	}

	// Delete old resource (telegram) data source
	{
		newConfig.DataSources = newConfig.DataSources[:3]
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_be_api.Node{
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
			&matreshka_be_api.Node{
				Name:  "DATA-SOURCES_TELEGRAM-BOT_API-KEY",
				Value: &tg.ApiKey,
			})
	}

	patchResp, err := testEnv.grpcApi.PatchConfig(s.ctx, patchReq)
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
