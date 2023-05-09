package install

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdInstall(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "install",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go install")),
		Long:                  green(i18n.T("go install")),
		Example:               green(i18n.T("go install")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
