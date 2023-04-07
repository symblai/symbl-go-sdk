// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package symbl

import (
	"flag"
	"strconv"

	klog "k8s.io/klog/v2"
)

type LogLevel int64

const (
	LogLevelDefault   LogLevel = iota
	LogLevelErrorOnly          = 1
	LogLevelStandard           = 2
	LogLevelElevated           = 3
	LogLevelFull               = 4
	LogLevelDebug              = 5
	LogLevelTrace              = 6
	LogLevelVerbose            = 7
)

type SybmlInit struct {
	LogLevel      LogLevel
	DebugFilePath string
}

func Init(init SybmlInit) {
	if init.LogLevel == LogLevelDefault {
		init.LogLevel = LogLevelStandard
	}

	klog.InitFlags(nil)
	flag.Set("v", strconv.FormatInt(int64(init.LogLevel), 10))
	if init.DebugFilePath != "" {
		flag.Set("logtostderr", "false")
		flag.Set("log_file", init.DebugFilePath)
	}
	flag.Parse()
}
