// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package async

import (
	"context"
	"fmt"

	klog "k8s.io/klog/v2"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
)

func (c *Client) getQueryParamFromContext(ctx context.Context) string {
	// additional query parameters to URL
	params := make(map[string][]string, 0)

	if parameters, ok := ctx.Value(interfaces.ParametersContext{}).(map[string][]string); ok {
		for k, vs := range parameters {
			klog.V(5).Infof("Key/Value: %s = %v\n", k, vs)
			params[k] = vs
		}
	}

	if len(params) > 0 {
		queryString := ""
		for k, vs := range params {
			for _, v := range vs {
				if len(queryString) > 1 {
					queryString += "&"
				}
				queryString += fmt.Sprintf("%s=%s", k, v)
			}
		}
		klog.V(5).Infof("Final Query String: %s\n", queryString)
		return queryString
	}

	klog.V(6).Infof("Final Query String is Empty\n")
	return ""
}
