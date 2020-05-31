package analysis

import (
	"automatedYTCoding/apiadapter"
	"automatedYTCoding/csv"
)

func SentimentAnalysis(devKey string, ids []string, filePath string) {
	comments := make(map[string][]string, len(ids))
	for _, id := range ids {
		comments[id], _ = apiadapter.GetVideoComments(devKey, id)
	}
	result := getSentimentValue(comments)
	csv.CreateCSV(filePath, result)
}
