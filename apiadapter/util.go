package apiadapter

import (
	"fmt"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"net/http"
)

var client *youtube.Service

func GetClient(developerKey string) *youtube.Service {
	if client != nil {
		return client
	}
	c := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}
	client, err := youtube.New(c)
	if err != nil {
		fmt.Errorf("Error creating new YouTube client: %v", err)
		panic(err)
	}
	return client

}
