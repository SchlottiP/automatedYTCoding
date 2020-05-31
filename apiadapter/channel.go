package apiadapter

import (
	"fmt"
	"google.golang.org/api/youtube/v3"
	"strings"
)

func getChannelData(service *youtube.Service, channelIds []string) []*youtube.Channel {
	var channels []*youtube.Channel
	length := len(channelIds)
	for i := 0; i < length; i += 10 {
		end := i + 10
		if end >= length {
			end = length
		}
		response, err := service.Channels.List("statistics").Id(strings.Join(channelIds[i:end], ",")).Do()
		if err != nil {
			fmt.Printf("error listing channels by id (ids: %v) %v ", strings.Join(channelIds, ","), err)
		}
		channels = append(channels, response.Items...)
	}
	return channels
}
