package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
	"strconv"
)

var (
	url        string
	endpoint   string
	port       int
	bodySrc    string
	validVerbs = []string{"GET", "POST"}
)

func init() {
	httpCmd.PersistentFlags().StringVarP(&url, "url", "u", "http://localhost", "url to hit")
	httpCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	httpCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "", "endpoint (optional)")
	postCmd.Flags().StringVarP(&bodySrc, "body", "b", "", "path to json file for body (required)")
	postCmd.MarkFlagRequired("body")
	httpCmd.AddCommand(getCmd)
	httpCmd.AddCommand(postCmd)
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Http client to send request",
	Args:  validVerb,
	Run: func(cmd *cobra.Command, args []string) {
		url := url + ":" + strconv.Itoa(port)
		fmt.Println("http called to:", url)
	},
}

var getCmd = &cobra.Command{
	Use:   "GET",
	Short: "sending a get request",
	Run: func(cmd *cobra.Command, args []string) {
		url := url + ":" + strconv.Itoa(port)
		fmt.Println("just sending a get:", url)
	},
}

var postCmd = &cobra.Command{
	Use:   "POST",
	Short: "sending a post request",
	Run: func(cmd *cobra.Command, args []string) {
		url := url + ":" + strconv.Itoa(port)
		fmt.Println("just sending a post:", url)
	},
}

func validVerb(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New(fmt.Sprintf("this action requires a verb argument:%s", validVerbs))
	}
	if slices.Contains(validVerbs, args[0]) {
		return nil
	}
	return fmt.Errorf("invalid verb specified: %s, supported: %s", args[0], validVerbs)
}
