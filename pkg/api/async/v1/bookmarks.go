// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// Use of this source code is governed by an Apache-2.0 license that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

/*
	Async package for processing Async conversations
*/
package async

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	asyncinterfaces "github.com/symblai/symbl-go-sdk/pkg/api/async/v1/interfaces"
	version "github.com/symblai/symbl-go-sdk/pkg/api/version"
	interfaces "github.com/symblai/symbl-go-sdk/pkg/client/interfaces"
)

// GetBookmarks to get bookmarks of a conversation
func (c *Client) GetBookmarks(ctx context.Context, conversationId string) (*asyncinterfaces.BookmarksResult, error) {
	klog.V(6).Infof("async.GetBookmarks ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetBookmarks LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetManagementAPI(version.BookmarksURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetBookmarks LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.BookmarksResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetBookmarks LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetBookmarks LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET Bookmarks succeeded\n")
	klog.V(6).Infof("async.GetBookmarks LEAVE\n")
	return &result, nil
}

// GetBookmarkById get bookmarks by ID
func (c *Client) GetBookmarkById(ctx context.Context, conversationId, bookmarkId string) (*asyncinterfaces.BookmarksResult, error) {
	klog.V(6).Infof("async.GetBookmarkById ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
		return nil, ErrInvalidInput
	}
	if bookmarkId == "" {
		klog.V(1).Infof("bookmarkId is empty\n")
		klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetManagementAPI(version.BookmarksByIdURI, conversationId, bookmarkId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.BookmarksResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET BookmarkById succeeded\n")
	klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
	return &result, nil
}

/*
	CreateBookmark creates a bookmark in a conversation

	TODO:
	When exercising the API and description is blank...

	HTTP Code: 400
	{
		"message":"\"description\" is not allowed to be empty"
	}
*/
func (c *Client) CreateBookmark(ctx context.Context, conversationId string, request asyncinterfaces.BookmarkRequest) (*asyncinterfaces.Bookmark, error) {
	klog.V(6).Infof("async.CreateBookmark ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	v := validator.New()
	err := v.Struct(request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("CreateBookmark validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, err
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetManagementAPI(version.BookmarksURI, conversationId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.Bookmark

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.CreateBookmark LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET Create Bookmark succeeded\n")
	klog.V(6).Infof("async.CreateBookmark LEAVE\n")
	return &result, nil
}

// UpdateBookmark updates an existing bookmark in a conversation
func (c *Client) UpdateBookmark(ctx context.Context, conversationId, bookmarkId string, request asyncinterfaces.BookmarkRequest) (*asyncinterfaces.Bookmark, error) {
	klog.V(6).Infof("async.UpdateBookmark ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	v := validator.New()
	err := v.Struct(request)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			klog.V(1).Infof("UpdateBookmark validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
		return nil, err
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
		return nil, ErrInvalidInput
	}
	if bookmarkId == "" {
		klog.V(1).Infof("bookmarkId is empty\n")
		klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetManagementAPI(version.BookmarksByIdURI, conversationId, bookmarkId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.V(1).Infof("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.Bookmark

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET Update Bookmark succeeded\n")
	klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
	return &result, nil
}

// DeleteBookmark removes a bookmark in a conversation
func (c *Client) DeleteBookmark(ctx context.Context, conversationId, bookmarkId string) error {
	klog.V(6).Infof("async.DeleteBookmark ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
		return ErrInvalidInput
	}
	if bookmarkId == "" {
		klog.V(1).Infof("bookmarkId is empty\n")
		klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
		return ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetManagementAPI(version.BookmarksByIdURI, conversationId, bookmarkId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "DELETE", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
		return err
	}

	// check the status
	err = c.Client.Do(ctx, req, nil)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
				return err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
		return err
	}

	klog.V(3).Infof("GET Delete Bookmark succeeded\n")
	klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
	return nil
}

// GetSummaryOfBookmark gets a summary of bookmarks
func (c *Client) GetSummaryOfBookmark(ctx context.Context, conversationId, bookmarkId string) (*asyncinterfaces.BookmarkSummaryResult, error) {
	klog.V(6).Infof("async.GetSummaryOfBookmark ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetSummaryOfBookmark LEAVE\n")
		return nil, ErrInvalidInput
	}
	if bookmarkId == "" {
		klog.V(1).Infof("bookmarkId is empty\n")
		klog.V(6).Infof("async.GetSummaryOfBookmark LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := fmt.Sprintf("%s%s",
		version.GetAsyncAPI(version.BookmarkSummaryURI, conversationId, bookmarkId),
		c.getQueryParamFromContext(ctx))
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetSummaryOfBookmark LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.BookmarkSummaryResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetSummaryOfBookmark LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetSummaryOfBookmark LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET SummaryOfBookmark succeeded\n")
	klog.V(6).Infof("async.GetSummaryOfBookmark LEAVE\n")
	return &result, nil
}

// GetSummaryOfBookmarks gets a list of bookmarks in a given conversation
func (c *Client) GetSummaryOfBookmarks(ctx context.Context, conversationId string, filters []string) (*asyncinterfaces.BookmarksSummaryResult, error) {
	klog.V(6).Infof("async.GetSummaryOfBookmarks ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.V(1).Infof("conversationId is empty\n")
		klog.V(6).Infof("async.GetSummaryOfBookmarks LEAVE\n")
		return nil, ErrInvalidInput
	}

	queryString := ""
	if len(filters) > 0 {
		queryString = "?"
		for _, filter := range filters {
			queryString += url.QueryEscape(filter)
		}
	}

	// request
	URI := version.GetAsyncAPI(version.SummariesOfBookmarksURI, conversationId)
	if len(filters) > 0 {
		URI = version.GetAsyncAPI(version.SummariesOfBookmarksURI, conversationId, queryString)
	}
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.V(1).Infof("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetSummaryOfBookmarks LEAVE\n")
		return nil, err
	}

	// check the status
	var result asyncinterfaces.BookmarksSummaryResult

	err = c.Client.Do(ctx, req, &result)

	if err != nil {
		if e, ok := err.(*interfaces.StatusError); ok {
			if e.Resp.StatusCode != http.StatusOK {
				klog.V(1).Infof("HTTP Code: %v\n", e.Resp.StatusCode)
				klog.V(6).Infof("async.GetSummaryOfBookmarks LEAVE\n")
				return nil, err
			}
		}

		klog.V(1).Infof("Platform Supplied Err: %v\n", err)
		klog.V(6).Infof("async.GetSummaryOfBookmarks LEAVE\n")
		return nil, err
	}

	klog.V(3).Infof("GET SummaryOfBookmarks succeeded\n")
	klog.V(6).Infof("async.GetSummaryOfBookmarks LEAVE\n")
	return &result, nil
}
