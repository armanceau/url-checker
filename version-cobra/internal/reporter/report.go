package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/armanceau/url-checker/version-cobra/internal/checker"
)

func ExportResultToJsonFile(filePath string, resultats []checker.ReportEntry) error {
	data, err := json.MarshalIndent(resultats, "", " ")
	if err != nil {
		return fmt.Errorf("Impossible d'encoder les résultats en json")
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("Impossible d'écrire")
	}
	return nil
}
