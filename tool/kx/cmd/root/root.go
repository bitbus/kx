// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package root

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "kx",
		Short: "kitex help toolkit",
		Long:  `kitex help toolkit`,
	}
)

// Setup set root command name,short-describe, long-describe
// return &cobra.Command to custom other options
func Setup(use, short, long string) *cobra.Command {
	rootCmd.Use = use
	rootCmd.Short = short
	rootCmd.Long = long
	return rootCmd
}

// register add sub-command
func Register(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

// Execute start application
func Execute() {
	rootCmd.Execute()
}
