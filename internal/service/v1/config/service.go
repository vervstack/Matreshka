package config

import (
	"go.vervstack.ru/matreshka/internal/service"
	"go.vervstack.ru/matreshka/internal/storage"
	"go.vervstack.ru/matreshka/internal/storage/tx_manager"
	"go.vervstack.ru/matreshka/internal/validation"
)

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
