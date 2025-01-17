package subscription

import (
	"go.verv.tech/matreshka-be/internal/domain"
)

type Subscriber struct {
	isStoppedChan chan struct{}
	updatesChan   chan domain.PatchConfigRequest
}

func NewSubscriber() *Subscriber {
	return &Subscriber{
		isStoppedChan: make(chan struct{}),
		updatesChan:   make(chan domain.PatchConfigRequest),
	}
}
func (s *Subscriber) Consume(request domain.PatchConfigRequest) {
	s.updatesChan <- request
}

func (s *Subscriber) GetUpdateChan() chan domain.PatchConfigRequest {
	return s.updatesChan
}

func (s *Subscriber) Stop() {
	close(s.updatesChan)
	<-s.isStoppedChan
}

func (s *Subscriber) NotifyStopped() {
	close(s.isStoppedChan)
}
