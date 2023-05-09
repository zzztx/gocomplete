package fmt

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdFmt(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "fmt",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go fmt")),
		Long:                  green(i18n.T("go fmt")),
		Example:               green(i18n.T("go fmt")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
