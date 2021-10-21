package notification

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/communication"
)

func (c *DummyNotificationService) List(cursor uint64, limit uint64) ([]communication.Notification, error) {
	notificationsCount := uint64(len(c.notifications))
	if cursor >= notificationsCount {
		return nil, fmt.Errorf("cursor is larger than notifications count")
	}
	resultNotifications := make([]communication.Notification, 0, limit)
	var cnt uint64 = 0
	for i := cursor; i < uint64(len(c.notifications)) && cnt < limit; i++ {
		resultNotifications = append(resultNotifications, c.notifications[i])
		cnt++
	}
	return resultNotifications, nil
}
