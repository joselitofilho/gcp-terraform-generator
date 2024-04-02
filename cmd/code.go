package cmd

import (
	"github.com/spf13/cobra"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/code"
)

// structureCmd represents the structure command.
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Manage Code",
	Run: func(cmd *cobra.Command, _ []string) {
		config, err := cmd.Flags().GetString(flagConfig)
		if err != nil {
			printErrorAndExit(err)
		}

		output, err := cmd.Flags().GetString(flagOutput)
		if err != nil {
			printErrorAndExit(err)
		}

		err = code.NewCode(config, output).Build()
		if err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	codeCmd.Flags().StringP(flagConfig, "c", "",
		"Path to the configuration file. For example: ./diagram.yaml")
	codeCmd.Flags().StringP(flagOutput, "o", "",
		"Path to the output folder. For example: ./output")

	_ = codeCmd.MarkFlagRequired(flagConfig)
	_ = codeCmd.MarkFlagRequired(flagOutput)
}
