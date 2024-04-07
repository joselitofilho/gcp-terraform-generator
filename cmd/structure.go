package cmd

import (
	"github.com/spf13/cobra"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/structure"
)

// structureCmd represents the structure command.
var structureCmd = &cobra.Command{
	Use:   "structure",
	Short: "Manage Structure",
	Run: func(cmd *cobra.Command, _ []string) {
		config, _ := cmd.Flags().GetString(flagConfig)
		output, _ := cmd.Flags().GetString(flagOutput)

		if err := structure.NewStructure(config, output).Build(); err != nil {
			printErrorAndExit(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(structureCmd)

	structureCmd.Flags().StringP(flagConfig, "c", "",
		"Path to the configuration file. For example: ./structure.config.yaml")
	structureCmd.Flags().StringP(flagOutput, "o", "",
		"Path to the output folder. For example: ./output")

	_ = structureCmd.MarkFlagRequired(flagConfig)
	_ = structureCmd.MarkFlagRequired(flagOutput)
}
