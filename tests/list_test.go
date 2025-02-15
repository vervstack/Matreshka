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
	start       time.Time

	req      *matreshka_be_api.ListConfigs_Request
	expected *matreshka_be_api.ListConfigs_Response
}

func (s *ListSuite) SetupSuite() {
	s.ctx = context.Background()

	s.start = time.Now().Add(-time.Minute).UTC()
	s.serviceName = s.T().Name()
	testEnv.create(s.T(), s.serviceName)

}

func (s *ListSuite) Test_ListWithPattern() {
	s.req = &matreshka_be_api.ListConfigs_Request{
		SearchPattern: s.serviceName,
	}

	s.expected = &matreshka_be_api.ListConfigs_Response{
		Services: []*matreshka_be_api.AppInfo{{
			Name:    s.serviceName,
			Version: "v0.0.1",
		}},
		TotalRecords: 1,
	}
}

func (s *ListSuite) TearDownTest() {
	resp, err := testEnv.matreshkaApi.ListConfigs(s.ctx, s.req)
	require.NoError(s.T(), err)

	tm := time.Unix(resp.Services[0].UpdatedAtUtcTimestamp, 0).UTC()
	require.WithinRange(s.T(), tm, s.start, time.Now().UTC())
	resp.Services[0].UpdatedAtUtcTimestamp = 0

	if !proto.Equal(s.expected, resp) {
		require.Equal(s.T(), s.expected, resp)
	}
}
func Test_List(t *testing.T) {
	suite.Run(t, new(ListSuite))
}
