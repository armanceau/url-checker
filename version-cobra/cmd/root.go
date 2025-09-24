package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gowatcher",
	Short: "gowatcher est un outil pour vérifier l'accessibilité des Urls.",
	Long:  "gowatcher est un outil pour vérifier l'accessibilité des Urls. Une description en plus long ...",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur %v\n", err)
		os.Exit(1)
	}
}

func init() {}
