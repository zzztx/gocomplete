/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	build "gocomplete/cmd/build"
	clean "gocomplete/cmd/clean"
	complete "gocomplete/cmd/completion"
	env "gocomplete/cmd/env"
	fix "gocomplete/cmd/fix"
	fmt "gocomplete/cmd/fmt"
	generate "gocomplete/cmd/generate"
	get "gocomplete/cmd/get"
	install "gocomplete/cmd/install"
	list "gocomplete/cmd/list"
	mod "gocomplete/cmd/mod"
	run "gocomplete/cmd/run"
	test "gocomplete/cmd/test"
	tool "gocomplete/cmd/tool"
	version "gocomplete/cmd/version"
	vet "gocomplete/cmd/vet"
	"log"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func init() {
	streams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	RootCmd.AddCommand(complete.NewCmdCompletion(streams))
	RootCmd.AddCommand(build.NewCmdBuild(streams))
	RootCmd.AddCommand(get.NewCmdGet(streams))
	RootCmd.AddCommand(run.NewCmdRun(streams))
	RootCmd.AddCommand(version.NewCmdVersion(streams))
	RootCmd.AddCommand(test.NewCmdTest(streams))
	RootCmd.AddCommand(mod.NewCmdMod(streams))
	RootCmd.AddCommand(list.NewCmdList(streams))
	RootCmd.AddCommand(install.NewCmdInstall(streams))
	RootCmd.AddCommand(tool.NewCmdTool(streams))
	RootCmd.AddCommand(vet.NewCmdVet(streams))
	RootCmd.AddCommand(clean.NewCmdClean(streams))
	RootCmd.AddCommand(env.NewCmdEnv(streams))
	RootCmd.AddCommand(fix.NewCmdFix(streams))
	RootCmd.AddCommand(fmt.NewCmdFmt(streams))
	RootCmd.AddCommand(generate.NewCmdGenerate(streams))
}

var RootCmd = &cobra.Command{
	Use: "go",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Go Cli Tools")
	},
}
