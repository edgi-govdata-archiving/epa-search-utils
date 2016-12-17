# Epa Search Scrape
This tool scrapes the EPA search engine, grabbing urls that match a list of queries, and combining them into a large CSV file that lists links

### Purpose

The purpose is to get a list of valid urls that match a given query from epa.gov

### How to Run

The go language is not necessary to run the tool, a compiled binary is attached

1. If you're not running os x, you'll need to compile the code. First make sure you've downloaded the [Go Programming languge][golang.org/download] and run `go build` from the epa_search_urls directory.
2. Edit queries.csv, placing a query on each line.
3. using the command line, run:
	`./epa_search_urls`
4. Wait for the tool to finish scraping.

Once finished, a series of csv files representing each page will be in the *results* directory, and a combined file of *all* results will be created at


### Next Steps
Sort & De-Duplicate the concatinated results