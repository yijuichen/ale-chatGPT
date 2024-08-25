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
	rootCmd.PersistentFlags().StringVar(&ConfigPath, "config", "./environment/local/config.json", "config file")
	rootCmd.PersistentFlags().StringVar(&Token, "token", "Give me a new perspective", "type question")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
