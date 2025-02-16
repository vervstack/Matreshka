package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.redsock.ru/toolbox"
	"google.golang.org/protobuf/proto"

	"go.vervstack.ru/matreshka-be/internal/domain"
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

func (s *ListSuite) SetupTest() {
	s.ctx = context.Background()

	s.start = time.Now().Add(-time.Minute).UTC()
	s.serviceName = getServiceNameFromTest(s.T())
}

func (s *ListSuite) Test_ListOneServiceWithOneVersion() {
	testEnv.create(s.T(), s.serviceName)

	s.req = &matreshka_be_api.ListConfigs_Request{
		SearchPattern: s.serviceName,
	}

	s.expected = &matreshka_be_api.ListConfigs_Response{
		Services: []*matreshka_be_api.AppInfo{{
			Name:           s.serviceName,
			ServiceVersion: "v0.0.1",
			ConfigVersions: []string{domain.MasterVersion},
		}},
		TotalRecords: 1,
	}
}

func (s *ListSuite) Test_ListOneServiceWithTwoVersion() {
	testEnv.create(s.T(), s.serviceName)

	patchReq := &matreshka_be_api.PatchConfig_Request{
		ServiceName: s.serviceName,
		Changes: []*matreshka_be_api.Node{
			{
				Name:  "ENVIRONMENT_IS_CRON_ACTIVE",
				Value: toolbox.ToPtr("true"),
				InnerNodes: []*matreshka_be_api.Node{
					{
						Name:  "TYPE",
						Value: toolbox.ToPtr("bool"),
					},
				},
			},
		},
		Version: toolbox.ToPtr("VERV-137"),
	}

	_, err := testEnv.matreshkaApi.PatchConfig(s.ctx, patchReq)
	require.NoError(s.T(), err)

	s.req = &matreshka_be_api.ListConfigs_Request{
		SearchPattern: s.serviceName,
	}

	s.expected = &matreshka_be_api.ListConfigs_Response{
		Services: []*matreshka_be_api.AppInfo{{
			Name:           s.serviceName,
			ServiceVersion: "v0.0.1",
			ConfigVersions: []string{domain.MasterVersion, "VERV-137"},
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
