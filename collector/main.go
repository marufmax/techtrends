package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
)

// https://jobs.bdjobs.com/jobsearch.asp?fcatId=8&icatId=&JobKeyword=php
type Job struct {
	//	CatId   int    `json:"cat_id"`
	Keyword string `json:"keyword"`
	Count   int    `json:"count"`
}

const CACHE_DIR = "./storage/cache"

func main() {
	count := getCount("Java Developer")

	fmt.Println(count)
}

func getRandomUserAgent() string {
	var userAgents = []string{
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36 Edge/14.14393",
		"Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 Firefox/95.0",
	}

	return userAgents[rand.Int()%len(userAgents)]
}

func getCount(keyword string) int {
	collector := colly.NewCollector(
		colly.AllowedDomains("jobs.bdjobs.com", "www.jobs.bdjobs.com"),
		colly.UserAgent(getRandomUserAgent()),
		colly.CacheDir(CACHE_DIR),
	)

	count := 0

	collector.OnHTML("#TopTotalRecord", func(e *colly.HTMLElement) {
		jobCount, err := getJobCount(e.Text)
		if err != nil {
			fmt.Println("Error in parsing value. " + err.Error())
		}
		count = jobCount
	})

	// Before making a request print "Visiting ..."
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	collector.Visit(getUrl(keyword))

	return count
}

func getUrl(keyword string) string {
	return "https://jobs.bdjobs.com/jobsearch.asp?fcatId=8&icatId=&JobKeyword=" + url.QueryEscape(keyword)
}

func getJobCount(value string) (int, error) {
	re := regexp.MustCompile("[0-9]+")
	count := re.FindAllString(value, -1)

	return strconv.Atoi(count[0])
}
