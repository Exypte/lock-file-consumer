package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Exypte/gh-actions-lock/internal/yamlprint"
)

// newPrintCmd returns the print command, which reads its settings from v.
func newPrintCmd(v *viper.Viper) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "print <file>",
		Short: "Parse a YAML file and print it",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := v.BindPFlag("indent", cmd.Flags().Lookup("indent")); err != nil {
				return err
			}
			opts := yamlprint.Options{Indent: v.GetInt("indent")}
			return yamlprint.Print(cmd.OutOrStdout(), args[0], opts)
		},
	}
	cmd.Flags().Int("indent", yamlprint.DefaultIndent, "number of spaces used for indentation")
	return cmd
}
