package api

import "log/slog"

type APIWithLogs struct {
	API API
}

func (a APIWithLogs) SendNotification(templateID string, userID string) error {
	l := slog.With(
		slog.String("template_id", templateID),
		slog.String("user_id", userID),
	)
	l.Info("Sending notification")

	err := a.API.SendNotification(templateID, userID)

	if err != nil {
		l.Error("Failed to send notification", "error", err)
	} else {
		l.Info("Notification sent")
	}
	return err
}
