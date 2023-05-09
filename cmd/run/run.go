package run

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/util/i18n"
)

func NewCmdRun(ioStreams genericclioptions.IOStreams) *cobra.Command {
	shells := []string{"main.go"}
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	theOs := runtime.GOOS
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			if theOs == "windows" && strings.HasSuffix(file.Name(), "exe") {
				shells = append(shells, file.Name())
			}
			if (theOs == "linux" || theOs == "darwin") && filepath.Ext(file.Name()) == "" {
				shells = append(shells, file.Name())
			}
		}
	}
	sort.Strings(shells)
	green := color.New(color.FgGreen).SprintFunc()

	cmd := &cobra.Command{
		Use:                   "run",
		DisableFlagsInUseLine: true,
		Short:                 green(i18n.T("go run main.go")),
		Long:                  green(i18n.T("go run main.go")),
		Example:               green(i18n.T("go run main.go")),
		Run: func(cmd *cobra.Command, args []string) {
		},
		ValidArgs: shells,
	}

	return cmd
}
