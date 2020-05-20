# automatedYTCoding
Commandline Tool to get video data for a regression analysis of videos about Direct-to-consumer genetic tests. Created for the seminar "Emerging Trends in Digital Health" at Karlsruher Institut of Technology

## Get and Run

## Get and Build
`git clone git@github.com:SchlottiP/automatedYTCoding.git`

`cd automatedYTCoding`

`go build`

## Commands
### list
- query: Searchterm.  | and - allowed (required)
- after: search for videos after this date. Format: "DD.MM.YYYY"
- maxResults: max. amount of results, default 25

### videos
- ids: list of ids "id1,id2" (required)
- path: path of the folder for the result file (result.csv), default: Temp Folder

### all
- path:  path of the folder for the result file (result.csv), default: Temp Folder
- query, after and maxResult, see list command

