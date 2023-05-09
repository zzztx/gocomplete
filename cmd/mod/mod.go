package mod

import (
	"sort"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdMod(ioStreams genericclioptions.IOStreams) *cobra.Command {
	shells := []string{"download", "edit", "graph", "init", "tidy", "vendor", "verify", "why"}
	sort.Strings(shells)
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "mod",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go mod")),
		Long:                  green(i18n.T("go mod")),
		Example:               green(i18n.T("go mod")),
		Run: func(cmd *cobra.Command, args []string) {
		},
		ValidArgs: shells,
	}

	return cmd
}
