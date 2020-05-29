package apiadapter

import (
	"fmt"
	"google.golang.org/api/youtube/v3"
	"regexp"
	"strconv"
	"strings"
)

type Video struct {
	Id                   string
	DurationInSeconds    int
	Category             string
	channelTitle         string
	Language             string
	Description          string
	PublishedAt          string
	TagsList             string
	Title                string
	CommentCount         uint64
	LikeDislikeRatio     float64
	ViewReactionRatio    float64
	DislikeCount         uint64
	LikeCount            uint64
	ViewCount            uint64
	Subscribers          uint64
	ViewSubscriberRation float64
}

func GetVideoData(devKey string, videoIds []string) map[string]*Video {
	client := GetClient(devKey)
	response, err := client.Videos.List("snippet,contentDetails,statistics, topicDetails").Id(strings.Join(videoIds, ",")).Do()
	if err != nil {
		panic(err)
	}
	categories := make(map[string]string)
	//map video response to the video struct
	results := make(map[string]*Video)
	for _, vid := range response.Items {
		results[vid.Id], categories = getVideoFromResponse(vid, categories, client)
	}
	//get Subscription data
	channelIdsToVideo := make(map[string]string, len(videoIds))
	channelIds := make([]string, len(videoIds))
	for i, vid := range response.Items {
		channelIdsToVideo[vid.Snippet.ChannelId] = vid.Id
		channelIds[i] = vid.Snippet.ChannelId
	}
	channelResponse, err := client.Channels.List("statistics").Id(strings.Join(channelIds, ",")).Do()
	if err != nil {
		panic(err)
	}
	for _, channel := range channelResponse.Items {
		id := channelIdsToVideo[channel.Id]
		results[id].Subscribers = channel.Statistics.SubscriberCount
		results[id].ViewSubscriberRation = float64(results[id].ViewCount) / float64(channel.Statistics.SubscriberCount)
	}

	return results
}

func getVideoFromResponse(vid *youtube.Video, categories map[string]string, client *youtube.Service) (*Video, map[string]string) {
	if _, in := categories[vid.Snippet.CategoryId]; !in {
		catResp, err := client.VideoCategories.List("snippet").Id(vid.Snippet.CategoryId).Do()
		if err != nil {
			fmt.Errorf("error getting categorie for video %v %v", vid.Id, err)
		}
		categories[vid.Snippet.CategoryId] = catResp.Items[0].Snippet.Title

	}
	video := &Video{
		Id:                vid.Id,
		DurationInSeconds: convertDuration(vid.ContentDetails.Duration),
		channelTitle:      vid.Snippet.ChannelTitle,
		Language:          vid.Snippet.DefaultAudioLanguage,
		Description:       vid.Snippet.Description,
		PublishedAt:       vid.Snippet.PublishedAt,
		TagsList:          strings.Join(vid.Snippet.Tags, ","),
		Title:             vid.Snippet.Title,
		Category:          categories[vid.Snippet.CategoryId],
		CommentCount:      vid.Statistics.CommentCount,
		DislikeCount:      vid.Statistics.DislikeCount,
		LikeCount:         vid.Statistics.LikeCount,
		ViewCount:         vid.Statistics.ViewCount,
	}
	video.LikeDislikeRatio = float64(video.LikeCount) / float64(video.DislikeCount)
	video.ViewReactionRatio = float64(video.ViewCount) / float64(video.LikeCount+video.DislikeCount)

	return video, categories
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
