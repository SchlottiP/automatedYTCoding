package analysis

import "github.com/cdipaolo/sentiment"

func getSentimentValue(videoIdsWithComments map[string][]string) map[string]int {
	model, e := sentiment.Restore()
	if e != nil {
		panic(e)
	}
	result := make(map[string]int, len(videoIdsWithComments))
	for videoId, comments := range videoIdsWithComments {
		if comments == nil {
			result[videoId] = -1
			continue
		}
		result[videoId] = 0
		for _, comment := range comments {
			result[videoId] += int(model.SentimentAnalysis(comment, sentiment.English).Score)
		}
	}
	return result
}
