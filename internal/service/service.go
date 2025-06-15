package service

//go:generate minimock -i Subscriber -o ./../../tests/mocks -g -s "_mock.go"

import (
	"context"

	"go.vervstack.ru/matreshka/internal/domain"
)

type Services interface {
	ConfigService() EvonConfigService
	PubSubService() PubSubService
}

type EvonConfigService interface {
	Replace(ctx context.Context, req domain.ReplaceConfigReq) error
	Patch(ctx context.Context, configPatch domain.PatchConfigRequest) error
	Create(ctx context.Context, name domain.ConfigName) (domain.AboutConfig, error)
	Rename(ctx context.Context, oldName, newName domain.ConfigName) error

	GetConfigWithNodes(ctx context.Context, name domain.ConfigName, version string) (domain.ConfigWithNodes, error)
	ListConfigs(ctx context.Context, req domain.ListConfigsRequest) (domain.ListConfigsResponse, error)
}

type PlainConfigService interface {
	Save(cfg string, version string) error
	Get(cfg string, version string) error
}

type PubSubService interface {
	PublisherService
	SubscriberService
}

type PublisherService interface {
	Publish(event domain.PatchConfigRequest)
}

type SubscriberService interface {
	Subscribe(c Subscriber, serviceNames ...string)
	Unsubscribe(c Subscriber, serviceNames ...string)
	StopSubscription(c Subscriber)
}

type Subscriber interface {
	Consume(request domain.PatchConfigRequest)
	GetUpdateChan() chan domain.PatchConfigRequest
	Stop()
}
