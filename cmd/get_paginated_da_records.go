package cmd

import (
	"github.com/georgepadayatti/pagination/usecase"
	"github.com/spf13/cobra"
)

var getPaginatedDaRecordsCmd = &cobra.Command{
	Use:   "get-paginated-da-records",
	Short: "Get pagianted da records",
	Args:  cobra.ExactArgs(0),
	Run:   getPaginatedDaRecordsCmdHandler,
}

func getPaginatedDaRecordsCmdHandler(cmd *cobra.Command, args []string) {
	usecase.GetDataAgreementRecords()
}
