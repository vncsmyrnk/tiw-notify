// vim: noexpandtab

//go:generate mockgen -source=notification.go -destination=mocks/mock_notification.go -package=mocks

package notification

import (
	"github.com/gen2brain/beeep"
)

type Notifier interface {
	Notify(string, string) error
}

type BeeepNotifier struct{}

func (bn BeeepNotifier) Notify(title, message string) error {
	iconPath := ""

	err := beeep.Notify(title, message, iconPath)
	if err != nil {
		return err
	}
	return nil
}
