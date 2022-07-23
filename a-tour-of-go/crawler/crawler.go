/*

Exercise: Web Crawler
In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!

*/

package crawler

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	go func() {
		if depth <= 0 {
			return
		}

		visitedUrlMap.Lock()

		if visited := visitedUrlMap.WasVisited(url); visited {
			visitedUrlMap.Unlock()
			return
		}

		body, urls, err := fetcher.Fetch(url)

		visitedUrlMap.SetAsVisited(url)

		visitedUrlMap.Unlock()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			Crawl(u, depth-1, fetcher)
		}
	}()
}

func Test() {
	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(1 * time.Second) // Sleeping so that the main function does not exit before every goroutine can finish
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// Ideally this would be a separate package

type VisitedUrlMap struct {
	mutex sync.Mutex
	v     map[string]bool
}

func (m *VisitedUrlMap) Lock() {
	m.mutex.Lock()
}

func (m *VisitedUrlMap) Unlock() {
	m.mutex.Unlock()
}

func (m *VisitedUrlMap) SetAsVisited(url string) {
	visitedUrlMap.v[url] = true
}

func (m *VisitedUrlMap) WasVisited(url string) bool {
	_, ok := visitedUrlMap.v[url]

	return ok
}

var visitedUrlMap = VisitedUrlMap{v: make(map[string]bool)}
