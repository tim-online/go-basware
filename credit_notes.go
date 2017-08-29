package basware

type CreditNotesService struct {
	client *Client
}

func NewCreditNotesService(client *Client) *CreditNotesService {
	return &CreditNotesService{client: client}
}
