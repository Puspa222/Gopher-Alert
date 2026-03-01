package service

import (
	"github.com/puspa222/gopher-alert/internal/notifier"
)

type NotificationService struct {
	registry *notifier.Registry
}

func NewNotificationService(r *notifier.Registry) *NotificationService {
	return &NotificationService{
		registry: r,
	}
}

func (s *NotificationService) Send(provider, to, message string) error {
	n, err := s.registry.Get(provider)
	if err != nil {
		return err
	}
	return n.Send(to, message)
}
