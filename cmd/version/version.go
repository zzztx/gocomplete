package version

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdVersion(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "version",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go version")),
		Long:                  green(i18n.T("go version")),
		Example:               green(i18n.T("go version")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
