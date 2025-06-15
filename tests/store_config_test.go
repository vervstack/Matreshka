package tests

import (
	"context"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/suite"

	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

//go:embed data/loki.example.yaml
var lokiConfig []byte

type StoreConfigSuite struct {
	suite.Suite

	ctx        context.Context
	configName string
	apiClient  api.MatreshkaBeAPIClient
}

func (s *StoreConfigSuite) SetupTest() {
	s.ctx = context.Background()
	s.apiClient = testEnv.matreshkaApi

	s.configName = testEnv.create(s.T())
}

func (s *StoreConfigSuite) TestStore() {
	storeReq := &api.StoreConfig_Request{
		Format:     api.Format_yaml,
		ConfigName: s.configName,
		Config:     lokiConfig,
	}

	_, err := s.apiClient.StoreConfig(s.ctx, storeReq)
	s.Require().NoError(err)

	getReq := &api.GetConfig_Request{
		ConfigName: s.configName,
		Format:     api.Format_yaml,
	}

	cfg, err := s.apiClient.GetConfig(s.ctx, getReq)
	s.Require().NoError(err)
	s.Require().YAMLEq(string(lokiConfig), string(cfg.Config))
}

func Test_StoreConfig(t *testing.T) {
	suite.Run(t, new(StoreConfigSuite))
}
