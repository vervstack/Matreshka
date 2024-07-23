package tests

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

type GetTestSuite struct {
	suite.Suite

	ctx context.Context
}

func (s *GetTestSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *GetTestSuite) Test_GetConfig_NotFound() {
	serviceName := s.T().Name()

	getReq := &matreshka_api.GetConfig_Request{
		ServiceName: serviceName,
	}
	resp, err := testEnv.grpcClient.GetConfig(s.ctx, getReq)

	expectedErr := status.New(codes.NotFound, "no nodes found").Err()
	s.Equal(expectedErr, err)
	s.Nil(resp)
}

func (s *GetTestSuite) Test_GetNodes_NotFound() {
	serviceName := s.T().Name()

	getReq := &matreshka_api.GetConfigNode_Request{
		ServiceName: serviceName,
	}
	resp, err := testEnv.grpcClient.GetConfigNodes(s.ctx, getReq)

	expectedErr := status.New(codes.NotFound, "no nodes found").Err()
	s.Equal(expectedErr, err)
	s.Nil(resp)
}

func (s *GetTestSuite) Test_GetNodes() {
	serviceName := s.T().Name()

	testEnv.create(s.T(), serviceName, fullConfigBytes)

	expectedConfig := []*matreshka_api.Node{
		{
			Name:  "APP-INFO_NAME",
			Value: toPtr("matreshka"),
		},
		{
			Name:  "APP-INFO_STARTUP-DURATION",
			Value: toPtr("10s")},
		{
			Name:  "APP-INFO_VERSION",
			Value: toPtr("v0.0.1")},
		{
			Name:  "DATA-SOURCES_GRPC-RSCLI-EXAMPLE_CONNECTION-STRING",
			Value: toPtr("0.0.0.0:50051")},
		{
			Name:  "DATA-SOURCES_GRPC-RSCLI-EXAMPLE_MODULE",
			Value: toPtr("github.com/Red-Sock/rscli_example")},
		{
			Name:  "DATA-SOURCES_POSTGRES_DB-NAME",
			Value: toPtr("matreshka")},
		{
			Name:  "DATA-SOURCES_POSTGRES_HOST",
			Value: toPtr("localhost")},
		{
			Name:  "DATA-SOURCES_POSTGRES_PORT",
			Value: toPtr("5432")},
		{
			Name:  "DATA-SOURCES_POSTGRES_PWD",
			Value: toPtr("matreshka")},
		{
			Name:  "DATA-SOURCES_POSTGRES_SSL-MODE",
			Value: toPtr("disable")},
		{
			Name:  "DATA-SOURCES_POSTGRES_USER",
			Value: toPtr("matreshka")},
		{
			Name:  "DATA-SOURCES_REDIS_DB",
			Value: toPtr("2")},
		{
			Name:  "DATA-SOURCES_REDIS_HOST",
			Value: toPtr("localhost")},
		{
			Name:  "DATA-SOURCES_REDIS_PORT",
			Value: toPtr("6379")},
		{
			Name:  "DATA-SOURCES_REDIS_PWD",
			Value: toPtr("redis_matreshka_pwd")},
		{
			Name:  "DATA-SOURCES_REDIS_USER",
			Value: toPtr("redis_matreshka")},
		{
			Name:  "DATA-SOURCES_TELEGRAM_API-KEY",
			Value: toPtr("some_api_key")},
		{
			Name:  "ENVIRONMENT_AVAILABLE-PORTS",
			Value: toPtr("[10,12,34,35,36,37,38,39,40]")},
		{
			Name:  "ENVIRONMENT_AVAILABLE-PORTS_TYPE",
			Value: toPtr("int")},
		{
			Name:  "ENVIRONMENT_CREDIT-PERCENT",
			Value: toPtr("0.01")},
		{
			Name:  "ENVIRONMENT_CREDIT-PERCENTS-BASED-ON-YEAR-OF-BIRTH",
			Value: toPtr("[0.01,0.02,0.03,0.04]")},
		{
			Name:  "ENVIRONMENT_CREDIT-PERCENTS-BASED-ON-YEAR-OF-BIRTH_TYPE",
			Value: toPtr("float")},
		{
			Name:  "ENVIRONMENT_CREDIT-PERCENT_TYPE",
			Value: toPtr("float")},
		{
			Name:  "ENVIRONMENT_DATABASE-MAX-CONNECTIONS",
			Value: toPtr("1")},
		{
			Name:  "ENVIRONMENT_DATABASE-MAX-CONNECTIONS_TYPE",
			Value: toPtr("int")},
		{
			Name:  "ENVIRONMENT_ONE-OF-WELCOME-STRING",
			Value: toPtr("one")},
		{
			Name:  "ENVIRONMENT_ONE-OF-WELCOME-STRING_ENUM",
			Value: toPtr("[one,two,three]")},
		{
			Name:  "ENVIRONMENT_ONE-OF-WELCOME-STRING_TYPE",
			Value: toPtr("string")},
		{
			Name:  "ENVIRONMENT_REQUEST-TIMEOUT",
			Value: toPtr("10s")},
		{
			Name:  "ENVIRONMENT_REQUEST-TIMEOUT_TYPE",
			Value: toPtr("duration")},
		{
			Name:  "ENVIRONMENT_TRUE-FALSER",
			Value: toPtr("true")},
		{
			Name:  "ENVIRONMENT_TRUE-FALSER_TYPE",
			Value: toPtr("bool")},
		{
			Name:  "ENVIRONMENT_USERNAMES-TO-BAN",
			Value: toPtr("[hacker228,mothe4acker]")},
		{
			Name:  "ENVIRONMENT_USERNAMES-TO-BAN_TYPE",
			Value: toPtr("string")},
		{
			Name:  "ENVIRONMENT_WELCOME-STRING",
			Value: toPtr("not so basic 🤡 string")},
		{
			Name:  "ENVIRONMENT_WELCOME-STRING_TYPE",
			Value: toPtr("string")},
		{
			Name:  "SERVERS_GRPC_PORT",
			Value: toPtr("50051")},
		{
			Name:  "SERVERS_REST_PORT",
			Value: toPtr("8080")},
	}

	getReq := &matreshka_api.GetConfigNode_Request{
		ServiceName: serviceName,
	}
	resp, err := testEnv.grpcClient.GetConfigNodes(s.ctx, getReq)
	s.NoError(err)

	sort.Slice(expectedConfig, func(i, j int) bool {
		return expectedConfig[i].Name > expectedConfig[j].Name
	})
	sort.Slice(resp.Root.InnerNodes, func(i, j int) bool {
		return resp.Root.InnerNodes[i].Name > resp.Root.InnerNodes[j].Name
	})

	s.Equal(expectedConfig, resp.Root.InnerNodes)
}

func Test_Get(t *testing.T) {
	suite.Run(t, new(GetTestSuite))
}

func toPtr[T any](a T) *T {
	return &a
}
