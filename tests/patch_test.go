package tests

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"go.vervstack.ru/matreshka/pkg/matreshka"
	"go.vervstack.ru/matreshka/pkg/matreshka/environment"
	"go.vervstack.ru/matreshka/pkg/matreshka/resources"
	"go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

type PatchConfigSuite struct {
	suite.Suite

	ctx         context.Context
	serviceName string
	cfg         matreshka.AppConfig
	patchReq    *matreshka_be_api.PatchConfig_Request
}

func (s *PatchConfigSuite) SetupTest() {
	s.ctx = context.Background()

	s.serviceName = getServiceNameFromTest(s.T())
	testEnv.create(s.T(), s.serviceName)

	s.cfg = getFullConfig(s.T())
	s.cfg.Name = s.serviceName
	testEnv.patchConfig(s.T(), s.cfg)

	s.patchReq = &matreshka_be_api.PatchConfig_Request{
		ServiceName: s.serviceName,
	}
}

func (s *PatchConfigSuite) Test_PatchConfigEnvironment() {
	// Change old environment variable
	{
		s.cfg.Environment[5] = environment.MustNewVariable(
			s.cfg.Environment[5].Name, []int{50051})

		patchStr := fmt.Sprint(s.cfg.Environment[5].Value)
		s.patchReq.Changes = append(s.patchReq.Changes,
			&matreshka_be_api.Node{
				Name:  "ENVIRONMENT_AVAILABLE-PORTS",
				Value: &patchStr,
			})
	}
	// Delete environment variable
	{
		s.patchReq.Changes = append(s.patchReq.Changes,
			&matreshka_be_api.Node{
				Name: "ENVIRONMENT_CREDIT-PERCENTS-BASED-ON-YEAR-OF-BIRTH",
			})
		s.cfg.Environment = s.cfg.Environment[:len(s.cfg.Environment)-1]
	}
	// Add new environment variable
	{
		someValue := "rand val"
		valueType := string(environment.VariableTypeStr)

		newEnvVar := environment.MustNewVariable("new_value", someValue)

		s.cfg.Environment = append(s.cfg.Environment, newEnvVar)

		s.patchReq.Changes = append(s.patchReq.Changes,
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
}

func (s *PatchConfigSuite) Test_PatchConfigServers() {
	const port = 50051
	s.cfg.Servers[port].GRPC["/{GRPC}"].Gateway = "/api/v2"
	s.patchReq.Changes = append(s.patchReq.Changes,
		&matreshka_be_api.Node{
			Name:  "SERVERS_MASTER2_/{GRPC}_GATEWAY",
			Value: &s.cfg.Servers[port].GRPC["/{GRPC}"].Gateway,
		})
}

func (s *PatchConfigSuite) Test_PatchConfigDataSources() {

	// Change old resource (pg) port
	{
		pg := s.cfg.DataSources[0].(*resources.Postgres)
		pg.Port = 5433
		portStr := strconv.Itoa(int(pg.Port))
		s.patchReq.Changes = append(s.patchReq.Changes,
			&matreshka_be_api.Node{
				Name:  "DATA-SOURCES_POSTGRES_PORT",
				Value: &portStr,
			})
	}

	// Delete old resource (telegram) data source
	// Add new telegram data source
	{
		tg := s.cfg.DataSources[2].(*resources.Telegram)
		tg.Name = "telegram_bot"
		tg.ApiKey = "jjggwwkk"
		s.patchReq.Changes = append(s.patchReq.Changes,
			&matreshka_be_api.Node{
				Name: "DATA-SOURCES_TELEGRAM",
			})
		s.patchReq.Changes = append(s.patchReq.Changes,
			&matreshka_be_api.Node{
				Name:  "DATA-SOURCES_TELEGRAM-BOT_API-KEY",
				Value: &tg.ApiKey,
			})
	}
}

func (s *PatchConfigSuite) TearDownTest() {
	patchResp, err := testEnv.matreshkaApi.PatchConfig(s.ctx, s.patchReq)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), patchResp)

	patchedConfig := testEnv.get(s.T(), s.serviceName)

	sort.Slice(s.cfg.Environment, func(i, j int) bool {
		return s.cfg.Environment[i].Name < s.cfg.Environment[j].Name
	})

	require.Equal(s.T(), patchedConfig.Environment, s.cfg.Environment)
}

func Test_PatchConfig(t *testing.T) {
	suite.Run(t, new(PatchConfigSuite))
}
