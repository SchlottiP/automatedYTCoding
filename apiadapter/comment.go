package apiadapter

import "google.golang.org/api/youtube/v3"

func GetVideoComments(devKey string, id string) ([]*youtube.CommentSnippet, error) {
	client := GetClient(devKey)
	call := client.CommentThreads.List("id,snippet").TextFormat("plainText").VideoId(id).MaxResults(100)
	res, err := call.Do()
	if err != nil {
		return nil, err
	}
	comments := make([]*youtube.CommentSnippet, 0)
	comments = putResultInComments(res, comments)
	res, err = call.PageToken(res.NextPageToken).Do()
	if err != nil {
		return nil, err
	}
	comments = putResultInComments(res, comments)
	return comments, nil
}

func putResultInComments(res *youtube.CommentThreadListResponse, comments []*youtube.CommentSnippet) []*youtube.CommentSnippet {
	for _, com := range res.Items {
		comments = append(comments, com.Snippet.TopLevelComment.Snippet)
	}
	return comments
}
