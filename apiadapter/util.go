package apiadapter

import (
	"fmt"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"net/http"
)

func GetClient(developerKey string) *youtube.Service {
	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		fmt.Errorf("Error creating new YouTube client: %v", err)
		panic(err)
	}
	return service
}
