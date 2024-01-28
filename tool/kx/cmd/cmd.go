// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	_ "github.com/bitbus/kx/tool/kx/cmd/config"
	"github.com/bitbus/kx/tool/kx/cmd/root"
)

func Execute() {
	root.Execute()
}
