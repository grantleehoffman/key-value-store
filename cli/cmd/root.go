package cmd

import (
	"fmt"
	"os"

	"github.com/grantleehoffman/key-value-store/cli/action"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:              "key-value",
	Short:            "A simple cli to interact with a consul kv service.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
}

func init() {
	cobra.OnInitialize()

	RootCmd.AddCommand(getKey)
	RootCmd.PersistentFlags().StringP("ServiceEndpoint", "s", "kvdemo.thehoff.xyz", "(string) key value service host endpoint.")
	getKey.Flags().StringP("key", "k", "", "(string) The key to retrieve the value of.")
	getKey.MarkFlagRequired("key")

	RootCmd.AddCommand(putKey)
	putKey.Flags().StringP("key", "k", "", "(string) The key to add to the key value store.")
	putKey.Flags().StringP("value", "v", "", "(string) The value of the key to add to the key value store.")
	putKey.MarkFlagRequired("key")
	putKey.MarkFlagRequired("value")

	RootCmd.AddCommand(deleteKey)
	deleteKey.Flags().StringP("key", "k", "", "(string) The key to delete from the key value store.")
	deleteKey.MarkFlagRequired("key")
}

var getKey = &cobra.Command{
	Use:   "get --key <key> ",
	Short: "Get the value of a stored key.",
	Run: func(cmd *cobra.Command, args []string) {
		key := cmd.Flag("key").Value.String()
		err := action.GetKey(key)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	},
}

var putKey = &cobra.Command{
	Use:   "put --key <key> --value <value>",
	Short: "Create a new stored key/value.",
	Run: func(cmd *cobra.Command, args []string) {
		key := cmd.Flag("key").Value.String()
		value := cmd.Flag("value").Value.String()

		err := action.PutKey(key, value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	},
}

var deleteKey = &cobra.Command{
	Use:   "delete --key <key>",
	Short: "Delete a key from the key value store.",
	Run: func(cmd *cobra.Command, args []string) {
		key := cmd.Flag("key").Value.String()

		err := action.DeleteKey(key)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	},
}
