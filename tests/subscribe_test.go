package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.redsock.ru/toolbox"

	api "go.verv.tech/matreshka-be/pkg/matreshka_be_api"
)

type SubscriptionSuite struct {
	suite.Suite

	ctx         context.Context
	serviceName string
	apiClient   api.MatreshkaBeAPIClient
}

func (s *SubscriptionSuite) SetupTest() {
	s.ctx = context.Background()
	// TODO
	s.apiClient = testEnv.grpcApi

	s.serviceName = getServiceNameFromTest(s.T())
	testEnv.create(s.T(), s.serviceName)

}

func (s *SubscriptionSuite) TestSubscribeOnChanges() {
	stream, err := s.apiClient.SubscribeOnChanges(s.ctx)
	require.NoError(s.T(), err)
	// Subscribe onto changes
	{
		subscribeRequest := &api.SubscribeOnChanges_Request{
			SubscribeServiceNames: []string{s.serviceName},
		}
		err = stream.Send(subscribeRequest)
		require.NoError(s.T(), err)
	}

	newVariableType := &api.Node{
		Name:  "ENVIRONMENT_SOME-VARIABLE_TYPE",
		Value: toolbox.ToPtr("string"),
	}

	newVariable := &api.Node{
		Name:  "ENVIRONMENT_SOME-VARIABLE",
		Value: toolbox.ToPtr("123"),
	}

	// Perform change in configuration
	{
		patch := &api.PatchConfig_Request{
			ServiceName: s.serviceName,
			Changes:     []*api.Node{newVariable, newVariableType},
		}
		_, err = s.apiClient.PatchConfig(s.ctx, patch)
		require.NoError(s.T(), err)
	}

	updatesExpected := &api.SubscribeOnChanges_Response{
		ServiceName: s.serviceName,
		Changes: &api.SubscribeOnChanges_Response_EnvVariables{
			EnvVariables: &api.SubscribeOnChanges_EnvChanges{
				EnvVariables: []*api.Node{
					{
						Name:  newVariable.Name,
						Value: newVariable.Value,
					},
					{
						Name:  newVariableType.Name,
						Value: newVariableType.Value,
					},
				},
			},
		},
	}

	updates, err := stream.Recv()
	require.NoError(s.T(), err)

	require.Equal(s.T(), updates.ServiceName, updatesExpected.ServiceName)
	require.Equal(s.T(), updates.Changes, updatesExpected.Changes)

}

func (s *SubscriptionSuite) TearDownSuite() {

}

func Test_Subscription(t *testing.T) {
	suite.Run(t, new(SubscriptionSuite))
}
