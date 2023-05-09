package list

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdList(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "list",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go list")),
		Long:                  green(i18n.T("go list")),
		Example:               green(i18n.T("go list")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
