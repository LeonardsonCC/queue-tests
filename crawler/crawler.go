package crawler

import (
	"net/http"
	"sync"
	"sync/atomic"
)

// how the crawler is
type Crawler struct {
	wg        sync.WaitGroup
	errsCount int64
	instances int
	url       string
}

// instantiate a new crawler with the url to get and how many instances
func NewCrawler(url string, instances int) *Crawler {
	return &Crawler{
		wg:        sync.WaitGroup{},
		instances: instances,
		url:       url,
	}
}

// start the crawler with instances as how many go routines it'll start
func (c *Crawler) Run(wg *sync.WaitGroup) {
	for i := 0; i < c.instances; i++ {
		c.wg.Add(1)
		go c.MakeRequest()
	}

	c.wg.Wait()

	wg.Done()
}

// makes the request to the url of the struct and add one to the error count if it needs
func (c *Crawler) MakeRequest() {
	_, err := http.Get(c.url)

	if err != nil {
		atomic.AddInt64(&c.errsCount, 1)
	}

	c.wg.Done()
}
