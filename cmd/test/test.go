package test

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdTest(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "test",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go test")),
		Long:                  green(i18n.T("go test")),
		Example:               green(i18n.T("go test")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
