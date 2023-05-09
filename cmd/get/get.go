package get

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdGet(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "get",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go get package")),
		Long:                  green(i18n.T("go get package")),
		Example:               green(i18n.T("go get package")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
