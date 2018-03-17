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
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package oauth2_test

import (
	"net/http/httptest"
	"net/url"
	"time"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/herodot"
	hc "github.com/ory/hydra/client"
	. "github.com/ory/hydra/oauth2"
	"github.com/ory/hydra/pkg"
	hydra "github.com/ory/hydra/sdk/go/hydra/swagger"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var hasher = &fosite.BCrypt{}

var store = &FositeMemoryStore{
	Manager:        hc.NewMemoryManager(hasher),
	AuthorizeCodes: make(map[string]fosite.Requester),
	IDSessions:     make(map[string]fosite.Requester),
	AccessTokens:   make(map[string]fosite.Requester),
	RefreshTokens:  make(map[string]fosite.Requester),
}

var fc = &compose.Config{
	AccessTokenLifespan: time.Second,
}

var handler = &Handler{
	OAuth2: compose.Compose(
		fc,
		store,
		&compose.CommonStrategy{
			CoreStrategy:               compose.NewOAuth2HMACStrategy(fc, []byte("some super secret secret secret secret")),
			OpenIDConnectTokenStrategy: compose.NewOpenIDConnectStrategy(pkg.MustINSECURELOWENTROPYRSAKEYFORTEST()),
		},
		nil,
		compose.OAuth2AuthorizeExplicitFactory,
		compose.OAuth2AuthorizeImplicitFactory,
		compose.OAuth2ClientCredentialsGrantFactory,
		compose.OAuth2RefreshTokenGrantFactory,
		compose.OpenIDConnectExplicitFactory,
		compose.OpenIDConnectHybridFactory,
		compose.OpenIDConnectImplicitFactory,
		compose.OAuth2TokenRevocationFactory,
		compose.OAuth2TokenIntrospectionFactory,
	),
	Consent: &DefaultConsentStrategy{
		Issuer:                   "http://hydra.localhost",
		ConsentManager:           consentManager,
		DefaultChallengeLifespan: time.Hour,
		DefaultIDTokenLifespan:   time.Hour * 24,
	},
	CookieStore:   sessions.NewCookieStore([]byte("foo-secret")),
	ForcedHTTP:    true,
	L:             logrus.New(),
	ScopeStrategy: fosite.HierarchicScopeStrategy,
	H:             herodot.NewJSONWriter(nil),
}

var router = httprouter.New()
var ts *httptest.Server
var oauthConfig *oauth2.Config
var oauthClientConfig *clientcredentials.Config

var consentHandler *ConsentSessionHandler
var consentManager = NewConsentRequestMemoryManager()
var consentClient *hydra.OAuth2Api

func init() {
	consentHandler = &ConsentSessionHandler{
		H: herodot.NewJSONWriter(nil),
		M: consentManager,
	}

	ts = httptest.NewServer(router)
	handler.Issuer = ts.URL

	handler.SetRoutes(router)
	consentHandler.SetRoutes(router)

	h, _ := hasher.Hash([]byte("secret"))
	consentClient = hydra.NewOAuth2ApiWithBasePath(ts.URL)

	c, _ := url.Parse(ts.URL + "/consent")
	handler.ConsentURL = *c

	store.Manager.(*hc.MemoryManager).Clients = append(store.Manager.(*hc.MemoryManager).Clients, hc.Client{
		ID:            "app-client",
		Secret:        string(h),
		RedirectURIs:  []string{ts.URL + "/callback"},
		ResponseTypes: []string{"id_token", "code", "token"},
		GrantTypes:    []string{"implicit", "refresh_token", "authorization_code", "password", "client_credentials"},
		Scope:         "hydra.* offline openid",
	})

	oauthConfig = &oauth2.Config{
		ClientID:     "app-client",
		ClientSecret: "secret",
		Endpoint: oauth2.Endpoint{
			AuthURL:  ts.URL + "/oauth2/auth",
			TokenURL: ts.URL + "/oauth2/token",
		},
		RedirectURL: ts.URL + "/callback",
		Scopes:      []string{"hydra.*", "offline", "openid"},
	}

	oauthClientConfig = &clientcredentials.Config{
		ClientID:     "app-client",
		ClientSecret: "secret",
		TokenURL:     ts.URL + "/oauth2/token",
		Scopes:       []string{"hydra.consent", "offline"},
	}
}
