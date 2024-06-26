package terraformtoresources

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	hcl "github.com/joselitofilho/hcl-parser-go/pkg/parser/hcl"
)

const (
	EmptyAttributeName = ""
)

func buildKeyValueFromLocals(tfLocals []*hcl.Local) map[string]string {
	keyValue := map[string]string{}

	for i := range tfLocals {
		for k, v := range tfLocals[i].Attributes {
			varName := fmt.Sprintf("local.%s", k)

			switch v := v.(type) {
			case string:
				keyValue[varName] = v
			case []string:
				buildSliceStringVars(varName, v, keyValue)
			case map[string]any:
				buildStringAnyMapVars(varName, v, keyValue)
			default:
				keyValue[varName] = varName
			}
		}
	}

	return keyValue
}

func buildKeyValueFromVariables(tfVariables []*hcl.Variable) map[string]string {
	keyValue := map[string]string{}

	for i := range tfVariables {
		for k, v := range tfVariables[i].Attributes {
			varName := fmt.Sprintf("var.%s", k)

			switch v := v.(type) {
			case string:
				keyValue[varName] = v
			case []string:
				buildSliceStringVars(varName, v, keyValue)
			case map[string]any:
				buildStringAnyMapVars(varName, v, keyValue)
			default:
				keyValue[varName] = varName
			}
		}
	}

	return keyValue
}

func buildSliceStringVars(varName string, values []string, keyValue map[string]string) {
	if len(values) > 0 {
		keyValue[varName] = values[0]
	} else {
		keyValue[varName] = varName
	}
}

func buildStringAnyMapVars(varName string, values map[string]any, keyValue map[string]string) {
	arr := make([]string, 0, len(values))
	for k := range values {
		arr = append(arr, k)
	}

	if len(arr) > 0 {
		slices.Sort(arr)
		keyValue[varName] = arr[0]
	} else {
		keyValue[varName] = varName
	}
}

func getResourceNameFromLabel(label, suffix string) string {
	result := label

	if strings.HasSuffix(result, suffix) {
		result = result[:len(label)-len(suffix)]
	}

	return result
}

func replaceVars(
	str string, tfVars []*hcl.Variable, tfLocals []*hcl.Local, replaceableStrs map[string]string,
) string {
	keyValue := buildKeyValueFromVariables(tfVars)
	str = replaceVariables(str, keyValue)

	keyValue = buildKeyValueFromLocals(tfLocals)
	str = replaceLocals(str, keyValue)

	str = replaceStrings(str, replaceableStrs)

	return str
}

func replaceVariables(str string, keyValue map[string]string) string {
	for i := 0; i <= len(keyValue); i++ {
		for varName, finalValue := range keyValue {
			str = strings.ReplaceAll(str, varName, finalValue)
		}

		if !strings.Contains(str, "var.") {
			break
		}
	}

	return str
}

func replaceLocals(str string, keyValue map[string]string) string {
	for i := 0; i <= len(keyValue); i++ {
		for varName, finalValue := range keyValue {
			str = strings.ReplaceAll(str, varName, finalValue)
		}

		if !strings.Contains(str, "local.") {
			break
		}
	}

	return str
}

func replaceStrings(str string, replaceableStrs map[string]string) string {
	for i := 0; i <= len(replaceableStrs); i++ {
		for varName, finalValue := range replaceableStrs {
			str = strings.ReplaceAll(str, varName, finalValue)
		}

		for varName := range replaceableStrs {
			if !strings.Contains(str, varName) {
				break
			}
		}
	}

	return str
}

func extractTextFromTFVar(s string) string {
	re := regexp.MustCompile(`\${(.*?)}`)
	matches := re.FindAllStringSubmatch(s, -1)

	for _, match := range matches {
		s = strings.Replace(s, match[0], match[1], 1)
	}

	return s
}
