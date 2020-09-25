package apiadapter

import "google.golang.org/api/youtube/v3"

func GetVideoComments(devKey string, id string) ([]*youtube.CommentSnippet, error) {
	client := GetClient(devKey)
	res, err := client.CommentThreads.List("id,snippet").TextFormat("plainText").VideoId(id).MaxResults(100).Do()
	if err != nil {
		return nil, err
	}
	comments := make([]*youtube.CommentSnippet, len(res.Items))
	for i, com := range res.Items {
		comments[i] = com.Snippet.TopLevelComment.Snippet
	}
	return comments, nil
}
