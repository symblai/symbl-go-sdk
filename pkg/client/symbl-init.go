// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package symbl

import (
	"flag"
	"strconv"

	klog "k8s.io/klog/v2"
)

func InitLogging(debugLevel int) {
	if debugLevel == 0 && debugLevel > 6 {
		debugLevel = 2
	}

	klog.InitFlags(nil)
	flag.Set("v", strconv.FormatInt(int64(debugLevel), 10))
	flag.Parse()
}

func InitLoggingToFile(debugLevel int) {
	if debugLevel == 0 && debugLevel > 6 {
		debugLevel = 2
	}

	klog.InitFlags(nil)
	flag.Set("logtostderr", "false")
	flag.Set("log_file", "symbl-go-sdk.log")
	flag.Set("v", strconv.FormatInt(int64(debugLevel), 10))
	flag.Parse()
}
