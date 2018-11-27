package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/fast", FastHandler)
	http.HandleFunc("/slow", SlowHandler)
	http.HandleFunc("/crawl", CrawlerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	w.WriteHeader(http.StatusOK)
}

func FastHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Millisecond)
	w.WriteHeader(http.StatusOK)
}

func CrawlerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	crawlResponse := make(chan string)
	go func() {
		url, error := Crawl("http://localhost:8080/slow")
		if error != nil {
			cancel()
		} else {
			crawlResponse <- url
			close(crawlResponse)
		}
	}()
	ProcessResponse(ctx, crawlResponse, w)
}

func ProcessResponse(ctx context.Context, crawlResponse <-chan string, w http.ResponseWriter) {
	select {
	case data := <-crawlResponse:
		fmt.Println(data)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	case <-ctx.Done():
		fmt.Println("X.x")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Crawl(url string) (string, error) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req = req.WithContext(ctx)
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return "", err
	}
	return url, nil
}
