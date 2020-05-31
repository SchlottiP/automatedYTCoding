package apiadapter

import (
	"fmt"
	"google.golang.org/api/youtube/v3"
	"regexp"
	"strconv"
	"strings"
)

// Video contains all Data needed for a regression analysis based on youtube videos
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
	client, videos := getVideos(devKey, videoIds)
	categories := make(map[string]string)
	//map video response to the video struct
	results := make(map[string]*Video)
	for _, vid := range videos {
		results[vid.Id], categories = getVideoFromResponse(vid, categories, client)
	}
	//get Subscription data
	channelIdsToVideo := make(map[string]string, len(videoIds))
	channelIds := make([]string, len(videoIds))
	for i, vid := range videos {
		channelIdsToVideo[vid.Snippet.ChannelId] = vid.Id
		channelIds[i] = vid.Snippet.ChannelId
	}
	channels := getChannelData(client, channelIds)
	for _, channel := range channels {
		id := channelIdsToVideo[channel.Id]
		results[id].Subscribers = channel.Statistics.SubscriberCount
		results[id].ViewSubscriberRation = float64(results[id].ViewCount) / float64(channel.Statistics.SubscriberCount)
	}

	return results
}

func getVideos(devKey string, videoIds []string) (*youtube.Service, []*youtube.Video) {
	client := GetClient(devKey)
	var videos []*youtube.Video
	length := len(videoIds)
	for i := 0; i < length; i += 10 {
		end := i + 10
		if end >= length {
			end = length
		}
		response, err := client.Videos.List("snippet,contentDetails,statistics, topicDetails").Id(strings.Join(videoIds[i:end], ",")).Do()
		if err != nil {
			fmt.Printf("error listing videos by id (ids: %v) %v ", strings.Join(videoIds, ","), err)
			panic(err)
		}
		videos = append(videos, response.Items...)
	}

	return client, videos
}

func getVideoFromResponse(vid *youtube.Video, categories map[string]string, client *youtube.Service) (*Video, map[string]string) {
	if _, in := categories[vid.Snippet.CategoryId]; !in {
		catResp, err := client.VideoCategories.List("snippet").Id(vid.Snippet.CategoryId).Do()
		if err != nil {
			fmt.Printf("error getting categorie for video %v %v", vid.Id, err)
			categories[vid.Snippet.CategoryId] = "Error!"
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
