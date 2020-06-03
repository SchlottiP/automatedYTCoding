# automatedYTCoding
Commandline Tool to get video data for a regression analysis of videos about Direct-to-consumer genetic tests. Created for the seminar "Emerging Trends in Digital Health" at Karlsruher Institut of Technology
uses the youtube api v3 (https://developers.google.com/youtube/v3/docs)

## Get and Build
`git clone git@github.com:SchlottiP/automatedYTCoding.git`

`cd automatedYTCoding`

`go build`

## Commands
To run any command the developer key of the youtube api is needed (in the environment `devkey=XXX`) (See https://developers.google.com/youtube/v3/getting-started)
### list
Quota impact = 100 units 

Lists videos with id, date and title
- query: Searchterm.  | and - allowed (required)
- after: search for videos after this date. Format: "DD.MM.YYYY"
- maxResults: max. amount of results, default 25, uses pagination and not the query parameter of the youtube api (because the api only allows up to 50 videos as maxResult value)

### videos
Quota impact =2+2+2+2+2= 10 per video and 2 per Video Category 

Lists the video information for each video in a csv file. 
Videos are ordered by views, default language is english, and the videos have medium length (4-20min)
- ids: list of ids "id1,id2" (required)
- path: path of the folder for the result file (result.csv), default: Temp Folder

### all
Quota impact: videos + list

Combines list and video command: First search for videos, then uses the resulting ids to list the video information for each video in a csv file. 
- path:  path of the folder for the result file (result.csv), default: Temp Folder
- query, after and maxResult, see list command

### sentiment
Quota impact: 2 per video

performs a sentiment analysis (https://github.com/cdipaolo/sentiment). The command uses up to 100 comments of each video 
and performance a sentiment analysis for each comment.
The score for each comment is -1(negative) or 1(positive). Result is a csv file with Id (video ID), All (average score), 
NoPos (number of positive comments) ,and NoNeg (number of negative comments) for each video

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