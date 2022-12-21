// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package symbl

import (
	"fmt"

	rtinterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/streaming/v1/interfaces"
	cfginterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
	rest "github.com/dvonthenen/symbl-go-sdk/pkg/client/rest"
	stream "github.com/dvonthenen/symbl-go-sdk/pkg/client/stream"
)

/*
	REST Client
*/
type RestClient struct {
	*rest.Client

	creds *interfaces.Credentials
	auth  *interfaces.AuthResp
}

/*
	Streaming Client
*/
type StreamingOptions struct {
	UUID           string
	ProxyAddress   string
	SymblConfig    *cfginterfaces.StreamingConfig
	Callback       rtinterfaces.InsightCallback
	SkipServerAuth bool
}

type StreamClient struct {
	*stream.WebSocketClient

	uuid           string
	restClient     *RestClient
	symblStreaming stream.WebSocketMessageCallback

	options *StreamingOptions
}

/*
	Symbl REST API Internals
*/
type HeadersContext struct{}

type StatusError struct {
	*rest.StatusError
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Resp.Request.Method, e.Resp.Request.URL, e.Resp.Status)
}
