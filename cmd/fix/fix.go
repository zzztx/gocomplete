package fix

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdFix(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "fix",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go fix")),
		Long:                  green(i18n.T("go fix")),
		Example:               green(i18n.T("go fix")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
