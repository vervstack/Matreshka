package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

type ListSuite struct {
	suite.Suite

	ctx         context.Context
	serviceName string
}

func (s *ListSuite) SetupSuite() {
	s.ctx = context.Background()

	s.serviceName = s.T().Name()
	testEnv.create(s.T(), s.serviceName, fullConfigBytes)
}

func (s *ListSuite) Test_ListWithPattern() {
	listReq := &matreshka_api.ListConfigs_Request{
		SearchPattern: s.serviceName,
	}
	resp, err := testEnv.grpcApi.ListConfigs(s.ctx, listReq)
	s.NoError(err)

	expectedList := []*matreshka_api.AppInfo{{
		Name:    s.serviceName,
		Version: "v0.0.1",
	}}
	s.Equal(expectedList, resp.Services)
}

func Test_List(t *testing.T) {
	suite.Run(t, new(ListSuite))
}
