package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu   sync.Mutex
	urls map[string]bool
}

func (c *Cache) InsertUrls(urls []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, url := range urls {
		c.urls[url] = false
	}
}

func (c *Cache) UpdateFetchedUrl(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.urls[url] = true
}

func (c *Cache) IsUrlFetched(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.urls[url]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *Cache, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	// insert new urls in cache
	// and update url fetched
	cache.InsertUrls(urls)
	cache.UpdateFetchedUrl(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		// check cache if url is already fetched
		if !cache.IsUrlFetched(u) {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, cache, wg)
		}
	}
	return
}

// synced crawl using sync.Metux and sync.WaitGroup
func SyncCrawl(url string, depth int, fetcher Fetcher) {
	var wg sync.WaitGroup
	cache := Cache{urls: make(map[string]bool)}

	wg.Add(1)
	go Crawl(url, depth, fetcher, &cache, &wg)
	wg.Wait()
}

func main() {
	SyncCrawl("https://golang.org/", 4, fetcher)
	fmt.Println("done")
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
