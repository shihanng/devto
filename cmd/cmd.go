package cmd

import (
	"context"
	"fmt"

	"github.com/shihanng/devto/pkg/devto"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "devto",
		Short: "Publish to dev.to from your terminal",
	}
	rootCmd.PersistentFlags().String("api-key", "", "API key for authentication")

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List published articles on dev.to",
		RunE:  listRunE,
	}

	rootCmd.AddCommand(listCmd)

	return rootCmd
}

func listRunE(cmd *cobra.Command, args []string) error {
	apiKey, err := cmd.Parent().PersistentFlags().GetString("api-key")
	if err != nil {
		return err
	}

	auth := context.WithValue(context.Background(), devto.ContextAPIKey, devto.APIKey{
		Key: apiKey,
	})

	client := devto.NewAPIClient(devto.NewConfiguration())
	articles, _, err := client.ArticlesApi.GetUserAllArticles(auth, nil)
	if err != nil {
		return err
	}

	for _, a := range articles {
		fmt.Println(a.Title, a.Id)
	}

	return nil
}
