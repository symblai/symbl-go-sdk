// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package symbl

import (
	"fmt"

	rest "github.com/dvonthenen/symbl-go-sdk/pkg/client/rest"
)

/*
	Symbl REST API
*/
type HeadersContext struct{}

type StatusError struct {
	*rest.StatusError
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Resp.Request.Method, e.Resp.Request.URL, e.Resp.Status)
}
