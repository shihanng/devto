package cmd

import (
	"context"
	"strings"

	"github.com/antihax/optional"
	"github.com/cockroachdb/errors"
	"github.com/shihanng/devto/pkg/article"
	"github.com/shihanng/devto/pkg/devto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	flagAPIKey = "api-key"
	flagDebug  = "debug"
)

func New() (*cobra.Command, func()) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("DEVTO")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	r := &runner{}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List published articles on dev.to",
		RunE:  r.listRunE,
	}

	submitCmd := &cobra.Command{
		Use:   "submit <Markdown file>",
		Short: "Submit article to dev.to",
		RunE:  r.submitRunE,
		Args:  cobra.ExactArgs(1),
	}

	rootCmd := &cobra.Command{
		Use:               "devto",
		Short:             "Publish to dev.to from your terminal",
		PersistentPreRunE: r.rootRunE,
	}

	rootCmd.PersistentFlags().String(flagAPIKey, "", "API key for authentication")
	rootCmd.PersistentFlags().BoolP(flagDebug, "d", false, "Print debug log on stderr")
	rootCmd.AddCommand(
		listCmd,
		submitCmd,
	)

	_ = viper.BindPFlag(flagAPIKey, rootCmd.PersistentFlags().Lookup(flagAPIKey))

	return rootCmd, func() { _ = r.log.Sync() }
}

type runner struct {
	log *zap.SugaredLogger
}

func (r *runner) rootRunE(cmd *cobra.Command, args []string) error {
	// Setup logger
	logConfig := zap.NewDevelopmentConfig()

	isDebug, err := cmd.Parent().PersistentFlags().GetBool(flagDebug)
	if err != nil {
		return errors.Wrap(err, "cmd: get flag --debug")
	}

	if !isDebug {
		logConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logger, err := logConfig.Build()
	if err != nil {
		return errors.Wrap(err, "cmd: create new logger")
	}

	r.log = logger.Sugar()

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

func (r *runner) listRunE(cmd *cobra.Command, args []string) error {
	apiKey := context.WithValue(context.Background(), devto.ContextAPIKey, devto.APIKey{
		Key: viper.GetString(flagAPIKey),
	})

	client := devto.NewAPIClient(devto.NewConfiguration())

	articles, _, err := client.ArticlesApi.GetUserAllArticles(apiKey, nil)
	if err != nil {
		return errors.Wrap(err, "cmd: get articles")
	}

	for _, a := range articles {
		r.log.Infow("", "title", a.Title, "ID", a.Id)
	}

	return nil
}

func (r *runner) submitRunE(cmd *cobra.Command, args []string) error {
	apiKey := context.WithValue(context.Background(), devto.ContextAPIKey, devto.APIKey{
		Key: viper.GetString(flagAPIKey),
	})

	body, err := article.Read(args[0])
	if err != nil {
		return err
	}

	client := devto.NewAPIClient(devto.NewConfiguration())

	article := &devto.ArticlesApiCreateArticleOpts{
		ArticleCreate: optional.NewInterface(devto.ArticleCreate{
			Article: devto.ArticleCreateArticle{
				BodyMarkdown: body,
			},
		},
		),
	}

	submitted, _, err := client.ArticlesApi.CreateArticle(apiKey, article)
	if err != nil {
		return errors.Wrap(err, "cmd: get articles")
	}

	r.log.Info(submitted.Id)

	return nil
}
