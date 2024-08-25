package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

type GetVersionSuite struct {
	suite.Suite
}

func (s *GetVersionSuite) Test_GetVersion() {
	ctx := context.Background()
	resp, err := testEnv.grpcApi.ApiVersion(ctx, &matreshka_be_api.ApiVersion_Request{})

	s.Require().NoError(err)
	s.Require().NotNil(resp)
}

func Test_GetVersion(t *testing.T) {
	suite.Run(t, new(GetVersionSuite))
}
