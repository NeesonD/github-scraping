package main

import (
	"fmt"
	"strings"
	"sync"
)

var (
	starUrls = []string{
		"https://github.com/NeesonD?tab=stars",
	}
)

func main() {
	var wg sync.WaitGroup
	for _, url := range starUrls {
		url := url
		wg.Add(1)
		go func() {
			defer wg.Done()
			Scraping(url, getCsvFileName(url))
		}()
	}
	wg.Wait()
}

func getCsvFileName(url string) string {
	url = strings.ReplaceAll(url, "?", "/")
	urlParts := strings.Split(url, "/")
	// 取出倒数第二个元素
	username := urlParts[len(urlParts)-2]
	fmt.Println(username)
	return fmt.Sprintf("%s_star_repo.csv", username)
}
