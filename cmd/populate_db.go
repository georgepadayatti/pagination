package cmd

import (
	"github.com/georgepadayatti/pagination/usecase"
	"github.com/spf13/cobra"
)

var populateDbCmd = &cobra.Command{
	Use:   "populate-db",
	Short: "Populate DB with 10 policies and policy authors",
	Args:  cobra.ExactArgs(0),
	Run:   populateDbCmdHandler,
}

func populateDbCmdHandler(cmd *cobra.Command, args []string) {
	usecase.CreateTenPolicyDocuments()
}
