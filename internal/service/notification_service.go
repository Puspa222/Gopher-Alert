package service

import (
	"github.com/puspa222/gopher-alert/internal/notifier"
	"github.com/puspa222/gopher-alert/internal/storage"
)

type NotificationService struct {
	registry *notifier.Registry
	db       *storage.DB
}

func NewNotificationService(r *notifier.Registry, db *storage.DB) *NotificationService {
	return &NotificationService{
		registry: r,
		db:       db,
	}
}

func (s *NotificationService) Send(providerName, to, message string) error {
	p, err := s.registry.Get(providerName)
	status := "success"
	errMsg := ""
	if err != nil {
		status = "failed"
		errMsg = err.Error()
		s.db.LogNotification(providerName, to, message, status, errMsg)
		return err
	}

	err = p.Send(to, message)
	if err != nil {
		status = "failed"
		errMsg = err.Error()
	}

	// Log to SQLite
	if dbErr := s.db.LogNotification(providerName, to, message, status, errMsg); dbErr != nil {
		// optional: log db error to console
	}

	return err
}
