package cmd

import (
	"ale-chatGPT/cmd/gpt"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var gptCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gpt",
		Short: "get response from gpt",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			app, err := gpt.Initialize(ctx, ConfigPath)
			if err != nil {
				log.Fatal(err)
			}
			resp, getErr := app.Usecase.GetGPTResponse(ctx, Token)
			if getErr != nil {
				log.Fatal(getErr)
			}
			fmt.Println(resp)
		},
	}

	return cmd
}
