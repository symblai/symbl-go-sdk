// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package symbl

import (
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
	UUID            string
	SymblEndpoint   string
	SymblConfig     *cfginterfaces.StreamingConfig
	Callback        rtinterfaces.InsightCallback
	SkipServerAuth  bool
	RedirectService bool
}

type StreamClient struct {
	*stream.WebSocketClient

	uuid           string
	restClient     *RestClient
	symblStreaming stream.WebSocketMessageCallback

	options *StreamingOptions
}
