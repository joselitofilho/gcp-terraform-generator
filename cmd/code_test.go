package cmd

import (
	_ "embed"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCode_Run(t *testing.T) {
	type args struct {
		configFile string
		output     string
	}

	tests := []struct {
		name             string
		args             args
		setup            func() (tearDown func())
		extraValidations func(testing.TB)
	}{
		{
			name: "happy path",
			args: args{
				configFile: path.Join(testdataFolder, "code.config.yaml"),
				output:     path.Join(testOutput, "teststack"),
			},
			extraValidations: func(tb testing.TB) {
				teststackPath := path.Join(testOutput, "teststack")

				require.FileExists(tb, path.Join(teststackPath, "mod", "my-engine.tf"))
			},
		},
		{
			name: "config file does not exist",
			args: args{
				configFile: "fileDoesNotExist.xml",
				output:     testOutput,
			},
			setup: func() (tearDown func()) {
				osExit = func(code int) {
					require.Equal(t, 1, code)
				}

				return func() {
					osExit = os.Exit
				}
			},
		},
	}

	defer func() {
		_ = os.RemoveAll(testOutput)
	}()

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			if tc.setup != nil {
				tearDown := tc.setup()
				defer tearDown()
			}

			_ = codeCmd.Flags().Set(flagConfig, tc.args.configFile)
			_ = codeCmd.Flags().Set(flagOutput, tc.args.output)

			codeCmd.Run(codeCmd, []string{})

			if tc.extraValidations != nil {
				tc.extraValidations(t)
			}
		})
	}
}
