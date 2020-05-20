package apiadapter

import (
	"regexp"
	"strconv"
	"strings"
)

type Video struct {
	Id                string
	DurationInSeconds int
	CategoryId        string
	channelTitle      string
	DefaultLanguage   string
	Description       string
	PublishedAt       string
	TagsList          string
	Title             string
	CommentCount      uint64
	DislikeCount      uint64
	FavoriteCount     uint64
	LikeCount         uint64
	ViewCount         uint64
	TopicCategories   []string
}

func GetVideoData(devKey string, videoIds []string) map[string]*Video {
	client := GetClient(devKey)
	response, err := client.Videos.List("snippet,contentDetails,statistics, topicDetails").Id(strings.Join(videoIds, ",")).Do()
	if err != nil {
		panic(err)
	}
	results := make(map[string]*Video)
	for _, vid := range response.Items {
		categories := make([]string, 0)
		if vid.TopicDetails != nil {
			categories = vid.TopicDetails.TopicCategories
		}
		results[vid.Id] = &Video{
			Id:                vid.Id,
			DurationInSeconds: convertDuration(vid.ContentDetails.Duration),
			CategoryId:        vid.Snippet.CategoryId,
			channelTitle:      vid.Snippet.ChannelTitle,
			DefaultLanguage:   vid.Snippet.DefaultLanguage,
			Description:       vid.Snippet.Description,
			PublishedAt:       vid.Snippet.PublishedAt,
			TagsList:          strings.Join(vid.Snippet.Tags, ","),
			Title:             vid.Snippet.Title,
			CommentCount:      vid.Statistics.CommentCount,
			DislikeCount:      vid.Statistics.DislikeCount,
			FavoriteCount:     vid.Statistics.FavoriteCount,
			LikeCount:         vid.Statistics.LikeCount,
			ViewCount:         vid.Statistics.ViewCount,
			TopicCategories:   categories,
		}
	}
	return results
}

func convertDuration(duration string) int {
	durationRegex := regexp.MustCompile(`PT(?P<minutes>\d+M)?(?P<seconds>\d+S)?`)
	matches := durationRegex.FindStringSubmatch(duration)

	minutes := ParseInt64(matches[1])
	seconds := ParseInt64(matches[2])

	return minutes*60 + seconds
}

func ParseInt64(value string) int {
	if len(value) == 0 {
		return 0
	}
	parsed, err := strconv.Atoi(value[:len(value)-1])
	if err != nil {
		return 0
	}
	return parsed
}
