package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pagination",
	Short: "Example app to demonstrate pagination in MongoDB",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(populateDbCmd)
	rootCmd.AddCommand(getPaginatedPoliciesCmd)
	rootCmd.AddCommand(getPaginatedRevisionsCmd)
	rootCmd.AddCommand(getPoliciesByNameCmd)
}
