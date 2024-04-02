package cmd

import (
	"fmt"

	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/diagram"

	"github.com/spf13/cobra"
)

// diagramCmd represents the structure command.
var diagramCmd = &cobra.Command{
	Use:   "diagram",
	Short: "Manage Diagram",
	Run: func(cmd *cobra.Command, _ []string) {
		diagramFilename, err := cmd.Flags().GetString(flagDiagram)
		if err != nil {
			printErrorAndExit(err)
		}

		configFile, err := cmd.Flags().GetString(flagConfig)
		if err != nil {
			printErrorAndExit(err)
		}

		output, err := cmd.Flags().GetString(flagOutput)
		if err != nil {
			printErrorAndExit(err)
		}

		if err := diagram.NewDiagram(diagramFilename, configFile, output).Build(); err != nil {
			printErrorAndExit(err)
		}

		fmt.Printf("Configuration file '%s' has been generated successfully\n", output)
	},
}

func init() {
	rootCmd.AddCommand(diagramCmd)

	diagramCmd.Flags().StringP(flagDiagram, "d", "", "Path to the XML file. For example: ./diagram.xml")
	diagramCmd.Flags().StringP(flagConfig, "c", "",
		"Path to the YAML config file. For example: ./diagram.config.yaml")
	diagramCmd.Flags().StringP(flagOutput, "o", "", "Path to the output file. For example: ./diagram.yaml")

	_ = diagramCmd.MarkFlagRequired(flagDiagram)
	_ = diagramCmd.MarkFlagRequired(flagConfig)
	_ = diagramCmd.MarkFlagRequired(flagOutput)
}
