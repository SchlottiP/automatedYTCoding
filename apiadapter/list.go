package apiadapter

import (
	"fmt"
	"github.com/biter777/countries"
	"google.golang.org/api/youtube/v3"
	"log"
	"strings"
	"time"
)

type VideoData struct {
	Id           string
	SearchKey    string
	Title        string
	Description  string
	PublishedAt  string
	ChannelTitle string
	ChannelId    string
	nr           int
}

// List uses keywords, maxResult and if present publishedAfter
// Defaults: order for viewCount, region Code: US
func List(developerKey string, keywords string, maxResult int64, publishedAfter *time.Time) []*VideoData {

	service := GetClient(developerKey)
	keywordsList := strings.Split(keywords, ";")
	var resultData []*VideoData

	// Make the API call to YouTube.
	for _, word := range keywordsList {
		resultData = append(resultData, makeCall(service, word, publishedAfter, maxResult)...)
	}
	return resultData
}

func makeCall(service *youtube.Service, keywords string, publishedAfter *time.Time, maxResult int64) []*VideoData {
	part := "id, snippet"
	call := service.Search.List(part).
		Q(keywords).Order("viewCount").MaxResults(50).Type("video").RegionCode(countries.US.Alpha2())
	if publishedAfter != nil {
		fmt.Printf("after: %v %v", publishedAfter.Format(time.RFC822), publishedAfter.Format(time.RFC3339))
		call.PublishedAfter(publishedAfter.Format(time.RFC3339))
	}
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error requesting Api: %v", err)
	}
	var results []*youtube.SearchResult
	if response.PageInfo.ResultsPerPage >= maxResult {
		results = response.Items
	} else {
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
	}
	nr := 1
	formated := make([]*VideoData, len(results))
	for _, res := range results {
		formated = append(formated, &VideoData{
			Id:           res.Id.VideoId,
			SearchKey:    keywords,
			Title:        res.Snippet.Title,
			Description:  res.Snippet.Description,
			PublishedAt:  res.Snippet.PublishedAt,
			ChannelTitle: res.Snippet.ChannelTitle,
			ChannelId:    res.Snippet.ChannelId,
			nr:           nr,
		})
		nr++
	}
	return formated
}
