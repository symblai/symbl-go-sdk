// Copyright 2023 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Nebula package for processing Nebula Async conversations
*/
package nebula

import (
	"context"
	"fmt"

	klog "k8s.io/klog/v2"

	interfaces "github.com/symblai/symbl-go-sdk/pkg/client/interfaces"
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

	// TODO: replace with https://github.com/google/go-querystring
	// API differs from how go-querystring works
	//
	//	go-querystring : []vals -> key=vals[0]&key=vals[1]
	//	symbl API: []vals -> key=["$vals[0]", "$vals[1]"]
	//
	// need to look into switching that behavior in order to use go-querystring lib
	if len(params) > 1 {
		queryString := "?"
		for k, vs := range params {
			if len(queryString) > 3 {
				queryString += "&"
			}
			if len(vs) == 1 {
				queryString += fmt.Sprintf("%s=%s", k, vs[0])
			} else {
				appended := false
				for _, v := range vs {
					if !appended {
						queryString += fmt.Sprintf("%s=[", k)
						appended = true
					} else {
						queryString += ","
					}
					queryString += fmt.Sprintf("%s", v)
				}
				if len(vs) > 0 {
					queryString += "]"
				}
			}
		}
		klog.V(5).Infof("Final Query String: %s\n", queryString)
		return queryString
	}

	klog.V(6).Infof("Final Query String is Empty\n")
	return ""
}
