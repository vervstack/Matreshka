package config

import (
	"go.verv.tech/matreshka-be/internal/service"
	"go.verv.tech/matreshka-be/internal/storage"
	"go.verv.tech/matreshka-be/internal/storage/tx_manager"
	"go.verv.tech/matreshka-be/internal/validation"
)

var allowedSegments = []string{
	appInfo,
	environmentSegment,
	dataSourceSegment,
	serverSegment,
}

type CfgService struct {
	configStorage storage.Data
	txManager     *tx_manager.TxManager

	validator  validation.Validator
	pubService service.PublisherService
}

func New(data storage.Data, txManager *tx_manager.TxManager, pubService service.PublisherService) *CfgService {
	return &CfgService{
		configStorage: data,
		txManager:     txManager,

		validator:  validation.NewValidator(),
		pubService: pubService,
	}
}
