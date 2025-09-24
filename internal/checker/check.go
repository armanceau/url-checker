package checker

import (
	"fmt"
	"net/http"
	"time"
)

type CheckResult struct {
	Target  string
	Statuts string
	Err     error
}

func CheckUrl(url string, results chan<- CheckResult) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(url)

	if err != nil {
		results <- CheckResult{
			Target: url, Err: fmt.Errorf("Request failed: %w", err),
		}
		return
	}

	defer resp.Body.Close()

	results <- CheckResult{
		Target:  url,
		Statuts: resp.Status,
	}
}
