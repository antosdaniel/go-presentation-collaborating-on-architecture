package infra

type SmsSender struct{}

func (SmsSender) Send(phone string, message string) {}
