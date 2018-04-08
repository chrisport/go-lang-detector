package benchme

import (
	"github.com/rcrowley/go-metrics"
	"io"
	"log"
	"net/http"
	"time"
)

func NewHttpBenchmark(name string, reqProvider func() *http.Request) HttpBenchmark {
	reg := metrics.NewPrefixedRegistry(name + "_")
	t := metrics.GetOrRegisterTimer("timer", reg)
	c := metrics.GetOrRegisterCounter("errorCount", reg)
	s := func(response *http.Response) bool {
		return response.StatusCode == 200
	}

	return &httpBenchmark{
		name:        name,
		timer:       t,
		errCount:    c,
		reqProvider: reqProvider,
		successFunc: s,
		registry:    reg,
	}
}

type HttpBenchmark interface {
	PrintStats(deltaTime int, out io.Writer)
	Warmup(N int)
	Benchmark(N int)
	Errs() []error
}

type httpBenchmark struct {
	name        string
	timer       metrics.Timer
	errCount    metrics.Counter
	reqProvider func() *http.Request
	registry    metrics.Registry
	successFunc func(response *http.Response) bool
	errors      []error
}

func (h *httpBenchmark) PrintStats(deltaTime int, out io.Writer) {
	go metrics.LogScaled(h.registry, 2*time.Second, time.Millisecond, log.New(out, h.name+"_metrics", log.Lmicroseconds))
}

func (h *httpBenchmark) Warmup(N int) {
	for i := 0; i < N; i++ {
		log.Printf("warmup %v/%v", i+1, N)
		h.executeRequest(false)
	}
}

func (h *httpBenchmark) Errs() []error {
	return h.errors
}
func (h *httpBenchmark) Benchmark(N int) {
	for i := 0; i < N; i++ {
		log.Printf("benchmark %v/%v", i+1, N)
		h.executeRequest(true)
	}
}

func (h *httpBenchmark) executeRequest(timed bool) {
	r := h.reqProvider()
	start := time.Now()
	res, err := http.DefaultClient.Do(r)
	diff := time.Since(start)

	if !timed {
		return
	}

	if err == nil && (h.successFunc == nil || h.successFunc(res)) {
		h.timer.Update(diff)
	} else {
		h.errors = append(h.errors, err)
		h.errCount.Count()
	}
}
