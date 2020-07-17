package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete company from internship search",
	Run: func(cmd *cobra.Command, args []string) {
		delete(args)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete(args []string) {
	company := args[0]

	viper.SetDefault(fmt.Sprintf("companies.%s", company), nil)
	viper.WriteConfig()
}
