package crawler

import (
	"errors"
	"fmt"
	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/gocolly/colly/v2"
	"log"
	"net/url"
	"regexp"
	"strconv"
)

const baseDomain string = "jobs.bdjobs.com"

var domains = []string{
	baseDomain,
	"www." + baseDomain,
}

// GetCount crawl the webpage for the latest job count by a keyword
func GetCount(keyword string) (Job, error) {
	// colly.allowed domain can be refactored
	c := collector(colly.AllowedDomains(domains...))
	count := 0
	vacancies := 0

	c.OnHTML("#TopTotalRecord", func(e *colly.HTMLElement) {
		count = parseCount(e)
	})

	c.OnHTML("#TopTotalVacancy", func(e *colly.HTMLElement) {
		vacancies = parseCount(e)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		bugsnag.Notify(fmt.Errorf("request URL: %s failed with response %v \n Error: %v", r.Request.URL, r, err), r.Ctx)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(getUrl(keyword))

	if count == 0 {
		bugsnag.Notify(fmt.Errorf("failed to get count for %s", keyword))

		return Job{}, errors.New("Failed to get count for " + keyword)
	}

	return Job{Keyword: keyword, Count: count, Vacancy: vacancies}, nil
}

func parseCount(e *colly.HTMLElement) int {
	re := regexp.MustCompile(`\d+`)
	count, err := strconv.Atoi(re.FindAllString(e.Text, -1)[0])

	if err != nil {
		log.Fatal("parseCount: could not retrieve the count")
	}

	return count
}

func getUrl(keyword string) string {
	return "https://" + baseDomain + "/jobsearch.asp?fcatId=8&ica" +
		"tId=&JobKeyword=" + url.QueryEscape(keyword)
}
