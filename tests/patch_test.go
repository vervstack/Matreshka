package tests

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/godverv/matreshka/environment"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/godverv/matreshka-be/pkg/matreshka_api"
)

const (
	patchConfig = "patch-config"
)

type PatchConfigSuite struct {
	suite.Suite

	ctx context.Context
}

func (s *PatchConfigSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *PatchConfigSuite) Test_PatchExistingConfig() {
	testEnv.create(s.T(), patchConfig, fullConfigBytes)

	newConfig := getFullConfig(s.T())

	patchReq := &matreshka_api.PatchConfig_Request{
		ServiceName: patchConfig,
	}

	// Change old environment value
	{
		newConfig.Environment[0].Value = []int{50051}

		portStr := fmt.Sprint(newConfig.Environment[0].Value)
		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name:  "ENVIRONMENT_AVAILABLE-PORTS",
				Value: &portStr,
			})
	}
	// Add new environment value
	{
		someValue := "rand val"
		valueType := string(environment.VariableTypeStr)

		newEnvVar := &environment.Variable{
			Name:  "new value",
			Value: someValue,
			Type:  environment.VariableTypeStr,
		}

		newConfig.Environment = append(newConfig.Environment, newEnvVar)

		patchReq.Changes = append(patchReq.Changes,
			&matreshka_api.Node{
				Name:  "ENVIRONMENT_NEW-VALUE",
				Value: &someValue,
				InnerNodes: []*matreshka_api.Node{
					{
						Name:  "TYPE",
						Value: &valueType,
					},
				},
			},
		)
	}
	patchResp, err := testEnv.grpcClient.PatchConfig(s.ctx, patchReq)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), patchResp)

	patchedConfig := testEnv.get(s.T(), patchConfig)

	sort.Slice(newConfig.Environment, func(i, j int) bool {
		return newConfig.Environment[i].Name < newConfig.Environment[j].Name
	})

	require.Equal(s.T(), patchedConfig, newConfig)
}

func Test_PatchConfig(t *testing.T) {
	suite.Run(t, new(PatchConfigSuite))
}
