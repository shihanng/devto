package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shihanng/devto/pkg/devto"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "devto",
		Short: "Publish to dev.to from your terminal",
		Run: func(cmd *cobra.Command, args []string) {

			auth := context.WithValue(context.Background(), devto.ContextAPIKey, devto.APIKey{
				Key: os.Getenv("DEV_TOKEN"),
			})

			client := devto.NewAPIClient(devto.NewConfiguration())
			articles, _, err := client.ArticlesApi.GetUserAllArticles(auth, nil)
			if err != nil {
				log.Fatal(err)
			}

			for _, a := range articles {
				fmt.Println(a.Title, a.Id)
			}
		},
	}
	return rootCmd
}
