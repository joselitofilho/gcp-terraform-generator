package cmd

import (
	"errors"
	"os"

	"github.com/joselitofilho/gcp-terraform-generator/internal/fmtcolor"
	"github.com/spf13/cobra"
)

const (
	flagConfig  = "config"
	flagDiagram = "diagram"
	flagFile    = "file"
	flagLeft    = "left"
	flagOutput  = "output"
	flagRight   = "right"
	flagWorkdir = "workdir"
)

var (
	ErrNoDiagramOrConfigFiles = errors.New("this directory does not contain any diagram or config files")
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "gcp-terraform-generator",
	Short: "GCP terraform generator",
	Run:   func(cmd *cobra.Command, _ []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP(flagWorkdir, "", ".",
		"Path to the directory where diagrams and configuration files are stored for the project. For example: ./example")
}

func printErrorAndExit(err error) {
	fmtcolor.Red.Printf("🚨 %s\n", err)
	os.Exit(1)
}
