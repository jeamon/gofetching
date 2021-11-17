# go-concurrency-urls-fetching

A go-based repository to demonstrate multiple techniques to concurrently fetch multiple url over http.

## Outputs


* **Below is the program execution output**

```

~$ go run main.go

https://twitter.com : 200 OK
https://google.com : 200 OK
https://microsoft.com : 200 OK
https://cisco.com : 200 OK
https://amazon.com : 200 OK
https://facebook.com : 200 OK

https://google.com : 200 OK
https://twitter.com : 200 OK
https://microsoft.com : 200 OK
https://facebook.com : 200 OK
https://amazon.com : 200 OK
https://cisco.com : 200 OK

https://twitter.com : 200 OK
https://google.com : 200 OK
https://amazon.com : 200 OK
https://facebook.com : 200 OK
https://microsoft.com : 200 OK
https://cisco.com : 200 OK

worker 1 :: https://google.com : 200 OK
worker 0 :: https://microsoft.com : 200 OK
worker 3 :: https://facebook.com : 200 OK
worker 1 :: https://amazon.com : 200 OK
worker 0 :: https://twitter.com : 200 OK
worker 2 :: https://cisco.com : 200 OK

```


* **Below is the tests execution output**

```

~$ go test -v -failfast
=== RUN   TestFetch_Valid
--- PASS: TestFetch_Valid (0.00s)
=== RUN   TestFetch_Invalid
--- PASS: TestFetch_Invalid (0.00s)
=== RUN   TestFetch_Timeout
--- PASS: TestFetch_Timeout (6.00s)
=== RUN   TestFetch_EndToEnd
=== RUN   TestFetch_EndToEnd/testing_status_200
=== RUN   TestFetch_EndToEnd/testing_status_400
=== RUN   TestFetch_EndToEnd/testing_status_401
=== RUN   TestFetch_EndToEnd/testing_status_404
=== RUN   TestFetch_EndToEnd/testing_status_503
=== RUN   TestFetch_EndToEnd/testing_status_n/a
--- PASS: TestFetch_EndToEnd (2.98s)
    --- PASS: TestFetch_EndToEnd/testing_status_200 (0.83s)
    --- PASS: TestFetch_EndToEnd/testing_status_400 (0.35s)
    --- PASS: TestFetch_EndToEnd/testing_status_401 (0.31s)
    --- PASS: TestFetch_EndToEnd/testing_status_404 (0.33s)
    --- PASS: TestFetch_EndToEnd/testing_status_503 (0.33s)
    --- PASS: TestFetch_EndToEnd/testing_status_n/a (0.84s)
=== RUN   ExampleFirstWorker
--- PASS: ExampleFirstWorker (1.43s)
=== RUN   ExampleSecondWorker
--- PASS: ExampleSecondWorker (1.06s)
=== RUN   ExampleThirdWorker
--- PASS: ExampleThirdWorker (1.10s)
PASS
ok      github.com/jeamon/go-concurrency-urls-fetching  12.725s

```


## License

Please check & read [the license details](https://github.com/jeamon/go-concurrency-urls-fetching/blob/master/LICENSE) 


## Contact

Feel free to [reach out to me](https://blog.cloudmentor-scale.com/contact) before any action.