package main

// This is a nice challenge to demonstrate different ways to fetch multiple urls concurrently
// and display on the terminal the response status code with the associated url.
// We could improve the fetch result by adding response body size and time taken to fetch.
// start := time.Now() then n, err := io.Copy(ioutil.Discard, resp.Body)
// secs := time.Since(start).Seconds() and fmt.Sprintf("%s :: %.2fs :: %7d", resp.Status, secs, n)

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 26 August 2021

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

// Fetch the link and print out the URL and HTTP response code for each link concurrently
func Fetch(link string) string {

	// set the http connection timeout.
	client := http.Client{Timeout: 5 * time.Second}

	// get the full content.
	resp, err := client.Get(link)
	if err != nil {
		// failed to fetch.
		return "n/a"
	}
	defer resp.Body.Close()

	return resp.Status
}

// FirstWorker demonstrates an approach to use Fetch concurrently.
// Each url is fetched by a unique goroutine. And the status is
// displayed by that same goroutine. Here we inject the url link
// directly into the goroutine as global variable.
func FirstWorker(links []string) {
	wg := &sync.WaitGroup{}
	for _, link := range links {
		wg.Add(1)
		// create a new variable.
		link := link
		go func() {
			defer wg.Done()
			status := Fetch(link)
			fmt.Printf("%s : %s\n", link, status)
		}()
	}
	wg.Wait()
}

// SecondWorker demonstrates an approach to use Fetch concurrently.
// Each url is fetched by a unique goroutine. And the status is
// displayed by that same goroutine. Here we inject the url link
// as input parameter to the goroutine function.
func SecondWorker(links []string) {
	wg := &sync.WaitGroup{}
	for _, link := range links {
		wg.Add(1)
		// pass variable to anonymous func.
		go func(url string) {
			defer wg.Done()
			status := Fetch(url)
			fmt.Printf("%s : %s\n", url, status)
		}(link)
	}
	wg.Wait()
}

// ThirdWorker demonstrates an approach to use Fetch concurrently.
// Each goroutine puts its result on a channel. Another goroutine
// monitor that channel and retrieve the result for displaying.
func ThirdWorker(links []string) {
	numberOfLinks := len(links)
	resultsChannel := make(chan string)
	done := make(chan bool)

	go func() {
		for i := 0; i < numberOfLinks; i++ {
			result := <-resultsChannel
			fmt.Print(result)
		}
		done <- true
	}()

	for _, link := range links {
		// pass variable to anonymous func.
		go func(url string) {
			status := Fetch(url)
			resultsChannel <- fmt.Sprintf("%s : %s\n", url, status)
		}(link)
	}
	// block until all results displayed.
	<-done
}

// FourthWorker demonstrates an approach to use Fetch concurrently.
// A pool (with a predefined number) of worker handles all the urls.
// Each worker fetches the url and builds the result then send it on
// a results channel. A deterministic loop read & displays each result.
func FourthWorker(links []string) {
	n := len(links)
	// lets use the number of Cores.
	numberOfWorkers := runtime.NumCPU()
	if n < numberOfWorkers {
		numberOfWorkers = n
	}
	// buffered channels with number of links.
	jobsChannel := make(chan string, n)
	resultsChannel := make(chan string, n)

	// spin up all workers.
	for i := 0; i < numberOfWorkers; i++ {
		go func(id int) {
			for url := range jobsChannel {
				status := Fetch(url)
				// build result and add to results channel for displaying.
				resultsChannel <- fmt.Sprintf("worker %d :: %s : %s\n", id, url, status)
			}
		}(i)
	}
	// asynchronously feed the jobs channel.
	go func() {
		for _, url := range links {
			jobsChannel <- url
		}
	}()

	// ensure to read n number of results.
	for r := 0; r < n; r++ {
		fmt.Print(<-resultsChannel)
	}

	// close to signal workers to terminate.
	close(jobsChannel)
}

func main() {

	// list of urls for testing.
	links := []string{

		"https://cisco.com",
		"https://google.com",
		"https://facebook.com",
		"https://microsoft.com",
		"https://amazon.com",
		"https://twitter.com",
	}

	fmt.Println()

	// launch the 1st technique.
	FirstWorker(links)

	fmt.Println()

	// launch the 2nd technique.
	SecondWorker(links)

	fmt.Println()

	// launch the 3rd technique.
	ThirdWorker(links)

	fmt.Println()

	// launch the 4th technique.
	FourthWorker(links)
}
