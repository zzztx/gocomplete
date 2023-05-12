package build

import (
	"log"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdBuild(ioStreams genericclioptions.IOStreams) *cobra.Command {
	shells := []string{}
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), "go") {
			shells = append(shells, file.Name())
		}
	}
	sort.Strings(shells)
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "build",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go build file.go")),
		Long:                  green(i18n.T("go build file.go")),
		Example:               green(i18n.T("go build file.go")),
		Run: func(cmd *cobra.Command, args []string) {
		},
		ValidArgs: shells,
	}

	return cmd
}
