// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/bitbus/kx/tool/kx/cmd/root"
	"github.com/spf13/cobra"
)

func init() {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "config custom setup",
		Long:  "config custom setup",
		Run:   configRun,
	}
	root.Register(configCmd)
}

func configRun(_cmd *cobra.Command, _args []string) {
	// TODO
}
