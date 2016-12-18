# Epa Search Scrape
This tool scrapes the EPA search engine, grabbing urls that match a list of queries, and combining them into a large CSV file that lists links

### Purpose

The purpose is to get a list of valid urls that match a given query from epa.gov

### How to Install

The Go language is not necessary to run the tool, a compiled binary is attached for Mac OS

If you're not running os x, you'll need to compile the code. First make sure you've downloaded the [Go Programming language](https://golang.org/dl) and run `go build` from the epa_search_urls directory.
For windows users:
1. Set environment variable : `set GOPATH=c:\YOUR_DESIRED_PATH`
2. Download required imports : `go get github.com/PuerkitoBio/goquery`
 
### How to Run 

1. Edit queries.csv, placing your search queries one per line.
2. using the command line, run:
	`./epa_search_urls` ( if using compiled binary) 
	or  `go run main.go`
3. Wait for the tool to finish scraping.

Once finished, a series of csv files representing each page will be in the *results* directory, and a combined file of *all* results will be created at


### Next Steps
Sort & De-Duplicate the concatinated results
