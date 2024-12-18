package servicev1

import (
	"go.verv.tech/matreshka-be/internal/storage"
	"go.verv.tech/matreshka-be/internal/storage/tx_manager"
)

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
