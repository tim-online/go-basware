package basware_test

import (
	"log"
	"os"
	"testing"

	uuid "github.com/satori/go.uuid"
	basware "github.com/tim-online/go-basware"
)

func TestDing(t *testing.T) {
	// var err error

	// get username & password
	username := os.Getenv("BASWARE_USERNAME")
	password := os.Getenv("BASWARE_PASSWORD")

	// build client
	client := basware.NewClient(nil, username, password)
	client.SetDebug(true)
	client.SetTestMode()

	params := client.Invoices.NewPostPathParams()
	params.BumID = uuid.NewV4().String()
	requestBody := client.Invoices.NewPostRequestBody()
	_, err := client.Invoices.Post(nil, params, requestBody)
	log.Println(err)
}
