package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var ConfigPath string
var Token string

var rootCmd = &cobra.Command{}

func init() {

	rootCmd.AddCommand(gptCmd())
	rootCmd.PersistentFlags().StringVarP(&ConfigPath, "config", "c", "./environment/local/config.json", "config file")
	rootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "Give me a new perspective", "type question")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
