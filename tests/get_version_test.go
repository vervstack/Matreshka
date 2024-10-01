//go:build integration

package tests

import (
	"context"
	"io"
	"net/http"
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

func (s *GetVersionSuite) Test_Gateway() {
	resp, err := http.Get("http://localhost:8080/api/version")
	s.Require().NoError(err)

	bd, err := io.ReadAll(resp.Body)
	s.Require().NoError(err)

	s.Require().Equal(`{"version":"v1.0.32"}`, string(bd))
}

func Test_GetVersion(t *testing.T) {
	suite.Run(t, new(GetVersionSuite))
}
