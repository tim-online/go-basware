package basware

type FilesService struct {
	client *Client
}

func NewFilesService(client *Client) *FilesService {
	return &FilesService{client: client}
}
