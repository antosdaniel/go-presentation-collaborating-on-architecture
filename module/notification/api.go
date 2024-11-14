package notification

type API interface {
	SendNotification(templateID string, userID string) error
}

func New(templateRepo TemplateRepo, userRepo UserRepo, smsSender SmsSender) API {
	return withLogs{
		API: api{
			TemplateRepo: templateRepo,
			UserRepo:     userRepo,
			SmsSender:    smsSender,
		},
	}
}

type api struct {
	TemplateRepo TemplateRepo
	UserRepo     UserRepo
	SmsSender    SmsSender
}

func (a api) SendNotification(templateID string, userID string) error {
	template := a.TemplateRepo.Find(templateID)
	user := a.UserRepo.Find(userID)

	message := template.FillWithName(user.FirstName)

	a.SmsSender.Send(user.Phone, message)

	return nil
}

type UserRepo interface {
	Find(id string) User
}

type TemplateRepo interface {
	Find(id string) Template
}

type SmsSender interface {
	Send(phone string, message string)
}
