package main

import (
	"automatedYTCoding/analysis"
	"automatedYTCoding/apiadapter"
	"automatedYTCoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var developerKey string

func main() {
	developerKey = os.Getenv("devkey")
	if developerKey == "" {
		panic(errors.New("missing developer key. Must be set via the envirment"))
	}

	// Subcommands
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	videoData := flag.NewFlagSet("videos", flag.ExitOnError)
	fileData := flag.NewFlagSet("file", flag.ExitOnError)

	//List subcommand flag pointers
	query := listCommand.String("query", "23AndMe", "Search term")
	maxResults := listCommand.Int64("max-results", 25, "Max YouTube results, optional ")
	publishedAfter := listCommand.String("after", "", "Date PublishedAfter in Format DD.MM.YYYY, optional")

	// video subcommand flag pointers
	ids := videoData.String("ids", "", "ids of videos, comma seperated")

	resultPath := fileData.String("path", os.TempDir(), "path to the result file")

	if len(os.Args) < 2 {
		fmt.Println("apiadapter or count subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		list(listCommand, resultPath, query, publishedAfter, maxResults)
	case "videos":
		listVideos(videoData, fileData, ids, resultPath)
	case "sentiment":
		listSentiment(videoData, fileData, ids, resultPath)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func listSentiment(videoData *flag.FlagSet, fileData *flag.FlagSet, ids *string, folderForResult *string) {
	_ = videoData.Parse(os.Args[2:])
	if len(os.Args) == 2 {
		_ = fileData.Parse(os.Args[3:])
	}
	result, err := filepath.Abs(*folderForResult + "/sentiment.csv")
	commentPath, err := filepath.Abs(*folderForResult + "/comments.csv")
	if err != nil {
		panic(err)
	}
	videoIds := strings.Split(*ids, ",")
	analysis.SentimentAnalysis(developerKey, videoIds, result, commentPath)
}

func makeInterfaceSliceVideo(data map[string]*apiadapter.Video) []interface{} {
	var res []interface{}
	for _, v := range data {
		res = append(res, v)
	}
	return res
}
func makeInterfaceSliceSearch(data []*apiadapter.VideoData) []interface{} {
	var res []interface{}
	for _, v := range data {
		res = append(res, v)
	}
	return res
}

func listVideos(videoData *flag.FlagSet, fileData *flag.FlagSet, ids *string, resultPath *string) {
	//FIRST VIDEO THEN FILE INFORMATION!
	_ = videoData.Parse(os.Args[2:])
	_ = fileData.Parse(os.Args[3:])

	// Required Flags
	if *ids == "" {
		videoData.PrintDefaults()
		os.Exit(1)
	}
	path, err := filepath.Abs(*resultPath + "/result.csv")
	if err != nil {
		panic(err)
	}
	fmt.Printf("resultFile: %v", path)
	idList := strings.Split(*ids, ",")
	csv.CreateCSV(path, makeInterfaceSliceVideo(apiadapter.GetVideoData(developerKey, idList)))
}

func list(listCommand *flag.FlagSet, resultPath *string, query *string, publishedAfter *string, maxResults *int64) {
	_ = listCommand.Parse(os.Args[2:])

	// Required Flags
	if *query == "" {
		listCommand.PrintDefaults()
		os.Exit(1)
	}
	//optional
	var after *time.Time
	if *publishedAfter != "" {
		date, err := time.Parse("02.01.2006", *publishedAfter)
		if err != nil {
			panic(fmt.Errorf("published After Date is not valid. Format: DD.MM.YYYY"))
		}
		after = &date
	}
	path, err := filepath.Abs(*resultPath + "/searchResult.csv")
	if err != nil {
		panic(err)
	}
	csv.CreateCSV(path, makeInterfaceSliceSearch(apiadapter.List(developerKey, *query, *maxResults, after)))
}
