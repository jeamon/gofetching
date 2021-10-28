package main

// run the unit tests : go test main.go main_test.go -v
// stop at first failing test : go test -v -failfast
// test specific unit : go test -run TestFetch_Valid -v
// test specific unit : go test -run TestFetch_EndToEnd -v
// test specific unit : go test -run ExampleFirstWorker -v
// view the test coverage : go test -cover

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestFetch unit tests Fetch function on returned status accuracy.
func TestFetch_Valid(t *testing.T) {
	expected := "200 OK"
	// mock a server which only sends status 200.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	status := Fetch(server.URL)
	if status != expected {
		t.Errorf("expected status to be %s but got %s", expected, status)
	}
}

// TestFetch unit tests Fetch function on returned status accuracy.
func TestFetch_Invalid(t *testing.T) {
	notExpected := "200 OK"
	// mock a server which only sends status <404 Not Found>.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	status := Fetch(server.URL)
	if status == notExpected {
		t.Errorf("status should not be %s but instead %s", notExpected, status)
	}
}

// TestFetch_Timeout unit tests Fetch function on timeout behaviour.
func TestFetch_Timeout(t *testing.T) {
	expected := "n/a"
	// mock a server which only sends status 200.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// delay the response for more than 5secs.
		time.Sleep(6 * time.Second)
		w.WriteHeader(http.StatusFound)
	}))
	defer server.Close()

	status := Fetch(server.URL)

	if status != expected {
		t.Errorf("expected status to be %s but got %s", expected, status)
	}
}

// TestFetch_EndToEnd performs a parallel table-driven end-to-end testing of Fetch function.
func TestFetch_EndToEnd(t *testing.T) {
	// slice of anonymous struct for functional end-to-end testings.
	var testsTable = []struct {
		name string
		url  string
		want string
	}{
		{name: "testing status 200", url: "https://httpstat.us/200", want: "200 OK"},
		{name: "testing status 400", url: "https://httpstat.us/400", want: "400 Bad Request"},
		{name: "testing status 401", url: "https://httpstat.us/401", want: "401 Unauthorized"},
		{name: "testing status 404", url: "https://httpstat.us/404", want: "404 Not Found"},
		{name: "testing status 503", url: "https://httpstat.us/503", want: "503 Service Unavailable"},
		{name: "testing status n/a", url: "https://hxxxxxxx.us/", want: "n/a"},
	}

	for _, tc := range testsTable {
		t.Run(tc.name, func(t *testing.T) {
			if got := Fetch(tc.url); got != tc.want {
				t.Errorf("expected: %v, got %v", tc.want, got)
			}
		})
	}
}

// ExampleFirstWorker performs end-to-end functional testing of FirstWorker function.
// Make sure that the website used here (https://cisco.com) is available & accessible.
func ExampleFirstWorker() {
	FirstWorker([]string{"https://cisco.com"})
	// Output:
	// https://cisco.com : 200 OK
}

// ExampleSecondWorker performs end-to-end functional testing of SecondWorker function.
// Make sure that the website used here (https://cisco.com) is available & accessible.
func ExampleSecondWorker() {
	SecondWorker([]string{"https://cisco.com"})
	// Output:
	// https://cisco.com : 200 OK
}

// ExampleThirdWorker performs end-to-end functional testing of ThirdWorker function.
// Make sure that the website used here (https://cisco.com) is available & accessible.
func ExampleThirdWorker() {
	ThirdWorker([]string{"https://cisco.com"})
	// Output:
	// https://cisco.com : 200 OK
}
