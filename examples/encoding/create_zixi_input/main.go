package main

import (
	"log"

	"github.com/dakdikduk/bitmovin-go/bitmovin"
	"github.com/dakdikduk/bitmovin-go/models"
	"github.com/dakdikduk/bitmovin-go/services"
)

func main() {
	bitmovin := bitmovin.NewBitmovinDefaultTimeout("YOUR API KEY", "https://api.bitmovin.com/v1/")

	log.Printf("Creating Zixi Input")
	zixiIS := services.NewZixiInputService(bitmovin)
	zixiInput := &models.ZixiInput{
		Name:   stringToPtr("My test Zixi Input"),
		Host:   stringToPtr("test.com"),
		Port:   int64ToPtr(2088),
		Stream: stringToPtr("test-stream"),
	}

	httpResp, err := zixiIS.Create(zixiInput)
	if err != nil {
		log.Fatalf("Could not create Zixi Input: %v", err)
	}
	log.Printf("Created Zixi Pull Input with ID: \n%s\n", httpResp.Data.Result.ID)
}

func stringToPtr(s string) *string {
	return &s
}

func intToPtr(i int) *int {
	return &i
}

func int64ToPtr(i int64) *int64 {
	return &i
}
