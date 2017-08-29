package basware

type NotificationsService struct {
	client *Client
}

func NewNotificationsService(client *Client) *NotificationsService {
	return &NotificationsService{client: client}
}
