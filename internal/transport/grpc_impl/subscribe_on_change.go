package grpc_impl

import (
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"go.redsock.ru/rerrors"

	"go.vervstack.ru/matreshka/internal/service/v1/subscription"
	api "go.vervstack.ru/matreshka/pkg/matreshka_be_api"
)

func (a *Impl) SubscribeOnChanges(stream api.MatreshkaBeAPI_SubscribeOnChangesServer) error {
	sub := subscription.NewSubscriber()

	defer a.subService.StopSubscription(sub)
	defer sub.NotifyStopped()

	subscriberEvents := consumeSubscriberStream(stream)

	errorCount := 0

	for {
		select {
		case req, ok := <-subscriberEvents:
			if !ok {
				return nil
			}

			a.subService.Subscribe(sub, req.SubscribeServiceNames...)
			a.subService.Unsubscribe(sub, req.UnsubscribeServiceNames...)

		case updates := <-sub.GetUpdateChan():

			patch := &api.SubscribeOnChanges_Response{
				ServiceName: updates.ServiceName,
				Timestamp:   uint32(time.Now().UTC().Unix()),
			}

			envChanges := make([]*api.Node, 0)

			for _, b := range updates.Batch {
				envChanges = append(envChanges, &api.Node{
					Name:  b.FieldName,
					Value: b.FieldValue,
				})
			}

			patch.Changes = &api.SubscribeOnChanges_Response_EnvVariables{
				EnvVariables: &api.SubscribeOnChanges_EnvChanges{
					EnvVariables: envChanges,
				},
			}

			err := stream.Send(patch)
			if err != nil {
				logrus.Errorf("error sending update to subscriber %s", err)
				errorCount++
				if errorCount >= 3 {
					return rerrors.Wrap(err)
				}
			}
		}
	}
}

func consumeSubscriberStream(stream api.MatreshkaBeAPI_SubscribeOnChangesServer) <-chan *api.SubscribeOnChanges_Request {
	subscriberEvents := make(chan *api.SubscribeOnChanges_Request, 1)

	errorCount := 0

	go func() {
		defer close(subscriberEvents)

		for {
			req, err := stream.Recv()
			if err != nil {
				if !rerrors.Is(err, io.EOF) {
					logrus.Error(rerrors.Wrap(err, "error receiving names from subscription"))
					errorCount++
					if errorCount > 3 {
						return
					}

					continue
				}

				return
			}
			subscriberEvents <- req
		}
	}()

	return subscriberEvents

}
