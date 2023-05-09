package env

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdEnv(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "env",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go env")),
		Long:                  green(i18n.T("go env")),
		Example:               green(i18n.T("go env")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
