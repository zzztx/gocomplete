package clean

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdClean(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "clean",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go clean")),
		Long:                  green(i18n.T("go clean")),
		Example:               green(i18n.T("go clean")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
