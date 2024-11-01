package storage

import (
	"context"
	"database/sql"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"
	"google.golang.org/grpc/codes"

	"github.com/godverv/matreshka-be/internal/domain"
)

var ErrNoNodes = errors.New("no nodes found", codes.NotFound)

type Data interface {
	// GetConfigNodes returns root node of parsed config
	// Might return ErrNoNodes
	GetConfigNodes(ctx context.Context, serviceName string) (*evon.Node, error)
	ListConfigs(ctx context.Context, req domain.ListConfigsRequest) ([]domain.ConfigListItem, error)

	SaveConfig(ctx context.Context, serviceConfig string, config matreshka.AppConfig) error

	UpsertValues(ctx context.Context, req domain.PatchConfigRequest) error
	DeleteValues(ctx context.Context, req domain.PatchConfigRequest) error

	WithTx(tx *sql.Tx) Data
}
