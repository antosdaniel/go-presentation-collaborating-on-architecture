package notification

type API interface {
	SendNotification(templateID string, userID string) error
}
