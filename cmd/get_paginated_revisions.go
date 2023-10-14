package cmd

import (
	"github.com/georgepadayatti/pagination/usecase"
	"github.com/spf13/cobra"
)

var getPaginatedRevisionsCmd = &cobra.Command{
	Use:   "get-paginated-revisions",
	Short: "Get paginated revisions",
	Args:  cobra.ExactArgs(0),
	Run:   getPaginatedRevisionsCmdHandler,
}

func getPaginatedRevisionsCmdHandler(cmd *cobra.Command, args []string) {
	usecase.GetPaginatedRevisions()
}
