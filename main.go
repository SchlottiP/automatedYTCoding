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
		list(listCommand, query, publishedAfter, maxResults)
	case "videos":
		listVideos(videoData, fileData, ids, resultPath)
	case "all":
		listAllVideos(listCommand, fileData, query, publishedAfter, maxResults, resultPath)
	case "sentiment":
		listSentiment(videoData, fileData, ids, resultPath)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func listSentiment(videoData *flag.FlagSet, fileData *flag.FlagSet, ids *string, resultPath *string) {
	videoData.Parse(os.Args[2:])
	if len(os.Args) == 2 {
		fileData.Parse(os.Args[3:])
	}
	path, err := filepath.Abs(*resultPath + "/sentiment.csv")
	if err != nil {
		panic(err)
	}
	videoIds := strings.Split(*ids, ",")
	analysis.SentimentAnalysis(developerKey, videoIds, path)
}

func listAllVideos(listCommand *flag.FlagSet, fileData *flag.FlagSet, query *string, publishedAfter *string, maxResults *int64, resultPath *string) {
	//FIRST FILE THEN SEARCH INFORMATION!
	if os.Args[2] == "path" {
		fileData.Parse(os.Args[2:])
		listCommand.Parse(os.Args[3:])
	} else {
		listCommand.Parse(os.Args[2:])
	}

	// Required Flags
	if *query == "" {
		listCommand.PrintDefaults()
		os.Exit(1)
	}
	//optional
	var after *time.Time
	if *publishedAfter != "" {
		time, err := time.Parse("02.01.2006", *publishedAfter)
		if err != nil {
			fmt.Errorf("published After Date is not valid. Format: DD.MM.YYYY")
			os.Exit(1)
		}
		after = &time
	}
	videos := apiadapter.List(developerKey, *query, *maxResults, after, false)
	ids := make([]string, len(videos))
	for i, vid := range videos {
		ids[i] = vid.Id.VideoId
	}
	path, err := filepath.Abs(*resultPath + "/result" + time.Now().Format("02-01-2006-15_04_05") + ".csv")
	if err != nil {
		panic(err)
	}
	fmt.Printf("query %v", *query)
	fmt.Printf("resultFile: %v", path)
	csv.CreateCSV(path, makeInterfaceSlice(apiadapter.GetVideoData(developerKey, ids)))
}

func makeInterfaceSlice(data map[string]*apiadapter.Video) []interface{} {
	var res []interface{}
	for _, v := range data {
		res = append(res, v)
	}
	return res
}

func listVideos(videoData *flag.FlagSet, fileData *flag.FlagSet, ids *string, resultPath *string) {
	//FIRST VIDEO THEN FILE INFORMATION!
	videoData.Parse(os.Args[2:])
	fileData.Parse(os.Args[3:])

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
	csv.CreateCSV(path, makeInterfaceSlice(apiadapter.GetVideoData(developerKey, idList)))
}

func list(listCommand *flag.FlagSet, query *string, publishedAfter *string, maxResults *int64) {
	listCommand.Parse(os.Args[2:])

	// Required Flags
	if *query == "" {
		listCommand.PrintDefaults()
		os.Exit(1)
	}
	//optional
	var after *time.Time
	if *publishedAfter != "" {
		time, err := time.Parse("02.01.2006", *publishedAfter)
		if err != nil {
			fmt.Errorf("published After Date is not valid. Format: DD.MM.YYYY")
			os.Exit(1)
		}
		after = &time
	}
	apiadapter.PrintIDs(apiadapter.List(developerKey, *query, *maxResults, after, true))
}
