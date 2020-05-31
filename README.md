# automatedYTCoding
Commandline Tool to get video data for a regression analysis of videos about Direct-to-consumer genetic tests. Created for the seminar "Emerging Trends in Digital Health" at Karlsruher Institut of Technology
uses the youtube api v3 (https://developers.google.com/youtube/v3/docs)

## Get and Build
`git clone git@github.com:SchlottiP/automatedYTCoding.git`

`cd automatedYTCoding`

`go build`

## Commands
To run any command the developer key of the youtube api is needed (in the environment `devkey=XXX`) (see https://developers.google.com/youtube/v3/getting-started)
### list
Quota impact = 100 units 
Lists videos with id, date and title
- query: Searchterm.  | and - allowed (required)
- after: search for videos after this date. Format: "DD.MM.YYYY"
- maxResults: max. amount of results, default 25

### videos
Quota impact =2+2+2+2+2= 10 per video and 2 per Video Category 

Lists the video information for each video in a csv file. 
Videos or ordered by views, default language is english and the videos have medium length (4-20min)
- ids: list of ids "id1,id2" (required)
- path: path of the folder for the result file (result.csv), default: Temp Folder

### all
Quota impact: videos + list
combines list and video command: first search for videos, then uses the resulting id to list the video information for each video in a csv file. 
- path:  path of the folder for the result file (result.csv), default: Temp Folder
- query, after and maxResult, see list command

### sentiment
Quota impact: 2 per video
performers a sentiment analysis (https://github.com/cdipaolo/sentiment). the command uses max. 50 comments of each video, performance a sentiment analysis for each comment and sums up the results. 
The score of one comment is 0(negative) or 1(positive). A higher sum indicates that more comments are positive. 
The resulting file contains the id and score for each video. 
the sentiment score is -1, if any error occurs while getting the comments (video has no comments, comments are not activated, video was deleted or any other error)


- ids: list of ids "id1,id2" (required)
- path:  path of the folder for the result file (result.csv), default: Temp Folder

## Video Data
the data for each video in the csv files (command all and videos):
- Id: text, videoId
- DurationInSeconds: int
- Category: string
- channelTitle: string
- Language: string, coded ("en")
- Description: string
- PublishedAt: string (e.g. 07-04-15 19:22)
- TagsList: string (comma separated)
- Title string
- CommentCount: int
- LikeDislikeRatio: float with 5 decimal places (no rounding)
- ViewReactionRatio: float with 5 decimal places (no rounding)
- DislikeCount: int
- LikeCount: int
- ViewCount: int
- Subscribers: int
- ViewSubscriberRation: float with 5 decimal places (no rounding)