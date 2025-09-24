package checker

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/armanceau/url-checker/version-cobra/internal/config"
)

type ReportEntry struct {
	Name    string
	URL     string
	Owner   string
	Statuts string
	ErrMsg  string
}

type CheckResult struct {
	InputTarget config.InputTarget
	Statuts     string
	Err         error
}

func CheckUrl(target config.InputTarget) CheckResult {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get(target.URL)

	if err != nil {
		return CheckResult{
			InputTarget: target,
			Err:         &UnreachableUrlError{URL: target.URL, Err: err},
		}
	}

	defer resp.Body.Close()

	return CheckResult{
		InputTarget: target,
		Statuts:     resp.Status,
	}
}

func ConvertToReportEntry(res CheckResult) ReportEntry {
	report := ReportEntry{
		Name:  res.InputTarget.Name,
		URL:   res.InputTarget.URL,
		Owner: res.InputTarget.Owner,
	}
	if res.Err != nil {
		var unreachable *UnreachableUrlError
		if errors.As(res.Err, &unreachable) {
			report.Statuts = "Inaccessible"
			report.ErrMsg = fmt.Sprintf("Unreachable URL: %v", unreachable)
		} else {
			report.Statuts = "Error"
			report.ErrMsg = fmt.Sprintf("Erreur générique:")
		}
	}
	return report
}
