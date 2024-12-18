//go:build untested

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.verv.tech/matreshka"
)

type CreateConfigSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *CreateConfigSuite) SetupSuite() {
	s.ctx = context.Background()

}

// Simply create and read created config
func (s *CreateConfigSuite) Test_CreateOnceAndRead() {
	serviceName := s.T().Name()
	testEnv.create(s.T(), serviceName, fullConfigBytes)
	readConfig := testEnv.get(s.T(), serviceName)

	expectedCfg := getFullConfig(s.T())

	require.Equal(s.T(), readConfig, expectedCfg)
}

// Create config, then replace it with different config
func (s *CreateConfigSuite) Test_FullUpdate() {
	serviceName := s.T().Name()
	testEnv.create(s.T(), serviceName, fullConfigBytes)

	newConfig := getFullConfig(s.T())
	newConfig.DataSources = matreshka.DataSources{}
	newConfigBytes, err := newConfig.Marshal()
	require.NoError(s.T(), err)

	testEnv.create(s.T(), serviceName, newConfigBytes)

	readConfig := testEnv.get(s.T(), serviceName)

	require.Equal(s.T(), readConfig, newConfig)
}

func Test_CreateConfigSuite(t *testing.T) {
	suite.Run(t, new(CreateConfigSuite))
}
