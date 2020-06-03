package apiadapter

func GetVideoComments(devKey string, id string) ([]string, error) {
	client := GetClient(devKey)
	res, err := client.CommentThreads.List("id,snippet").TextFormat("plainText").VideoId(id).MaxResults(100).Do()
	if err != nil {
		return nil, err
	}
	comments := make([]string, len(res.Items))
	for i, com := range res.Items {
		comments[i] = com.Snippet.TopLevelComment.Snippet.TextDisplay
	}
	return comments, nil
}
