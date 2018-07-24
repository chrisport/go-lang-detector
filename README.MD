# Playground

A collection of small things I played around with,
that are not worth an own project.


## Shut - Simple Sleep Timer
Usage:

$ ./sleep.sh   
Sleeping in 6 minutes

$ ./sleep.sh 12   
Sleeping in 12 minutes

## internetmonitor

install using `go install github.com/chrisport/go-playground/internetmonitor`
run `internetmonitor`

Sample console output:
```
time            status  last    total
14:32:46        online  0s      0s
14:32:51        online  0s      0s
14:32:57        offline 5s      5s
14:33:02        offline 10s     10s
14:33:07        online  0s      10s
14:33:12        online  0s      10s
14:33:17        offline 5s      15s
14:33:22        offline 10s     20s
14:33:27        online  0s      20s
```


## Benchme - HTTP Request benchmark
Executes a Request N times for measuring response time.
Uses github.com/rcrowley/go-metrics to collect and present results.
Sample usage:

```
func main(){
	b := benchme.NewHttpBenchmark("google", createRequest)
	// warmup
	b.Warmup(5)

	// start printing result to stdout
	b.PrintStats(1, os.Stdout)
	// benchmark 10 times
	b.Benchmark(10)

    // printing occured errors to console
	errs := b.Errs()
	for _, err := range errs {
		log.Println(err)
	}
	time.Sleep(2*time.Second)
}

func createRequest() *http.Request {
	testURL := "https://www.google.com"
	req, _ := http.NewRequest("GET", testURL, nil)
	req.Header.Set("Authorization", "Bearer whatevertoken")
	req.Header.Set("Cache-Control", "public")
	return req
}
```
