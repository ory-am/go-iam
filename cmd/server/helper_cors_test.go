/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @Copyright 	2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/fosite"
	"github.com/ory/hydra/client"
	"github.com/ory/hydra/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCORSMiddleware(t *testing.T) {
	handler := httprouter.New()
	handler.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusNoContent)
	})

	c := new(config.Config)
	for k, tc := range []struct {
		d            string
		mw           func(http.Handler) http.Handler
		code         int
		header       http.Header
		expectHeader http.Header
	}{
		{
			d:            "should ignore when disabled",
			mw:           newCORSMiddleware(false, nil, nil, nil),
			code:         204,
			header:       http.Header{},
			expectHeader: http.Header{},
		},
		{
			d: "should reject when basic auth but client does not exist",
			mw: newCORSMiddleware(true, c, nil, func(id string) (*client.Client, error) {
				return nil, errors.New("")
			}),
			code:         204,
			header:       http.Header{"Origin": {"http://foobar.com"}, "Authorization": {"Basic Zm9vOmJhcg=="}},
			expectHeader: http.Header{"Vary": {"Origin"}},
		},
		{
			d: "should reject when basic auth client exists but origin not allowed",
			mw: newCORSMiddleware(true, c, nil, func(id string) (*client.Client, error) {
				return &client.Client{AllowedCORSOrigins: []string{"http://not-foobar.com"}}, nil
			}),
			code:         204,
			header:       http.Header{"Origin": {"http://foobar.com"}, "Authorization": {"Basic Zm9vOmJhcg=="}},
			expectHeader: http.Header{"Vary": {"Origin"}},
		},
		{
			d: "should accept when basic auth client exists and origin allowed",
			mw: newCORSMiddleware(true, c, nil, func(id string) (*client.Client, error) {
				return &client.Client{AllowedCORSOrigins: []string{"http://foobar.com"}}, nil
			}),
			code:         204,
			header:       http.Header{"Origin": {"http://foobar.com"}, "Authorization": {"Basic Zm9vOmJhcg=="}},
			expectHeader: http.Header{"Vary": {"Origin"}, "Access-Control-Allow-Origin": {"http://foobar.com"}},
		},
		{
			d: "should fail when token introspection fails",
			mw: newCORSMiddleware(true, c, func(ctx context.Context, token string, tokenType fosite.TokenType, session fosite.Session, scope ...string) (fosite.TokenType, fosite.AccessRequester, error) {
				return "", nil, errors.New("")
			}, func(id string) (*client.Client, error) {
				return &client.Client{AllowedCORSOrigins: []string{"http://foobar.com"}}, nil
			}),
			code:         204,
			header:       http.Header{"Origin": {"http://foobar.com"}, "Authorization": {"Basic Zm9vOmJhcg=="}},
			expectHeader: http.Header{"Vary": {"Origin"}, "Access-Control-Allow-Origin": {"http://foobar.com"}},
		},
		{
			d: "should fail when token introspection fails",
			mw: newCORSMiddleware(true, c, func(ctx context.Context, token string, tokenType fosite.TokenType, session fosite.Session, scope ...string) (fosite.TokenType, fosite.AccessRequester, error) {
				if token != "1234" {
					return "", nil, errors.New("")
				}
				return "", &fosite.AccessRequest{Request: fosite.Request{Client: &client.Client{ClientID: "asdf"}}}, nil
			}, func(id string) (*client.Client, error) {
				if id != "asdf" {
					return nil, errors.New("")
				}
				return &client.Client{AllowedCORSOrigins: []string{"http://foobar.com"}}, nil
			}),
			code:         204,
			header:       http.Header{"Origin": {"http://foobar.com"}, "Authorization": {"bearer 1234"}},
			expectHeader: http.Header{"Vary": {"Origin"}, "Access-Control-Allow-Origin": {"http://foobar.com"}},
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://foobar.com/", nil)
			require.NoError(t, err)
			for k := range tc.header {
				req.Header.Set(k, tc.header.Get(k))
			}

			res := httptest.NewRecorder()
			tc.mw(handler).ServeHTTP(res, req)
			require.NoError(t, err)
			assert.EqualValues(t, tc.expectHeader, res.Header())
		})
	}
}
