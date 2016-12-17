package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	queryData, err := ioutil.ReadFile("queries.csv")
	if err != nil {
		panic(err)
	}

	csvr := csv.NewReader(bytes.NewBuffer(queryData))
	for {
		query, err := csvr.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		err = CrawlEntry(query[0])
		if err != nil {
			panic(err)
		}
	}

	if err := ConcatResults(); err != nil {
		panic(err)
	}
}

func SearchPageUrl(query string, page, resultCount int) string {
	start := ((page - 1) * resultCount) + 1
	return fmt.Sprintf("https://nlquery.epa.gov/epasearch/epasearch?querytext=%s&start=%d&results_per_page=%d&cluster=no&filter=&fld=&url_directory=&federated=no&max_results=1000&result_template=2col.ftl&areaname=&areapagehead=epafiles_pagehead&areapagefoot=epafiles_pagefoot&areasidebar=search_sidebar&stylesheet=&sort=term_relevancy&faq=true&cluster=both&sessionid=2950002FD05190D4090549F0E6EB934A&doctype=all&typeofsearch=epa&site=epa_default", query, start, resultCount)
}

func CrawlEntry(query string) error {
	page := 1
	resultCount := 500
	for {
		url := SearchPageUrl(query, page, resultCount)
		fmt.Println(url)

		doc, err := goquery.NewDocument(url)
		if err != nil {
			return err
		}

		kontinue, err := CrawlPage(url, doc, query, page, resultCount)
		if err != nil {
			return err
		}

		if !kontinue {
			break
		}
		time.Sleep(time.Second)
		page++
	}

	return nil
}

func CrawlPage(url string, doc *goquery.Document, query string, page, resultCount int) (bool, error) {
	kontinue := true
	sel := doc.Find("#main-content dl").Find("a")

	buf := &bytes.Buffer{}
	csvw := csv.NewWriter(buf)
	now := time.Now().Unix()

	fmt.Println("num entries:", len(sel.Nodes))
	for i := range sel.Nodes {
		node := sel.Eq(i)
		href, _ := node.Attr("href")
		// fmt.Println(i+1, ",", href)
		csvw.Write([]string{strconv.FormatInt(now, 10), href})
	}

	if len(sel.Nodes) < resultCount {
		kontinue = false
	}

	csvw.Flush()
	if err := ioutil.WriteFile(fmt.Sprintf("results/%s-%d-%d.csv", query, page, resultCount), buf.Bytes(), os.ModePerm); err != nil {
		return kontinue, err
	}

	return kontinue, nil
}

func ConcatResults() error {
	buf := &bytes.Buffer{}
	wr := csv.NewWriter(buf)
	err := filepath.Walk("results/", func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.IsDir() && filepath.Ext(fi.Name()) == ".csv" {
			fmt.Println(filepath.Join("results", fi.Name()))
			data, err := ioutil.ReadFile(filepath.Join("results", fi.Name()))
			if err != nil {
				return err
			}

			r := csv.NewReader(bytes.NewBuffer(data))
			rows, err := r.ReadAll()
			if err != nil {
				return err
			}

			return wr.WriteAll(rows)
		}
		return nil
	})

	if err != nil {
		return err
	}

	wr.Flush()
	if err := ioutil.WriteFile("results.csv", buf.Bytes(), os.ModePerm); err != nil {
		return err
	}

	return nil
}
