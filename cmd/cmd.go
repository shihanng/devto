package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/shihanng/devto/pkg/devto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagAPIKey = "api-key"
)

func New() *cobra.Command {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("DEVTO")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List published articles on dev.to",
		RunE:  listRunE,
	}

	rootCmd := &cobra.Command{
		Use:               "devto",
		Short:             "Publish to dev.to from your terminal",
		PersistentPreRunE: rootRunE,
	}

	rootCmd.PersistentFlags().String(flagAPIKey, "", "API key for authentication")
	rootCmd.AddCommand(listCmd)

	viper.BindPFlag(flagAPIKey, rootCmd.PersistentFlags().Lookup(flagAPIKey))

	return rootCmd
}

func rootRunE(cmd *cobra.Command, args []string) error {
	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return errors.Wrap(err, "cmd: read config")
		}
	}

	config := struct {
		APIKey string `mapstructure:"api-key"`
	}{}

	if err := viper.Unmarshal(&config); err != nil {
		return errors.Wrap(err, "cmd: unmarshal config")
	}

	return nil
}

func listRunE(cmd *cobra.Command, args []string) error {
	apiKey := context.WithValue(context.Background(), devto.ContextAPIKey, devto.APIKey{
		Key: viper.GetString(flagAPIKey),
	})

	client := devto.NewAPIClient(devto.NewConfiguration())
	articles, _, err := client.ArticlesApi.GetUserAllArticles(apiKey, nil)
	if err != nil {
		return err
	}

	for _, a := range articles {
		fmt.Println(a.Title, a.Id)
	}

	return nil
}
