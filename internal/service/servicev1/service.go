package servicev1

import (
	errors "github.com/Red-Sock/trace-errors"
	"google.golang.org/grpc/codes"

	"github.com/godverv/matreshka-be/internal/storage"
	"github.com/godverv/matreshka-be/internal/storage/tx_manager"
)

var ErrInvalidPatchName = errors.New("invalid patched env var name", codes.InvalidArgument)

var allowedSegments = []string{
	appInfo,
	environmentSegment,
	dataSourceSegment,
	serverSegment,
}

type ConfigService struct {
	configStorage storage.Data
	txManager     *tx_manager.TxManager

	validator validator
}

func New(data storage.Data, txManager *tx_manager.TxManager) *ConfigService {
	return &ConfigService{
		configStorage: data,
		txManager:     txManager,

		validator: newValidator(),
	}
}
