package v1

import (
	"go.verv.tech/matreshka-be/internal/service"
	"go.verv.tech/matreshka-be/internal/service/v1/config"
	"go.verv.tech/matreshka-be/internal/service/v1/subscription"
	"go.verv.tech/matreshka-be/internal/storage"
	"go.verv.tech/matreshka-be/internal/storage/tx_manager"
)

type Services struct {
	configService *config.CfgService
	pubSubService *subscription.PubSubService
}

func New(data storage.Data, txManager *tx_manager.TxManager) *Services {
	pubSubService := subscription.New()

	return &Services{
		configService: config.New(data, txManager, pubSubService),
		pubSubService: pubSubService,
	}
}

func (s *Services) ConfigService() service.ConfigService {
	return s.configService
}
func (s *Services) PubSubService() service.PubSubService {
	return s.pubSubService
}
