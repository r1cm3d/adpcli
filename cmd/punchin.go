package cmd

import (
	"fmt"
	"github.com/mniak/adpexpert"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var punchinCmd = &cobra.Command{
	Use:   "punchin",
	Short: "Punch the clock",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			os.Exit(1)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			os.Exit(1)
		}
		cli := adpexpert.Client{}

		log.Print(fmt.Sprintf("Trying to log into Username (%s) account.", username))
		if err := cli.Login(username, password); err != nil {
			log.Fatalf("Cannot login into Username (%s) account.", username)
		}
		log.Print(fmt.Sprintf("Logged succesfully into Username (%s) account.", username))
		log.Print(fmt.Sprintf("Trying to Punchin into Username (%s) account.", username))
		if err := cli.PunchIn(); err != nil {
			log.Fatalf("Cannot Punchin into Username (%s) account.", username)
		}
		log.Print(fmt.Sprintf("Punchin succesfully done into Username (%s) account.", username))

		os.Exit(0)
	},
}
