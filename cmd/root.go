// Package cmd implements the gh-actions-lock command-line interface.
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// envPrefix is the prefix of environment variables read as configuration,
	// e.g. GH_ACTIONS_LOCK_INDENT for the "indent" key.
	envPrefix = "GH_ACTIONS_LOCK"
	// configName is the name, without extension, of the default config file.
	configName = ".gh-actions-lock"
)

// Execute runs the root command and returns any error it produced.
func Execute() error {
	return newRootCmd().Execute()
}

// newRootCmd returns the root command wired to a fresh viper instance.
func newRootCmd() *cobra.Command {
	v := viper.New()
	var cfgFile string

	root := &cobra.Command{
		Use:          "gh-actions-lock",
		Short:        "Parse and print YAML files",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig(v, cfgFile)
		},
	}
	root.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is ./"+configName+".yaml or $HOME/"+configName+".yaml)")
	root.AddCommand(newPrintCmd(v))
	return root
}

// initConfig loads configuration into v from the environment and from
// cfgFile, or from the default config file locations when cfgFile is empty.
// A missing default config file is not an error.
func initConfig(v *viper.Viper, cfgFile string) error {
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		v.SetConfigName(configName)
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		if home, err := os.UserHomeDir(); err == nil {
			v.AddConfigPath(home)
		}
	}

	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if cfgFile == "" && errors.As(err, &notFound) {
			return nil
		}
		return fmt.Errorf("read config: %w", err)
	}
	return nil
}
