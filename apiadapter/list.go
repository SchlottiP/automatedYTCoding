package apiadapter

import (
	"fmt"
	"google.golang.org/api/youtube/v3"
	"log"
	"time"
)

// List uses keywords, maxResult and if present publishedAfter
// Defaults: order for viewCount, videoduration medium (4-20 min), relevance language: English
func List(developerKey string, keywords string, maxResult int64, publishedAfter *time.Time, toPrint bool) []*youtube.SearchResult {

	service := GetClient(developerKey)

	// Make the API call to YouTube.
	part := "id"
	if toPrint {
		part += ",snippet"
	}
	call := service.Search.List(part).
		Q(keywords).Order("viewCount").VideoDuration("medium").RelevanceLanguage("en").Type("video")
	if publishedAfter != nil {
		fmt.Printf("after: %v %v", publishedAfter.Format(time.RFC822), publishedAfter.Format(time.RFC3339))
		call.PublishedAfter(publishedAfter.Format(time.RFC3339))
	}
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error requesting Api: %v", err)
	}
	if response.PageInfo.ResultsPerPage >= maxResult {
		return response.Items
	}
	var results []*youtube.SearchResult
	results = append(results, response.Items...)
	for int64(len(results)) < maxResult {
		response, err = call.PageToken(response.NextPageToken).Do()
		if err != nil {
			log.Fatalf("Error requesting Api: %v", err)
		}
		missing := maxResult - int64(len(results))
		if missing >= response.PageInfo.ResultsPerPage {
			results = append(results, response.Items...)
		} else {
			results = append(results, response.Items[0:missing]...)
		}
	}
	return results
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
