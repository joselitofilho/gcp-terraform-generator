package config

import (
	"errors"
	"os"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/require"
)

var (
	testdataFolder = "../testdata"

	errDummy = errors.New("dummy error")
)

func TestYAML_Parse(t *testing.T) {
	type fields struct {
		fileName string
	}

	tests := []struct {
		setup     func(testing.TB) func(testing.TB)
		name      string
		fields    fields
		want      *Config
		targetErr error
	}{
		{
			setup:  func(_ testing.TB) func(testing.TB) { return func(_ testing.TB) {} },
			name:   "Diagram",
			fields: fields{fileName: testdataFolder + "/diagram.config.yaml"},
			want: &Config{Diagram: &Diagram{
				StackName: "teststack",
			}},
		},
		{
			setup: func(_ testing.TB) func(testing.TB) {
				osReadFile = func(name string) ([]byte, error) {
					require.Empty(t, name)
					return nil, errDummy
				}

				return func(_ testing.TB) {
					osReadFile = os.ReadFile
				}
			},
			name:      "Empty File",
			fields:    fields{fileName: ""},
			targetErr: errDummy,
		},
		{
			setup: func(_ testing.TB) func(testing.TB) {
				yamlUnmarshal = func(_ []byte, _ any) error {
					return errDummy
				}
				return func(_ testing.TB) {
					yamlUnmarshal = yaml.Unmarshal
				}
			},
			name:      "Invalid YAML Syntax",
			fields:    fields{fileName: testdataFolder + "/invalid_sintax.yaml"},
			targetErr: errDummy,
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			tearDown := tc.setup(t)
			defer tearDown(t)

			got, err := NewYAML(tc.fields.fileName).Parse()

			require.ErrorIs(t, err, tc.targetErr)
			require.Equal(t, tc.want, got)
		})
	}
}
