package build

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdBuild(ioStreams genericclioptions.IOStreams) *cobra.Command {
	shells := []string{}
	shells = append(shells, "main.go")
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "completion SHELL",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go build main.go")),
		Long:                  green(i18n.T("go build main.go")),
		Example:               green(i18n.T("go build main.go")),
		Run: func(cmd *cobra.Command, args []string) {
		},
		ValidArgs: shells,
	}

	return cmd
}
