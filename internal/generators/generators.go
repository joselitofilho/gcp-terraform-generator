package generators

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/ettle/strcase"

	"github.com/joselitofilho/gcp-terraform-generator/internal/fmtcolor"
	"github.com/joselitofilho/gcp-terraform-generator/internal/generators/config"
	"github.com/joselitofilho/gcp-terraform-generator/internal/utils"
)

var ErrUnsupportedFileType = errors.New("unsupported file type")

// Build executes the provided templateContent using the data supplied and returns the resulting string.
func Build(data any, templateName, templateContent string) (string, error) {
	tmpl, err := buildAndParseTemplate(templateName, templateContent)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	var output bytes.Buffer

	err = tmpl.Execute(&output, data)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return output.String(), nil
}

// CreateFilesMap takes a slice of config.File structs and creates a map where the keys are strings and the values are
// File structs. Each element in the slice is used to populate the corresponding key-value pair in the map.
func CreateFilesMap(files []config.File) map[string]File {
	filesConf := map[string]File{}
	for i := range files {
		filesConf[files[i].Name] = File{
			Tmpl:    files[i].Tmpl,
			Imports: files[i].Imports,
		}
	}

	return filesConf
}

// CreateTemplatesMap creates a map[string]string from a slice of config.FilenameTemplateMap. It takes a slice of
// config.FilenameTemplateMap as input. Each element in the slice represents a map where keys are filenames (strings)
// and values are corresponding templates (strings). It iterates over each map in the slice and merges them into a
// single map[string]string where each filename is mapped to its respective template.
func CreateTemplatesMap(filenameTemplatesList []config.FilenameTemplateMap) map[string]string {
	templatesMap := map[string]string{}

	for i := range filenameTemplatesList {
		for filename, tmpl := range filenameTemplatesList[i] {
			templatesMap[filename] = tmpl
		}
	}

	return templatesMap
}

// FilterTemplatesMap filters a map of filenames to templates based on a given filter string. It takes a filter string
// and a map of filenames to templates as input. It iterates over each key-value pair in the input map and adds those
// pairs to a new map if the filename contains the filter string. The function then returns the filtered map.
func FilterTemplatesMap(filter string, templatesMap map[string]string) map[string]string {
	filtred := map[string]string{}

	for filename, tmpl := range templatesMap {
		if strings.Contains(filename, filter) {
			filtred[filename] = tmpl
		}
	}

	return filtred
}

// GenerateFile generates a file using the provided template and data, and writes the output to the specified
// outputFile. If fileTmpl is provided, it uses that specific template; otherwise, it looks up the template based on the
// fileName in the templatesMap. The generated file is then formatted based on its extension.
func GenerateFile(templatesMap map[string]string, fileName, fileTmpl, outputFile string, data any) error {
	var (
		tmpl     string
		tmplName = fmt.Sprintf("%s-template", strings.ReplaceAll(fileName, ".", "-"))
	)

	if fileTmpl == "" {
		tmpl = templatesMap[fileName]
	} else {
		tmpl = fileTmpl
	}

	err := buildFile(data, tmplName, tmpl, outputFile)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = formatFileBasedOnExt(fileName, outputFile)
	if err != nil && !errors.Is(err, ErrUnsupportedFileType) {
		fmtcolor.Yellow.Println(err)
		err = nil
	}

	return err
}

// GenerateFiles generates multiple files using the provided templates and data, and writes the outputs to the specified
// output directory. It iterates over templates from the templatesMap and files from the filesMap, merging them into a
// single map of templates. For each template, it generates a file, applies formatting based on its extension, and saves
// it to the output directory.
func GenerateFiles(templatesMap map[string]string, filesMap map[string]File, data any, output string) error {
	mergedTemplates := map[string]string{}

	for filename, tmpl := range templatesMap {
		mergedTemplates[filename] = tmpl
	}

	for filename, file := range filesMap {
		mergedTemplates[filename] = file.Tmpl
	}

	for filename, fileTmpl := range mergedTemplates {
		tmplName := fmt.Sprintf("%s-template", strings.ReplaceAll(filename, ".", "-"))

		outputFile := path.Join(output, filename)

		err := buildFile(data, tmplName, fileTmpl, outputFile)
		if err != nil {
			// TODO: Append error
			fmtcolor.Yellow.Println(err)
		}

		err = formatFileBasedOnExt(filename, outputFile)
		if err != nil && !errors.Is(err, ErrUnsupportedFileType) {
			// TODO: Append error
			fmtcolor.Yellow.Println(err)
		}
	}

	// TODO: Return errors
	return nil
}

func buildAndParseTemplate(name, content string) (*template.Template, error) {
	tmpl, err := template.New(name).
		Funcs(template.FuncMap{
			"getFileByName":  func(files map[string]File, name string) File { return files[name] },
			"getFileImports": func(files map[string]File, name string) []string { return files[name].Imports },
			"ToCamel":        strcase.ToCamel,
			"ToKebab":        strcase.ToKebab,
			"ToLower":        strings.ToLower,
			"ToPascal":       strcase.ToPascal,
			"ToSpace":        func(s string) string { return strings.ReplaceAll(strcase.ToKebab(s), "-", " ") },
			"ToSnake":        strcase.ToSnake,
			"ToUpper":        strings.ToUpper,
		}).
		Parse(content)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return tmpl, nil
}

func buildFile(data any, templateName, templateContent, outputPath string) error {
	tmpl, err := buildAndParseTemplate(templateName, templateContent)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer output.Close()

	err = tmpl.Execute(output, data)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func formatFileBasedOnExt(fileName, outputFile string) error {
	var err error

	ext := path.Ext(fileName)

	switch ext {
	case ".go":
		err = utils.GoFormat(outputFile)
	case ".tf":
		err = utils.TerraformFormat(outputFile)
	default:
		err = nil
	}

	return err
}
