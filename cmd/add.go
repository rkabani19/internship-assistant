package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add company and url to search",
	Run: func(cmd *cobra.Command, args []string) {
		add(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(args []string) {
	company := args[0]
	url := args[1]

	viper.SetDefault(fmt.Sprintf("companies.%s", company), url)
	viper.WriteConfig()
}
