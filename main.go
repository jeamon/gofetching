package main

// This is a nice challenge to demonstrate different ways to fetch multiple urls concurrently.
// Version  : 1.0
// Author   : Jerome AMON
// Created  : 26 August 2021

import (
	"fmt"
	"net/http"
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
func FirstWorker(links *[]string) {
	wg := &sync.WaitGroup{}
	for _, link := range *links {
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
func SecondWorker(links *[]string) {
	wg := &sync.WaitGroup{}
	for _, link := range *links {
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
func ThirdWorker(links *[]string) {
	numberOfLinks := len(*links)
	resultsChannel := make(chan string)
	done := make(chan bool)

	go func() {
		for i := 0; i < numberOfLinks; i++ {
			result := <-resultsChannel
			fmt.Print(result)
		}
		done <- true
	}()

	for _, link := range *links {
		// pass variable to anonymous func.
		go func(url string) {
			status := Fetch(url)
			resultsChannel <- fmt.Sprintf("%s : %s\n", url, status)
		}(link)
	}
	// block until all results displayed.
	<-done
}

func main() {

	// list of urls for testing.
	links := []string{"https://cisco.com", "https://google.com", "https://facebook.com", "https://microsoft.com", "https://amazon.com", "https://twitter.com"}

	// launch the 1st technique.
	FirstWorker(&links)

	fmt.Println()

	// launch the 2nd technique.
	SecondWorker(&links)

	fmt.Println()

	// launch the 3rd technique.
	ThirdWorker(&links)
}

/*
// Output:

[22:26:49] {nxos-geek}:~$ go run main.go
https://twitter.com : 200 OK
https://google.com : 200 OK
https://microsoft.com : 200 OK
https://cisco.com : 200 OK
https://amazon.com : 200 OK
https://facebook.com : 200 OK

https://twitter.com : 200 OK
https://google.com : 200 OK
https://facebook.com : 200 OK
https://amazon.com : 200 OK
https://microsoft.com : 200 OK
https://cisco.com : 200 OK

https://twitter.com : 200 OK
https://google.com : 200 OK
https://facebook.com : 200 OK
https://microsoft.com : 200 OK
https://amazon.com : 200 OK
https://cisco.com : 200 OK

[22:27:10] {nxos-geek}:~$

*/
