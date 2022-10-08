// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package async

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"
	klog "k8s.io/klog/v2"

	version "github.com/dvonthenen/symbl-go-sdk/pkg/api/version"
	symbl "github.com/dvonthenen/symbl-go-sdk/pkg/client"

	interfaces "github.com/dvonthenen/symbl-go-sdk/pkg/api/async/v1/interfaces"
)

func (c *Client) GetBookmarks(ctx context.Context, conversationId string) (*interfaces.BookmarksResult, error) {
	klog.V(6).Infof("async.GetBookmarks ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetBookmarks LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetManagementAPI(version.BookmarksURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetBookmarks LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.BookmarksResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.Errorf("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetBookmarks LEAVE\n")
			return nil, err
		}
	}

	klog.V(4).Infof("GET Bookmarks succeeded\n")
	klog.V(6).Infof("async.GetBookmarks LEAVE\n")
	return &result, nil
}

func (c *Client) GetBookmarkById(ctx context.Context, conversationId, bookmarkId string) (*interfaces.BookmarksResult, error) {
	klog.V(6).Infof("async.GetBookmarkById ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
		return nil, ErrInvalidInput
	}
	if bookmarkId == "" {
		klog.Errorf("bookmarkId is empty\n")
		klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetManagementAPI(version.BookmarksByIdURI, conversationId, bookmarkId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "GET", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.BookmarksResult

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.Errorf("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
			return nil, err
		}
	}

	klog.V(4).Infof("GET BookmarkById succeeded\n")
	klog.V(6).Infof("async.GetBookmarkById LEAVE\n")
	return &result, nil
}

/*
	TODO: description is required
	HTTP Code: 400
	{
		"message":"\"description\" is not allowed to be empty"
	}

	If using MessageRefs, then BeginTimeOffset and Duration cannot be zero (aka present in the struct)
	which contradicts the docs: https://docs.symbl.ai/docs/create-bookmarks-guide
	The reverse is also true... using BeginTimeOffset and Duration, MessageRefs must not be present in the struct
*/
func (c *Client) CreateBookmarkByMessageRefs(ctx context.Context, conversationId string, request interfaces.BookmarkByMessageRefsRequest) (*interfaces.Bookmark, error) {
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
			klog.Errorf("CreateBookmark validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, err
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetManagementAPI(version.BookmarksURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	// let's auth
	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.Errorf("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.Bookmark

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.Errorf("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.CreateBookmark LEAVE\n")
			return nil, err
		}
	}

	klog.V(4).Infof("GET Create Bookmark succeeded\n")
	klog.V(6).Infof("async.CreateBookmark LEAVE\n")
	return &result, nil
}
func (c *Client) CreateBookmarkByTimeDuration(ctx context.Context, conversationId string, request interfaces.BookmarkBtTimeDurationsRequest) (*interfaces.Bookmark, error) {
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
			klog.Errorf("CreateBookmark validation failed. Err: %v\n", e)
		}
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, err
	}
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, ErrInvalidInput
	}

	// request
	URI := version.GetManagementAPI(version.BookmarksURI, conversationId)
	klog.V(6).Infof("Calling %s\n", URI)

	// let's auth
	jsonStr, err := json.Marshal(request)
	if err != nil {
		klog.Errorf("json.Marshal failed. Err: %v\n", err)
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", URI, bytes.NewBuffer(jsonStr))
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.CreateBookmark LEAVE\n")
		return nil, err
	}

	// check the status
	var result interfaces.Bookmark

	err = c.Client.Do(ctx, req, &result)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.Errorf("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.CreateBookmark LEAVE\n")
			return nil, err
		}
	}

	klog.V(4).Infof("GET Create Bookmark succeeded\n")
	klog.V(6).Infof("async.CreateBookmark LEAVE\n")
	return &result, nil
}

/*
	??? TODO: This appears broken... This is the error we get back:

	HTTP Code: 400
	{
		"message":"Either create/update bookmark using 'beginTimeOffset + duration' or 'messageRefs'. None of them provided"
	}

	Even if we only supply MessageRefs OR beginTimeOffset + duration
*/
// func (c *Client) UpdateBookmark(ctx context.Context, conversationId, bookmarkId string, request interfaces.BookmarkByMessageRefsRequest) (*interfaces.Bookmark, error) {
// 	klog.V(6).Infof("async.UpdateBookmark ENTER\n")

// 	// checks
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}

// 	// validate input
// 	v := validator.New()
// 	err := v.Struct(request)
// 	if err != nil {
// 		for _, e := range err.(validator.ValidationErrors) {
// 			klog.Errorf("UpdateBookmark validation failed. Err: %v\n", e)
// 		}
// 		klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
// 		return nil, err
// 	}
// 	if conversationId == "" {
// 		klog.Errorf("conversationId is empty\n")
// 		klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
// 		return nil, ErrInvalidInput
// 	}
// 	if bookmarkId == "" {
// 		klog.Errorf("bookmarkId is empty\n")
// 		klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
// 		return nil, ErrInvalidInput
// 	}

// 	// request
// 	URI := version.GetManagementAPI(version.BookmarksByIdURI, conversationId, bookmarkId)
// 	klog.V(6).Infof("Calling %s\n", URI)

// 	req, err := http.NewRequestWithContext(ctx, "PUT", URI, nil)
// 	if err != nil {
// 		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
// 		klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
// 		return nil, err
// 	}

// 	// check the status
// 	var result interfaces.Bookmark

// 	err = c.Client.Do(ctx, req, &result)

// 	if e, ok := err.(*symbl.StatusError); ok {
// 		if e.Resp.StatusCode != http.StatusOK {
// 			klog.Errorf("HTTP Code: %v\n", e.Resp.StatusCode)
// 			klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
// 			return nil, err
// 		}
// 	}

// 	klog.V(4).Infof("GET Update Bookmark succeeded\n")
// 	klog.V(6).Infof("async.UpdateBookmark LEAVE\n")
// 	return &result, nil
// }

func (c *Client) DeleteBookmark(ctx context.Context, conversationId, bookmarkId string) error {
	klog.V(6).Infof("async.DeleteBookmark ENTER\n")

	// checks
	if ctx == nil {
		ctx = context.Background()
	}

	// validate input
	if conversationId == "" {
		klog.Errorf("conversationId is empty\n")
		klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
		return ErrInvalidInput
	}
	if bookmarkId == "" {
		klog.Errorf("bookmarkId is empty\n")
		klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
		return ErrInvalidInput
	}

	// request
	URI := version.GetManagementAPI(version.BookmarksByIdURI, conversationId, bookmarkId)
	klog.V(6).Infof("Calling %s\n", URI)

	req, err := http.NewRequestWithContext(ctx, "DELETE", URI, nil)
	if err != nil {
		klog.Errorf("http.NewRequestWithContext failed. Err: %v\n", err)
		klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
		return err
	}

	// check the status
	err = c.Client.Do(ctx, req, nil)

	if e, ok := err.(*symbl.StatusError); ok {
		if e.Resp.StatusCode != http.StatusOK {
			klog.Errorf("HTTP Code: %v\n", e.Resp.StatusCode)
			klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
			return err
		}
	}

	klog.V(4).Infof("GET Delete Bookmark succeeded\n")
	klog.V(6).Infof("async.DeleteBookmark LEAVE\n")
	return nil
}
