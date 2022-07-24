package crawler

import (
	"github.com/gocolly/colly/v2"
	"math/rand"
)

type Job struct {
	Keyword string
	Count   int
	Vacancy int
}

// Crawler interface is not using rn.
// Planning to use it when implement other crawler, like LinkedIn, Glassdoor etc
type Crawler interface {
	GetCount(keyword string) Job
}

// collector gets the colly collector for scrapping
func collector(options ...colly.CollectorOption) *colly.Collector {
	return colly.NewCollector(append(options, colly.UserAgent(getRandomUserAgent()))...)
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
