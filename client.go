package basware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	uuid "github.com/satori/go.uuid"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-basware/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "api.basware.com",
		Path:   "",
	}

	BaseURLTest = url.URL{
		Scheme: "https",
		Host:   "test-api.basware.com",
		Path:   "",
	}
)

// NewClient returns a new MEWS API client
func NewClient(httpClient *http.Client, username string, password string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		http: httpClient,
	}

	c.SetUsername(username)
	c.SetPassword(password)
	c.SetBaseURL(BaseURL)
	c.SetProductionMode()
	c.SetUserAgent(userAgent)
	c.SetUserAgent(userAgent)
	c.SetMediaType(mediaType)
	c.SetCharset(charset)

	// Services
	c.Notifications = NewNotificationsService(c)
	c.Invoices = NewInvoicesService(c)
	c.CreditNotes = NewCreditNotesService(c)
	c.Files = NewFilesService(c)

	return c
}

// Client manages communication with MEWS API
type Client struct {
	// HTTP client used to communicate with the API.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	username string
	password string

	// User agent for client
	userAgent string

	mediaType string
	charset   string

	// Optional function called after every successful request made to the DO APIs
	onRequestCompleted RequestCompletionCallback

	// Services used for communicating with the API
	Notifications *NotificationsService
	Invoices      *InvoicesService
	CreditNotes   *CreditNotesService
	Files         *FilesService
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) Username() string {
	return c.username
}

func (c *Client) SetUsername(username string) {
	c.username = username
}

func (c *Client) Password() string {
	return c.password
}

func (c *Client) SetPassword(password string) {
	c.password = password
}

func (c *Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetProductionMode() {
	c.baseURL = BaseURL
}

func (c *Client) SetTestMode() {
	c.baseURL = BaseURLTest
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) GetEndpointURL(path string) (url.URL, error) {
	baseURL := c.BaseURL()
	apiURL, err := url.Parse(baseURL.String())
	if err != nil {
		return url.URL{}, err
	}

	apiURL.Path = apiURL.Path + path
	return *apiURL, nil
}

func (c *Client) NewRequest(ctx context.Context, method string, URL url.URL, body interface{}) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	req, err := http.NewRequest("POST", URL.String(), buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username(), c.Password())

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// send uuid as header for request/response identification
	uuid := uuid.NewV4()
	req.Header.Add("X-BW-REQUEST-ID", uuid.String())

	// set other headers
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())

	return req, nil
}

// Do sends an API request and returns the API response. The API response is json decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, err
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	err = json.NewDecoder(httpResp.Body).Decode(responseBody)
	if err != nil {
		// create a simple error response
		errorResponse := &ErrorResponse{Response: httpResp}
		errorResponse.Errors.Message = err.Error()
		return httpResp, errorResponse
	}

	return httpResp, nil
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. API error responses are expected to have either no response
// body, or a XML response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	err := checkContentType(r)
	if err != nil {
		errorResponse.Errors.Message = err.Error()
	}

	if r.Header.Get("Content-Length") == "0" {
		errorResponse.Errors.Message = r.Status
		return errorResponse
	}

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert xml to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		errorResponse.Errors.Message = err.Error()
		return errorResponse
	}

	return errorResponse
}

// An ErrorResponse reports the error caused by an API request
// {
//    "version" : "1.0",
//    "errors" : {
//       "validationErrors" : [
//          {
//             "fieldId" : "data.invoiceLine",
//             "fieldMessage" : "instance type (null) does not match any allowed primitive type (allowed: [\"array\"])"
//          }
//       ],
//       "message" : "Required field is missing from the request sent by the API client, or a field in the request does not match the expected pattern. For example, a date is given in a false format.",
//       "id" : "9ee67962-d927-4235-b557-46267e8b743d",
//       "type" : "VALIDATION",
//       "info" : "Required field is missing from the request sent by the API client, or a field in the request does not match the expected pattern. For example, a date is given in a false format.",
//       "code" : "Error.004.0002"
//    }
// }
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	// Fault code
	Version string `json:"version"`

	Errors Errors `json:"errors"`
}

type Errors struct {
	// List of validation errors
	ValidationErrors ValidationErrors `json:"validationErrors"`

	// Fault message
	Message string `json:"message"`

	// ID
	ID string `json:"id"`

	// Type: VALIDATION
	Type string `json:"type"`

	Info string `json:"info"`

	Code string `json:"code"`
}

func (err Errors) Error() string {
	return err.ValidationErrors.Error()
}

type ValidationErrors []ValidationError

func (errs ValidationErrors) Error() string {
	var errors error
	for _, err := range errs {
		errors = multierror.Append(errors, err)
	}
	return errors.Error()
}

type ValidationError struct {
	FieldID      string `json:"fieldId"`
	FieldMessage string `json:"fieldMessage"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.FieldID, e.FieldMessage)
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d (%v %v)",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Errors.Error(), r.Errors.Info)
}

func checkContentType(response *http.Response) error {
	// check content-type (application/soap+xml; charset=utf-8)
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}
