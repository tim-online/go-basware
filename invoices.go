package basware

import (
	"context"
	"net/http"
	"strings"
)

var (
	endpointInvoices = "v1/invoices/{bumId}"
)

type InvoicesService struct {
	client *Client
}

func NewInvoicesService(client *Client) *InvoicesService {
	return &InvoicesService{client: client}
}

func (s *InvoicesService) Get(ctx context.Context, pathParams *InvoiceGetPathParams) (*InvoicesGetResponse, error) {
	// @TODO: create wrapper?
	responseBody := s.NewGetResponse()
	method := http.MethodGet

	path := endpointInvoices
	path = strings.Replace(path, "{bumId}", pathParams.BumID, 1)
	apiURL, err := s.client.GetEndpointURL(path)
	if err != nil {
		return nil, err
	}

	// create new request
	httpReq, err := s.client.NewRequest(ctx, method, apiURL, nil)
	if err != nil {
		return nil, err
	}

	// process query parameters
	// utils.AddQueryParamsToRequest(queryParams, httpReq, false)

	// submit the request
	_, err = s.client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *InvoicesService) NewGetResponse() *InvoicesGetResponse {
	return &InvoicesGetResponse{}
}

func (s *InvoicesService) NewGetPathParams() *InvoiceGetPathParams {
	return &InvoiceGetPathParams{}
}

type InvoiceGetPathParams struct {
	BumID string `json:"bumId"`
}

func (s *InvoicesService) Post(ctx context.Context, pathParams *InvoicePostPathParams, requestBody *InvoicesPostRequestBody) (*InvoicesPostResponseBody, error) {
	// @TODO: create wrapper?
	method := http.MethodPost
	responseBody := s.NewPostResponseBody()

	path := endpointInvoices
	path = strings.Replace(path, "{bumId}", pathParams.BumID, 1)
	apiURL, err := s.client.GetEndpointURL(path)
	if err != nil {
		return nil, err
	}

	// create new request
	httpReq, err := s.client.NewRequest(ctx, method, apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	// process query parameters
	// utils.AddQueryParamsToRequest(queryParams, httpReq, false)

	// submit the request
	_, err = s.client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *InvoicesService) NewPostPathParams() *InvoicePostPathParams {
	return &InvoicePostPathParams{}
}

type InvoicePostPathParams struct {
	BumID string `json:"bumId"`
}

func (s *InvoicesService) NewPostRequestBody() *InvoicesPostRequestBody {
	return &InvoicesPostRequestBody{}
}

func (s *InvoicesService) NewPostResponseBody() *InvoicesPostResponseBody {
	return &InvoicesPostResponseBody{}
}

type InvoicesPostResponseBody struct {
}
