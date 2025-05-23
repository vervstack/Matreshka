package config

import (
	"go.vervstack.ru/matreshka/internal/service"
	"go.vervstack.ru/matreshka/internal/storage"
	"go.vervstack.ru/matreshka/internal/storage/tx_manager"
	"go.vervstack.ru/matreshka/internal/utils/validation"
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

		validator:  validation.New(),
		pubService: pubService,
	}
}
