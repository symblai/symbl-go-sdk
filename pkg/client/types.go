// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package symbl

import (
	rtinterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/streaming/v1/interfaces"
	cfginterfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/client/interfaces"
	rest "github.com/dvonthenen/symbl-go-sdk/pkg/client/rest"
	stream "github.com/dvonthenen/symbl-go-sdk/pkg/client/stream"
)

// RestClient extends the pkg/client/rest Client and also keeps tabs on the
// Creds and Auth Cookie in order to reconnect
type RestClient struct {
	*rest.Client

	creds *interfaces.Credentials
	auth  *interfaces.AuthResp
}

// StreamingOptions are connection options for the Real-Time Websocket client
type StreamingOptions struct {
	UUID            string
	SymblEndpoint   string
	SymblConfig     *cfginterfaces.StreamingConfig
	Callback        rtinterfaces.InsightCallback
	SkipServerAuth  bool
	RedirectService bool
}

// StreamClient is a representation of the Symbl Platform streaming client over a Websocket interface
type StreamClient struct {
	*stream.WebSocketClient

	uuid           string
	restClient     *RestClient
	symblStreaming stream.WebSocketMessageCallback

	options *StreamingOptions
}

// NebulaClient extends the pkg/client/rest Client and also keeps tabs on the auth token
type NebulaClient struct {
	*rest.Client

	creds *interfaces.Credentials
	auth  *interfaces.AuthResp
}
