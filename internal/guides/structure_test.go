package guides

import (
	"testing"

	"github.com/AlecAivazis/survey/v2"

	surveyasker "github.com/joselitofilho/gcp-terraform-generator/internal/survey"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGuideStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		surveyAsker surveyasker.Asker
		workdir     string
		fileMap     map[string][]string
	}

	tests := []struct {
		name         string
		args         args
		prepareMocks func(surveyasker.Asker)
		want         *StructureAnswers
		targetErr    error
	}{
		{
			name: "happy path",
			args: args{
				surveyAsker: &fakeStructureSurveyAsker{
					TB: t,
					Answers: StructureAnswers{
						Config: "structure.config.yaml",
						Output: ".//testoutput",
					},
					AskPerCall: []survey.Prompt{
						&survey.Select{
							Message: "Choose a config:",
							Default: 1,
							Options: []string{"config.yaml", "structure.config.yaml"},
						},
						&survey.Input{
							Message: "Enter the output folder:",
							Default: "./output",
						},
					},
				},
				workdir: "./testoutput/teststack",
				fileMap: map[string][]string{"config": {"config.yaml", "structure.config.yaml"}},
			},
			prepareMocks: func(_ surveyasker.Asker) {},
			want: &StructureAnswers{
				Config: "./testoutput/teststack/structure.config.yaml",
				Output: "./testoutput",
			},
		},
		{
			name: "when there is no config files should return an error",
			args: args{
				workdir: "./testoutput/teststack",
				fileMap: map[string][]string{"config": {}},
			},
			prepareMocks: func(_ surveyasker.Asker) {},
			targetErr:    ErrDirDoesNotContainAnyConfigFile,
		},
		{
			name: "when survey to choose the config fails should return an error",
			args: args{
				surveyAsker: surveyasker.NewMockAsker(ctrl),
				workdir:     "./testoutput/teststack",
				fileMap:     map[string][]string{"config": {"diagram.yaml"}},
			},
			prepareMocks: func(a surveyasker.Asker) {
				msa := a.(*surveyasker.MockAsker)

				gomock.InOrder(
					msa.EXPECT().AskOne(gomock.Any(), gomock.Any()).Return(errDummy),
				)
			},
			targetErr: errDummy,
		},
		{
			name: "when survey to enter the output fails should return an error",
			args: args{
				surveyAsker: surveyasker.NewMockAsker(ctrl),
				workdir:     "./testoutput/teststack",
				fileMap:     map[string][]string{"config": {"diagram.yaml"}},
			},
			prepareMocks: func(a surveyasker.Asker) {
				msa := a.(*surveyasker.MockAsker)

				gomock.InOrder(
					msa.EXPECT().AskOne(gomock.Any(), gomock.Any()).Return(nil),
					msa.EXPECT().AskOne(gomock.Any(), gomock.Any()).Return(errDummy),
				)
			},
			targetErr: errDummy,
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			tc.prepareMocks(tc.args.surveyAsker)

			got, err := GuideStructure(tc.args.surveyAsker, tc.args.workdir, tc.args.fileMap)

			if tc.targetErr == nil {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			} else {
				require.ErrorIs(t, err, tc.targetErr)
			}
		})
	}
}

type fakeStructureSurveyAsker struct {
	TB         testing.TB
	Answers    StructureAnswers
	AskPerCall []survey.Prompt
	callCount  int
}

func (f *fakeStructureSurveyAsker) Ask(_ []*survey.Question, _ any, _ ...survey.AskOpt) error {
	return nil
}

func (f *fakeStructureSurveyAsker) AskOne(prompt survey.Prompt, response any, _ ...survey.AskOpt) error {
	f.callCount++

	ans := response.(*string)
	require.Empty(f.TB, ans)

	switch f.callCount {
	case 1:
		require.IsType(f.TB, &survey.Select{}, prompt)
		require.Equal(f.TB, f.AskPerCall[f.callCount-1], prompt)

		*ans = f.Answers.Config
	case 2:
		require.IsType(f.TB, &survey.Input{}, prompt)
		require.Equal(f.TB, f.AskPerCall[f.callCount-1], prompt)

		*ans = f.Answers.Output
	}

	return nil
}
