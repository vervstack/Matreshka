package grpc_impl

import (
	"context"

	"github.com/Red-Sock/toolbox"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/godverv/matreshka-be/internal/domain"
	api "github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

func (a *Impl) ListConfigs(ctx context.Context, req *api.ListConfigs_Request) (*api.ListConfigs_Response, error) {
	listReq := domain.ListConfigsRequest{
		Paging: domain.Paging{
			Limit:  toolbox.Coalesce(req.GetPaging().GetLimit(), 10),
			Offset: req.GetPaging().GetOffset(),
		},
		Sort: domain.Sort{
			SortType: req.Sort.GetType(),
			Desc:     req.Sort.GetDesc(),
		},

		SearchPattern: req.GetSearchPattern(),
	}

	configs, err := a.storage.ListConfigs(ctx, listReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &api.ListConfigs_Response{
		Services:     make([]*api.AppInfo, 0, len(configs.List)),
		TotalRecords: configs.TotalRecords,
	}

	for _, item := range configs.List {
		resp.Services = append(resp.Services,
			&api.AppInfo{
				Name:    item.Name,
				Version: item.Version,
			})
	}

	return resp, nil
}
