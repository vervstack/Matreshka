package tests

import (
	"context"
	"testing"

	"github.com/godverv/matreshka"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	createOnce  = "create-once-config"
	createTwice = "create-twice-config"
)

type CreateConfigSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *CreateConfigSuite) SetupSuite() {
	s.ctx = context.Background()

}

// Simply create and read created config
func (s *CreateConfigSuite) Test_CreateAndRead() {
	testEnv.create(s.T(), createOnce, fullConfigBytes)
	readConfig := testEnv.get(s.T(), createOnce)

	expectedCfg := getFullConfig(s.T())

	require.Equal(s.T(), readConfig, expectedCfg)
}

// Create config, then replace it with different config
func (s *CreateConfigSuite) Test_FullUpdate() {
	testEnv.create(s.T(), createTwice, fullConfigBytes)

	newConfig := getFullConfig(s.T())
	newConfig.DataSources = matreshka.DataSources{}
	newConfigBytes, err := newConfig.Marshal()
	require.NoError(s.T(), err)

	testEnv.create(s.T(), createTwice, newConfigBytes)

	readConfig := testEnv.get(s.T(), createTwice)

	require.Equal(s.T(), readConfig, newConfig)
}

func Test_CreateConfigSuite(t *testing.T) {
	suite.Run(t, new(CreateConfigSuite))
}
