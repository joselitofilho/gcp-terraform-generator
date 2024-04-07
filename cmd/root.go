package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"github.com/joselitofilho/gcp-terraform-generator/internal/fmtcolor"
	"github.com/joselitofilho/gcp-terraform-generator/internal/guides"
	surveyasker "github.com/joselitofilho/gcp-terraform-generator/internal/survey"
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

const (
	optionGuideDiagram          = "Generate a diagram config file"
	optionGuideInitialStructure = "Generate the initial structure"
	optionGuideCode             = "Generate code"
	optionExit                  = "Exit"
)

var (
	ErrNoDiagramOrConfigFiles = errors.New("this directory does not contain any diagram or config files")
)

var osExit = os.Exit

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "gcp-terraform-generator",
	Short: "GCP terraform generator",
	Run: func(cmd *cobra.Command, _ []string) {
		workdir, _ := cmd.Flags().GetString(flagWorkdir)

		title := `

		 ██████╗ ██████╗ ██████╗ ███████╗     ██████╗ ███████╗███╗   ██╗
		██╔════╝██╔═══██╗██╔══██╗██╔════╝    ██╔════╝ ██╔════╝████╗  ██║
		██║     ██║   ██║██║  ██║█████╗      ██║  ███╗█████╗  ██╔██╗ ██║
		██║     ██║   ██║██║  ██║██╔══╝      ██║   ██║██╔══╝  ██║╚██╗██║
		╚██████╗╚██████╔╝██████╔╝███████╗    ╚██████╔╝███████╗██║ ╚████║
		 ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝     ╚═════╝ ╚══════╝╚═╝  ╚═══╝
                                                                             GCP

`
		fmtcolor.White.Println(title)

		surveyAsker := &surveyasker.RealAsker{}

		shouldContinue := true
		for shouldContinue {
			// Create a map to store file extensions and their respective files.
			fileMap := make(map[string][]string)

			// Read files in the current directory.
			files, err := os.ReadDir(workdir)
			if err != nil {
				printErrorAndExit(fmt.Errorf("error reading directory: %w", err))
			}

			// Iterate through the files and populate the map.
			for _, file := range files {
				if !file.IsDir() {
					ext := strings.ToLower(path.Ext(file.Name()))

					switch ext {
					case ".xml":
						fileMap[flagDiagram] = append(fileMap[flagDiagram], file.Name())
					case ".yaml", ".yml":
						fileMap[flagConfig] = append(fileMap[flagConfig], file.Name())
					}
				}
			}

			if len(fileMap) == 0 {
				printErrorAndExit(ErrNoDiagramOrConfigFiles)
			}

			var commandName string
			if err := survey.AskOne(&survey.Select{
				Message: "What would you like to do?",
				Options: []string{
					optionGuideDiagram,
					optionGuideInitialStructure,
					optionGuideCode,
					optionExit,
				},
			}, &commandName); err != nil {
				printErrorAndExit(err)
			}

			switch commandName {
			case optionGuideDiagram:
				answers, err := guides.GuideDiagram(surveyAsker, workdir, fileMap)
				if err != nil {
					printErrorAndExit(err)
				}

				_ = diagramCmd.Flags().Set(flagDiagram, answers.Diagram)
				_ = diagramCmd.Flags().Set(flagConfig, answers.Config)
				_ = diagramCmd.Flags().Set(flagOutput, answers.Output)
				diagramCmd.Run(diagramCmd, []string{})
			case optionGuideInitialStructure:
				answers, err := guides.GuideStructure(surveyAsker, workdir, fileMap)
				if err != nil {
					printErrorAndExit(err)
				}

				_ = structureCmd.Flags().Set(flagConfig, answers.Config)
				_ = structureCmd.Flags().Set(flagOutput, answers.Output)
				structureCmd.Run(structureCmd, []string{})
			case optionGuideCode:
				answers, err := guides.GuideCode(surveyAsker, workdir, fileMap)
				if err != nil {
					printErrorAndExit(err)
				}

				stackOutput := fmt.Sprintf("%s/%s", answers.Output, answers.StackName)

				_ = codeCmd.Flags().Set(flagConfig, answers.Config)
				_ = codeCmd.Flags().Set(flagOutput, stackOutput)
				codeCmd.Run(codeCmd, []string{})
				fmt.Println()
			default:
				shouldContinue = false
			}

			if shouldContinue {
				fmt.Println()
			}
		}

		fmtcolor.White.Println("👋 Goodbye. Until next time!")
	},
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
	osExit(1)
}
