package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "version of the app")
}

var rootCmd = &cobra.Command{
	Use:   "klient",
	Short: "Client to make http request and send tcp messages",
	Long:  `Client to make http request and send tcp messages: see more on the corresponding commands..`,
	Run: func(cmd *cobra.Command, args []string) {
		if v, err := cmd.Flags().GetBool("versi"); err != nil || v {
			fmt.Println("klient version: 0.0.1")
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
