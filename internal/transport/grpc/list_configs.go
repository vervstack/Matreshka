package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/internal/domain"
	"github.com/godverv/matreshka-be/internal/transport/grpc/converters"
	api "github.com/godverv/matreshka-be/pkg/matreshka_api"
)

func (a *App) ListConfigs(ctx context.Context, req *api.ListConfigs_Request,
) (*api.ListConfigs_Response, error) {

	listReq := domain.ListConfigsRequest{
		ListRequest: domain.ListRequest{
			Limit:  req.GetListRequest().GetLimit(),
			Offset: req.GetListRequest().GetOffset(),
		},
		SearchPattern: req.GetServiceName(),
	}

	infos, err := a.storage.ListConfigs(ctx, listReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &api.ListConfigs_Response{
		Services: make([]*api.Config_AppConfig, 0, len(infos)),
	}

	for _, c := range infos {
		resp.Services = append(resp.Services, converters.ToProtoAppInfo(c))
	}

	return resp, nil
}
