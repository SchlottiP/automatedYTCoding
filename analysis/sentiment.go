package analysis

import (
	"automatedYTCoding/apiadapter"
	"automatedYTCoding/csv"
)

func SentimentAnalysis(devKey string, ids []string, resultPath string, commentPath string) {
	comments := make(map[string][]string, len(ids))
	fullComments := make([]interface{}, len(ids))
	for _, id := range ids {
		videoComments, _ := apiadapter.GetVideoComments(devKey, id)
		//cast to interface
		for _, com := range videoComments {
			fullComments = append(fullComments, com)
		}
		//get comment string for sentiment calculation
		var listOfCommentsAsString []string
		for _, comments := range videoComments {
			listOfCommentsAsString = append(listOfCommentsAsString, comments.TextDisplay)
		}
		comments[id] = listOfCommentsAsString
	}
	result := getSentimentValue(comments)

	var resultAsInterface []interface{}
	for _, v := range result {
		resultAsInterface = append(resultAsInterface, v)
	}
	csv.CreateCSV(resultPath, resultAsInterface)
	csv.CreateCSV(commentPath, fullComments)
}
