package service

//go:generate minimock -i Subscriber -o ./../../tests/mocks -g -s "_mock.go"

import (
	"context"

	"go.redsock.ru/evon"

	"go.vervstack.ru/matreshka-be/internal/domain"
)

type Services interface {
	ConfigService() ConfigService
	PubSubService() PubSubService
}

type ConfigService interface {
	Patch(ctx context.Context, configPatch domain.PatchConfigRequest) error
	Create(ctx context.Context, serviceName string) (int64, error)
	Rename(ctx context.Context, oldName, newName string) error

	GetNodes(ctx context.Context, serviceName string) (*evon.Node, error)
	ListConfigs(ctx context.Context, req domain.ListConfigsRequest) (domain.ListConfigsResponse, error)
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
