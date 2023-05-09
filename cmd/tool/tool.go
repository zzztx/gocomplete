package tool

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdTool(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "tool",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go tool")),
		Long:                  green(i18n.T("go tool")),
		Example:               green(i18n.T("go tool")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
