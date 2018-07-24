package main

import (
	"net/http"
	"fmt"
	"time"
	"github.com/bugsnag/bugsnag-go/errors"
)

var timeFormat = "15:04:05"

func main() {
	lastOnline := time.Now()
	fmt.Printf("time\t\tstatus\tlast\ttotal\n")
	totalOfflineTime := time.Second * 0
	for {
		time.Sleep(5 * time.Second)

		t := time.Now()
		err := testConnection()
		stat := "offline"
		if err == nil {
			stat = "online"
			lastOnline = t
		} else {
			totalOfflineTime = totalOfflineTime + 5*time.Second
		}
		fmt.Printf("%v\t%v\t%vs\t%v\n", t.Format(timeFormat), stat, t.Sub(lastOnline).Truncate(time.Second).Seconds(), totalOfflineTime)
	}
}

func testConnection() error {
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		return err
	}
	cli := http.Client{Timeout: 4 * time.Second}
	res, err := cli.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.Errorf("not 200 %v", res)
	}
	return nil
}
