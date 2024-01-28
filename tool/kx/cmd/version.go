// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/bitbus/kx/tool/kx/cmd/root"
	"github.com/spf13/cobra"
)

const (
	version = "v0.0.1"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "show version information",
		Long:  "show version information",
		Run:   versionRun,
	}
	root.Register(versionCmd)
}

func versionRun(_cmd *cobra.Command, _args []string) {
	fmt.Println(version)
}
