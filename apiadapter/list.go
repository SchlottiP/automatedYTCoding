package apiadapter

import (
	"fmt"
	"google.golang.org/api/youtube/v3"
	"log"
	"time"
)

// List uses keywords, maxResult and if present publishedAfter
// Defaults: order for viewCount, videoduration medium (4-20 min), relevance language: English
func List(developerKey string, keywords string, maxResult int64, publishedAfter *time.Time) []*youtube.SearchResult {

	service := GetClient(developerKey)

	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		Q(keywords).MaxResults(maxResult).Order("viewCount").VideoDuration("medium").RelevanceLanguage("en").Type("video")
	if publishedAfter != nil {
		fmt.Printf("after: %v %v", publishedAfter.Format(time.RFC822), publishedAfter.Format(time.RFC3339))
		call.PublishedAfter(publishedAfter.Format(time.RFC3339))
	}
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error requesting Api: %v", err)
	}
	return response.Items
}

func PrintIDs(items []*youtube.SearchResult) {
	videos := make(map[string]string)
	// Iterate through each item and add it to the correct apiadapter.
	for _, item := range items {
		videos[item.Id.VideoId] = item.Snippet.Title + " published: " + item.Snippet.PublishedAt
	}

	for id, title := range videos {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
