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
	var resultAsInterface []interface{}
	for _, v := range result {
		resultAsInterface = append(resultAsInterface, v)
	}
	csv.CreateCSV(filePath, resultAsInterface)
}
