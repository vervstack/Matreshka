package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.vervstack.ru/matreshka/internal/api/server/matreshka_api"
)

const missingConfigName = "missing_service"

type NotFoundSuite struct {
	suite.Suite

	ctx context.Context
	app AppEnv
}

func (s *NotFoundSuite) SetupTest() {
	s.ctx = context.Background()
	s.app = InitAppEnvironment(s.T())
}

func (s *NotFoundSuite) Test_GetConfigNodes_MissingConfig() {
	req := newGetConfigNodesRequest(missingConfigName)

	_, err := s.app.matreshkaApi.GetConfigNodes(s.ctx, req)
	require.Equal(s.T(), codes.NotFound, status.Code(err))
}

func (s *NotFoundSuite) Test_GetConfig_MissingConfig() {
	req := newGetConfigRequest(missingConfigName)

	_, err := s.app.matreshkaApi.GetConfig(s.ctx, req)
	require.Equal(s.T(), codes.NotFound, status.Code(err))
}

func newGetConfigNodesRequest(configName string) *matreshka_api.GetConfigNode_Request {
	return &matreshka_api.GetConfigNode_Request{
		ConfigName: configName,
	}
}

func newGetConfigRequest(configName string) *matreshka_api.GetConfig_Request {
	return &matreshka_api.GetConfig_Request{
		ConfigName: configName,
	}
}

func Test_NotFound(t *testing.T) {
	suite.Run(t, new(NotFoundSuite))
}
