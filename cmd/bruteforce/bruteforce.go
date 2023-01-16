/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package bruteforce

import (
	"github.com/spf13/cobra"
)

// bruteforceCmd represents the bruteforce command
var BruteforceCmd = &cobra.Command{
	Use:   "bruteforce",
	Short: "暴力破解模块选择",
	Long:  `暴力破解模块，支持 ssh`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// rootCmd.AddCommand(bruteforceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bruteforceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bruteforceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
