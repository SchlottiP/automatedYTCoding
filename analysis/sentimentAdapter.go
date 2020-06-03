package analysis

import (
	"github.com/cdipaolo/sentiment"
)

type SentimentResult struct {
	Id string
	//positive/negative is the average of the score of all comments, it uses -1 for negative and +1 for positive comments
	All   float64
	NoPos int
	NoNeg int
}

func getSentimentValue(videoIdsWithComments map[string][]string) map[string]*SentimentResult {
	model, e := sentiment.Restore()
	if e != nil {
		panic(e)
	}
	result := make(map[string]*SentimentResult, len(videoIdsWithComments))
	for videoId, comments := range videoIdsWithComments {
		if comments == nil {
			result[videoId] = &SentimentResult{
				Id:    videoId,
				NoPos: -1,
				NoNeg: -1,
				All:   -2,
			}
			continue
		}
		curSentResult := &SentimentResult{
			Id: videoId,
		}
		for _, comment := range comments {
			res := model.SentimentAnalysis(comment, sentiment.English)
			if res.Score == uint8(0) {
				curSentResult.NoNeg++
				curSentResult.All--
			} else {
				curSentResult.NoPos++
				curSentResult.All++
			}
		}
		curSentResult.All = curSentResult.All / float64(len(comments))
		result[videoId] = curSentResult
	}
	return result
}
