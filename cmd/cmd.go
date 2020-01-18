package cmd

import (
	"os"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/shihanng/devto/pkg/article"
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
		Short: "List published articles (maximum 30) on dev.to",
		Long: `List published articles (maximum 30) on dev.to in the following format:

   [<article_id>] <title>
`,
		RunE: r.listRunE,
		Args: cobra.ExactArgs(0),
	}

	submitCmd := &cobra.Command{
		Use:   "submit <Markdown file>",
		Short: "Submit article to dev.to",
		RunE:  r.submitRunE,
		Args:  cobra.ExactArgs(1),
	}

	generateCmd := &cobra.Command{
		Use:   "generate <Markdown file>",
		Short: "Genenerate a devto.yml configuration file for the <Markdown file>",
		RunE:  r.generateRunE,
		Args:  cobra.ExactArgs(1),
	}

	rootCmd := &cobra.Command{
		Use:               "devto",
		Short:             "A tool to help you publish to dev.to from your terminal",
		PersistentPreRunE: r.rootRunE,
	}

	rootCmd.PersistentFlags().String(flagAPIKey, "", "API key for authentication")
	rootCmd.PersistentFlags().BoolP(flagDebug, "d", false, "Print debug log on stderr")
	rootCmd.AddCommand(
		listCmd,
		submitCmd,
		generateCmd,
	)

	_ = viper.BindPFlag(flagAPIKey, rootCmd.PersistentFlags().Lookup(flagAPIKey))

	return rootCmd, func() {
		if r.log != nil {
			_ = r.log.Sync()
		}
	}
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
	client, err := article.NewClient(viper.GetString(flagAPIKey))
	if err != nil {
		return err
	}

	return client.ListArticle(os.Stdout)
}

func (r *runner) submitRunE(cmd *cobra.Command, args []string) error {
	filename := args[0]

	client, err := article.NewClient(viper.GetString(flagAPIKey), article.SetConfig(filename))
	if err != nil {
		return err
	}

	return client.SubmitArticle(filename)
}

func (r *runner) generateRunE(cmd *cobra.Command, args []string) error {
	filename := args[0]

	client, err := article.NewClient(viper.GetString(flagAPIKey), article.SetConfig(filename))
	if err != nil {
		return err
	}

	return client.GenerateImageLinks(filename)
}
