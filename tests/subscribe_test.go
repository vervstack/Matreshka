package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.redsock.ru/toolbox"

	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

type SubscriptionSuite struct {
	suite.Suite

	ctx        context.Context
	configName string
	apiClient  api.MatreshkaBeAPIClient
}

func (s *SubscriptionSuite) SetupTest() {
	s.ctx = context.Background()

	s.apiClient = testEnv.matreshkaApi

	s.configName = getServiceNameFromTest(s.T())
	testEnv.create(s.T(), s.configName)

}

func (s *SubscriptionSuite) TestSubscribeOnChanges() {
	stream, err := s.apiClient.SubscribeOnChanges(s.ctx)
	require.NoError(s.T(), err)

	// Subscribe onto changes
	{
		subscribeRequest := &api.SubscribeOnChanges_Request{
			SubscribeConfigNames: []string{s.configName},
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

	doneC := make(chan struct{})

	go func() {
		defer close(doneC)

		updates, err := stream.Recv()
		require.NoError(s.T(), err)

		updatesExpected := &api.SubscribeOnChanges_Response{
			ConfigName: s.configName,
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

		require.Equal(s.T(), updates.ConfigName, updatesExpected.ConfigName)
		require.Equal(s.T(), updates.Changes, updatesExpected.Changes)
	}()
	// Perform change in configuration
	{
		patch := &api.PatchConfig_Request{
			ConfigName: s.configName,
			// TODO implement
			//Patches:    []*api.Node{newVariable, newVariableType},
		}
		_, err = s.apiClient.PatchConfig(s.ctx, patch)
		require.NoError(s.T(), err)
	}

	select {
	case <-time.After(time.Second):
		s.T().Fatal("timed out waiting for subscription to be received")
	case <-doneC:

	}

}

func Test_Subscription(t *testing.T) {
	suite.Run(t, new(SubscriptionSuite))
}
