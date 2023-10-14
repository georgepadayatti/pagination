package cmd

import (
	"github.com/georgepadayatti/pagination/usecase"
	"github.com/spf13/cobra"
)

var getPaginatedPoliciesCmd = &cobra.Command{
	Use:   "get-paginated-policies",
	Short: "Get paginated policies",
	Args:  cobra.ExactArgs(0),
	Run:   getPaginatedPoliciesCmdHandler,
}

func getPaginatedPoliciesCmdHandler(cmd *cobra.Command, args []string) {
	usecase.GetPaginatedPolicies()
}
