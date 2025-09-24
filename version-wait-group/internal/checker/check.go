package checker

import (
	"net/http"
	"time"
)

type CheckResult struct {
	Target  string
	Statuts string
	Err     error
}

func CheckUrl(url string) CheckResult {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(url)

	if err != nil {
		return CheckResult{
			Target: url,
			Err:    &UnreachableUrlError{URL: url, Err: err},
		}
	}

	defer resp.Body.Close()

	return CheckResult{
		Target:  url,
		Statuts: resp.Status,
	}
}
