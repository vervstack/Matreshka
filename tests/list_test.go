package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"

	"go.vervstack.ru/matreshka-be/pkg/matreshka_be_api"
)

type ListSuite struct {
	suite.Suite

	ctx         context.Context
	serviceName string
}

func (s *ListSuite) SetupSuite() {
	s.ctx = context.Background()

	s.serviceName = s.T().Name()
	testEnv.create(s.T(), s.serviceName)
}

func (s *ListSuite) Test_ListWithPattern() {
	listReq := &matreshka_be_api.ListConfigs_Request{
		SearchPattern: s.serviceName,
	}
	start := time.Now().Add(-time.Minute).UTC()

	resp, err := testEnv.matreshkaApi.ListConfigs(s.ctx, listReq)
	require.NoError(s.T(), err)

	expectedList := &matreshka_be_api.ListConfigs_Response{
		Services: []*matreshka_be_api.AppInfo{{
			Name:    s.serviceName,
			Version: "v0.0.1",
		}},
		TotalRecords: 1,
	}

	tm := time.Unix(resp.Services[0].UpdatedAtUtcTimestamp, 0).UTC()
	require.True(s.T(), tm.After(start), "new service's config timestamp MUST be over %v got %v", start, tm)
	resp.Services[0].UpdatedAtUtcTimestamp = 0

	if !proto.Equal(expectedList, resp) {
		require.Equal(s.T(), expectedList, resp)
	}
}

func Test_List(t *testing.T) {
	suite.Run(t, new(ListSuite))
}
