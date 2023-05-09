package generate

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdGenerate(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "generate",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go generate")),
		Long:                  green(i18n.T("go generate")),
		Example:               green(i18n.T("go generate")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
