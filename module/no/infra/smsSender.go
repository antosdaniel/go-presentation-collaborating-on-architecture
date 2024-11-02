package infra

import "log/slog"

type SmsSender struct{}

func (SmsSender) Send(phone string, message string) {
	slog.Info("SMS notification sent", "phone", phone)
}
