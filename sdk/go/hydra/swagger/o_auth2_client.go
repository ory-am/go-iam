/*
 * Hydra OAuth2 & OpenID Connect Server
 *
 * Please refer to the user guide for in-depth documentation: https://ory.gitbooks.io/hydra/content/   Hydra offers OAuth 2.0 and OpenID Connect Core 1.0 capabilities as a service. Hydra is different, because it works with any existing authentication infrastructure, not just LDAP or SAML. By implementing a consent app (works with any programming language) you build a bridge between Hydra and your authentication infrastructure. Hydra is able to securely manage JSON Web Keys, and has a sophisticated policy-based access control you can use if you want to. Hydra is suitable for green- (new) and brownfield (existing) projects. If you are not familiar with OAuth 2.0 and are working on a greenfield project, we recommend evaluating if OAuth 2.0 really serves your purpose. Knowledge of OAuth 2.0 is imperative in understanding what Hydra does and how it works.   The official repository is located at https://github.com/ory/hydra   ### Important REST API Documentation Notes  The swagger generator used to create this documentation does currently not support example responses. To see request and response payloads click on **\"Show JSON schema\"**: ![Enable JSON Schema on Apiary](https://storage.googleapis.com/ory.am/hydra/json-schema.png)   The API documentation always refers to the latest tagged version of ORY Hydra. For previous API documentations, please refer to https://github.com/ory/hydra/blob/<tag-id>/docs/api.swagger.yaml - for example:  0.9.13: https://github.com/ory/hydra/blob/v0.9.13/docs/api.swagger.yaml 0.8.1: https://github.com/ory/hydra/blob/v0.8.1/docs/api.swagger.yaml
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type OAuth2Client struct {

	// Name is the human-readable string name of the client to be presented to the end-user during authorization.
	ClientName string `json:"client_name,omitempty"`

	// Secret is the client's secret. The secret will be included in the create request as cleartext, and then never again. The secret is stored using BCrypt so it is impossible to recover it. Tell your users that they need to write the secret down as it will not be made available again.
	ClientSecret string `json:"client_secret,omitempty"`

	// ClientURI is an URL string of a web page providing information about the client. If present, the server SHOULD display this URL to the end-user in a clickable fashion.
	ClientUri string `json:"client_uri,omitempty"`

	// Contacts is a array of strings representing ways to contact people responsible for this client, typically email addresses.
	Contacts []string `json:"contacts,omitempty"`

	// GrantTypes is an array of grant types the client is allowed to use.
	GrantTypes []string `json:"grant_types,omitempty"`

	// ID is the id for this client.
	Id string `json:"id,omitempty"`

	// LogoURI is an URL string that references a logo for the client.
	LogoUri string `json:"logo_uri,omitempty"`

	// Owner is a string identifying the owner of the OAuth 2.0 Client.
	Owner string `json:"owner,omitempty"`

	// PolicyURI is a URL string that points to a human-readable privacy policy document that describes how the deployment organization collects, uses, retains, and discloses personal data.
	PolicyUri string `json:"policy_uri,omitempty"`

	// Public is a boolean that identifies this client as public, meaning that it does not have a secret. It will disable the client_credentials grant type for this client if set.
	Public bool `json:"public,omitempty"`

	// RedirectURIs is an array of allowed redirect urls for the client, for example: http://mydomain/oauth/callback .
	RedirectUris []string `json:"redirect_uris,omitempty"`

	// ResponseTypes is an array of the OAuth 2.0 response type strings that the client can use at the authorization endpoint.
	ResponseTypes []string `json:"response_types,omitempty"`

	// Scope is a string containing a space-separated list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client can use when requesting access tokens.
	Scope string `json:"scope,omitempty"`

	// TermsOfServiceURI is a URL string that points to a human-readable terms of service document for the client that describes a contractual relationship between the end-user and the client that the end-user accepts when authorizing the client.
	TosUri string `json:"tos_uri,omitempty"`
}
