package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type CacheMap struct {
	visited map[string]bool
	mux sync.Mutex
}

func Crawl(url string, depth int, fetcher Fetcher, ch chan response, cacheMap CacheMap) {
	defer close(ch)
	if depth <= 0 {
		return
	}
	cacheMap.mux.Lock()
	if cacheMap.visited[url] {
		cacheMap.mux.Unlock()
		return
	} else {
		cacheMap.visited[url] = true
		cacheMap.mux.Unlock()
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- response{url, body}
	results := make([]chan response, len(urls))
	for i, u := range urls {
		results[i] = make(chan response)
		go Crawl(u, depth - 1, fetcher, results[i], cacheMap)
	}
	for i := range results {
		for r := range results[i] {
			ch <- r
		}
	}
	return
}

func main() {
	ch := make(chan response)
	go Crawl("http://golang.org/", 4, fetcher, ch, CacheMap{visited: make(map[string]bool)})
	for r := range ch {
		fmt.Printf("found: %s %q\n", r.url, r.body)
	}
}

type response struct {
	url, body string
}

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

var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
