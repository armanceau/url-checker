package cmd

import (
	"errors"
	"fmt"
	"sync"

	"github.com/armanceau/url-checker/version-cobra/internal/checker"
	"github.com/armanceau/url-checker/version-cobra/internal/config"
	"github.com/armanceau/url-checker/version-cobra/internal/reporter"
	"github.com/spf13/cobra"
)

var (
	inputFilePath  string
	outputFilePath string
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Vérifier l'accessibilité d'un liste d'Urls",
	Long:  "Vérifier l'accessibilité d'un liste d'Urls en plus long",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFilePath == "" {
			fmt.Println("Erreur : le chemin du fichier d'entrée est vide")
			return
		}
		targets, err := config.LoadTargetFromFile(inputFilePath)

		if err != nil {
			fmt.Printf("Erreur lors du chargement des Urls: %v\n", err)
			return
		}

		if len(targets) == 0 {
			fmt.Println("Aucune Url à vérifier dans le fichier d'entrée")
			return
		}

		resultsChan := make(chan checker.CheckResult, len(targets))
		var wg sync.WaitGroup

		wg.Add(len(targets))
		for _, target := range targets {
			go func(t config.InputTarget) {
				defer wg.Done()
				result := checker.CheckUrl(t)
				resultsChan <- result

			}(target)
		}
		wg.Wait()
		close(resultsChan)

		var finalReport []checker.ReportEntry
		for res := range resultsChan {
			reportEntry := checker.ConvertToReportEntry(res)
			finalReport = append(finalReport, reportEntry)
			if res.Err != nil {
				var unreachable *checker.UnreachableUrlError
				if errors.As(res.Err, &unreachable) {
					fmt.Printf("KO %v: %v\n", unreachable.URL, unreachable.Err)
				} else {
					fmt.Printf("KO %v: %v\n", res.InputTarget, res.Err)
				}
			} else {
				fmt.Printf("OK %v: %v\n", res.InputTarget, res.Statuts)
			}
			if outputFilePath != "" {
				err := reporter.ExportResultToJsonFile(outputFilePath, finalReport)
				if err != nil {
					fmt.Printf("Erreur lors de l'exportation des résultats")
				} else {
					fmt.Printf("Résultats exportés")
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVar(&inputFilePath, "input", "i", "")
	checkCmd.Flags().StringVar(&outputFilePath, "output", "o", "")
}
