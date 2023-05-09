package vet

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdVet(ioStreams genericclioptions.IOStreams) *cobra.Command {
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "vet",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go vet")),
		Long:                  green(i18n.T("go vet")),
		Example:               green(i18n.T("go vet")),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
