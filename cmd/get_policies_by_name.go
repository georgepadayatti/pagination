package cmd

import (
	"github.com/georgepadayatti/pagination/usecase"
	"github.com/spf13/cobra"
)

var getPoliciesByNameCmd = &cobra.Command{
	Use:   "get-policies-by-name",
	Short: "Get policies by name",
	Args:  cobra.ExactArgs(0),
	Run:   getPoliciesByNameCmdHandler,
}

func getPoliciesByNameCmdHandler(cmd *cobra.Command, args []string) {
	usecase.GetPoliciesByName()
}
