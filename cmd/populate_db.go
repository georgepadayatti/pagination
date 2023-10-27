package cmd

import (
	"github.com/georgepadayatti/pagination/usecase"
	"github.com/spf13/cobra"
)

var populateDbCmd = &cobra.Command{
	Use:   "populate-db",
	Short: "Populate DB with dummy records",
	Args:  cobra.ExactArgs(0),
	Run:   populateDbCmdHandler,
}

func populateDbCmdHandler(cmd *cobra.Command, args []string) {
	usecase.CreateDataAgreement()
}
