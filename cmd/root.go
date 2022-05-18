package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "adpcli",
	Long: `
           _            _ _ 
  __ _  __| |_ __   ___| (_)
 / _ |/ _ | _ |  \ / __| | |
| (_| | (_| | |_) | (__| | |
 \__,_|\__,_| .__/ \___|_|_|
            |_|             
Interacts with ADP Expert application via command line.
`,
}

func init() {
	rootCmd.PersistentFlags().StringP("username", "u", "", "The username of ADP Expert application.")
	rootCmd.PersistentFlags().StringP("password", "p", "", "The password of ADP Expert application.")
	rootCmd.AddCommand(punchinCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
