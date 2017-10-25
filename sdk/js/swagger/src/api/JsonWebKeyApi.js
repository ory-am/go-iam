/**
 * Hydra OAuth2 & OpenID Connect Server
 * Please refer to the user guide for in-depth documentation: https://ory.gitbooks.io/hydra/content/   Hydra offers OAuth 2.0 and OpenID Connect Core 1.0 capabilities as a service. Hydra is different, because it works with any existing authentication infrastructure, not just LDAP or SAML. By implementing a consent app (works with any programming language) you build a bridge between Hydra and your authentication infrastructure. Hydra is able to securely manage JSON Web Keys, and has a sophisticated policy-based access control you can use if you want to. Hydra is suitable for green- (new) and brownfield (existing) projects. If you are not familiar with OAuth 2.0 and are working on a greenfield project, we recommend evaluating if OAuth 2.0 really serves your purpose. Knowledge of OAuth 2.0 is imperative in understanding what Hydra does and how it works.   The official repository is located at https://github.com/ory/hydra   ### Important REST API Documentation Notes  The swagger generator used to create this documentation does currently not support example responses. To see request and response payloads click on **\"Show JSON schema\"**: ![Enable JSON Schema on Apiary](https://storage.googleapis.com/ory.am/hydra/json-schema.png)   The API documentation always refers to the latest tagged version of ORY Hydra. For previous API documentations, please refer to https://github.com/ory/hydra/blob/<tag-id>/docs/api.swagger.yaml - for example:  0.9.13: https://github.com/ory/hydra/blob/v0.9.13/docs/api.swagger.yaml 0.8.1: https://github.com/ory/hydra/blob/v0.8.1/docs/api.swagger.yaml
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

;(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(
      [
        'ApiClient',
        'model/InlineResponse401',
        'model/JsonWebKey',
        'model/JsonWebKeySet',
        'model/JsonWebKeySetGeneratorRequest'
      ],
      factory
    )
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(
      require('../ApiClient'),
      require('../model/InlineResponse401'),
      require('../model/JsonWebKey'),
      require('../model/JsonWebKeySet'),
      require('../model/JsonWebKeySetGeneratorRequest')
    )
  } else {
    // Browser globals (root is window)
    if (!root.HydraOAuth2OpenIdConnectServer) {
      root.HydraOAuth2OpenIdConnectServer = {}
    }
    root.HydraOAuth2OpenIdConnectServer.JsonWebKeyApi = factory(
      root.HydraOAuth2OpenIdConnectServer.ApiClient,
      root.HydraOAuth2OpenIdConnectServer.InlineResponse401,
      root.HydraOAuth2OpenIdConnectServer.JsonWebKey,
      root.HydraOAuth2OpenIdConnectServer.JsonWebKeySet,
      root.HydraOAuth2OpenIdConnectServer.JsonWebKeySetGeneratorRequest
    )
  }
})(this, function(
  ApiClient,
  InlineResponse401,
  JsonWebKey,
  JsonWebKeySet,
  JsonWebKeySetGeneratorRequest
) {
  'use strict'

  /**
   * JsonWebKey service.
   * @module api/JsonWebKeyApi
   * @version Latest
   */

  /**
   * Constructs a new JsonWebKeyApi. 
   * @alias module:api/JsonWebKeyApi
   * @class
   * @param {module:ApiClient} apiClient Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance

    /**
     * Callback function to receive the result of the createJsonWebKeySet operation.
     * @callback module:api/JsonWebKeyApi~createJsonWebKeySetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/JsonWebKeySet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Generate a new JSON Web Key
     * This endpoint is capable of generating JSON Web Key Sets for you. There a different strategies available, such as symmetric cryptographic keys (HS256, HS512) and asymetric cryptographic keys (RS256, ECDSA).   If the specified JSON Web Key Set does not exist, it will be created.   The subject making the request needs to be assigned to a policy containing:  &#x60;&#x60;&#x60; { \&quot;resources\&quot;: [\&quot;rn:hydra:keys:&lt;set&gt;:&lt;kid&gt;\&quot;], \&quot;actions\&quot;: [\&quot;create\&quot;], \&quot;effect\&quot;: \&quot;allow\&quot; } &#x60;&#x60;&#x60;
     * @param {String} set The set
     * @param {Object} opts Optional parameters
     * @param {module:model/JsonWebKeySetGeneratorRequest} opts.body 
     * @param {module:api/JsonWebKeyApi~createJsonWebKeySetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/JsonWebKeySet}
     */
    this.createJsonWebKeySet = function(set, opts, callback) {
      opts = opts || {}
      var postBody = opts['body']

      // verify the required parameter 'set' is set
      if (set === undefined || set === null) {
        throw new Error(
          "Missing the required parameter 'set' when calling createJsonWebKeySet"
        )
      }

      var pathParams = {
        set: set
      }
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = ['oauth2']
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = JsonWebKeySet

      return this.apiClient.callApi(
        '/keys/{set}',
        'POST',
        pathParams,
        queryParams,
        headerParams,
        formParams,
        postBody,
        authNames,
        contentTypes,
        accepts,
        returnType,
        callback
      )
    }

    /**
     * Callback function to receive the result of the deleteJsonWebKey operation.
     * @callback module:api/JsonWebKeyApi~deleteJsonWebKeyCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete a JSON Web Key
     * The subject making the request needs to be assigned to a policy containing:  &#x60;&#x60;&#x60; { \&quot;resources\&quot;: [\&quot;rn:hydra:keys:&lt;set&gt;:&lt;kid&gt;\&quot;], \&quot;actions\&quot;: [\&quot;delete\&quot;], \&quot;effect\&quot;: \&quot;allow\&quot; } &#x60;&#x60;&#x60;
     * @param {String} kid The kid of the desired key
     * @param {String} set The set
     * @param {module:api/JsonWebKeyApi~deleteJsonWebKeyCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteJsonWebKey = function(kid, set, callback) {
      var postBody = null

      // verify the required parameter 'kid' is set
      if (kid === undefined || kid === null) {
        throw new Error(
          "Missing the required parameter 'kid' when calling deleteJsonWebKey"
        )
      }

      // verify the required parameter 'set' is set
      if (set === undefined || set === null) {
        throw new Error(
          "Missing the required parameter 'set' when calling deleteJsonWebKey"
        )
      }

      var pathParams = {
        kid: kid,
        set: set
      }
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = ['oauth2']
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = null

      return this.apiClient.callApi(
        '/keys/{set}/{kid}',
        'DELETE',
        pathParams,
        queryParams,
        headerParams,
        formParams,
        postBody,
        authNames,
        contentTypes,
        accepts,
        returnType,
        callback
      )
    }

    /**
     * Callback function to receive the result of the deleteJsonWebKeySet operation.
     * @callback module:api/JsonWebKeyApi~deleteJsonWebKeySetCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete a JSON Web Key
     * The subject making the request needs to be assigned to a policy containing:  &#x60;&#x60;&#x60; { \&quot;resources\&quot;: [\&quot;rn:hydra:keys:&lt;set&gt;\&quot;], \&quot;actions\&quot;: [\&quot;delete\&quot;], \&quot;effect\&quot;: \&quot;allow\&quot; } &#x60;&#x60;&#x60;
     * @param {String} set The set
     * @param {module:api/JsonWebKeyApi~deleteJsonWebKeySetCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteJsonWebKeySet = function(set, callback) {
      var postBody = null

      // verify the required parameter 'set' is set
      if (set === undefined || set === null) {
        throw new Error(
          "Missing the required parameter 'set' when calling deleteJsonWebKeySet"
        )
      }

      var pathParams = {
        set: set
      }
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = ['oauth2']
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = null

      return this.apiClient.callApi(
        '/keys/{set}',
        'DELETE',
        pathParams,
        queryParams,
        headerParams,
        formParams,
        postBody,
        authNames,
        contentTypes,
        accepts,
        returnType,
        callback
      )
    }

    /**
     * Callback function to receive the result of the getJsonWebKey operation.
     * @callback module:api/JsonWebKeyApi~getJsonWebKeyCallback
     * @param {String} error Error message, if any.
     * @param {module:model/JsonWebKeySet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieve a JSON Web Key
     * This endpoint can be used to retrieve JWKs stored in ORY Hydra.   The subject making the request needs to be assigned to a policy containing:  &#x60;&#x60;&#x60; { \&quot;resources\&quot;: [\&quot;rn:hydra:keys:&lt;set&gt;:&lt;kid&gt;\&quot;], \&quot;actions\&quot;: [\&quot;get\&quot;], \&quot;effect\&quot;: \&quot;allow\&quot; } &#x60;&#x60;&#x60;
     * @param {String} kid The kid of the desired key
     * @param {String} set The set
     * @param {module:api/JsonWebKeyApi~getJsonWebKeyCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/JsonWebKeySet}
     */
    this.getJsonWebKey = function(kid, set, callback) {
      var postBody = null

      // verify the required parameter 'kid' is set
      if (kid === undefined || kid === null) {
        throw new Error(
          "Missing the required parameter 'kid' when calling getJsonWebKey"
        )
      }

      // verify the required parameter 'set' is set
      if (set === undefined || set === null) {
        throw new Error(
          "Missing the required parameter 'set' when calling getJsonWebKey"
        )
      }

      var pathParams = {
        kid: kid,
        set: set
      }
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = ['oauth2']
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = JsonWebKeySet

      return this.apiClient.callApi(
        '/keys/{set}/{kid}',
        'GET',
        pathParams,
        queryParams,
        headerParams,
        formParams,
        postBody,
        authNames,
        contentTypes,
        accepts,
        returnType,
        callback
      )
    }

    /**
     * Callback function to receive the result of the getJsonWebKeySet operation.
     * @callback module:api/JsonWebKeyApi~getJsonWebKeySetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/JsonWebKeySet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Retrieve a JSON Web Key Set
     * This endpoint can be used to retrieve JWK Sets stored in ORY Hydra.   The subject making the request needs to be assigned to a policy containing:  &#x60;&#x60;&#x60; { \&quot;resources\&quot;: [\&quot;rn:hydra:keys:&lt;set&gt;:&lt;kid&gt;\&quot;], \&quot;actions\&quot;: [\&quot;get\&quot;], \&quot;effect\&quot;: \&quot;allow\&quot; } &#x60;&#x60;&#x60;
     * @param {String} set The set
     * @param {module:api/JsonWebKeyApi~getJsonWebKeySetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/JsonWebKeySet}
     */
    this.getJsonWebKeySet = function(set, callback) {
      var postBody = null

      // verify the required parameter 'set' is set
      if (set === undefined || set === null) {
        throw new Error(
          "Missing the required parameter 'set' when calling getJsonWebKeySet"
        )
      }

      var pathParams = {
        set: set
      }
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = ['oauth2']
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = JsonWebKeySet

      return this.apiClient.callApi(
        '/keys/{set}',
        'GET',
        pathParams,
        queryParams,
        headerParams,
        formParams,
        postBody,
        authNames,
        contentTypes,
        accepts,
        returnType,
        callback
      )
    }

    /**
     * Callback function to receive the result of the updateJsonWebKey operation.
     * @callback module:api/JsonWebKeyApi~updateJsonWebKeyCallback
     * @param {String} error Error message, if any.
     * @param {module:model/JsonWebKey} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update a JSON Web Key
     * Use this method if you do not want to let Hydra generate the JWKs for you, but instead save your own.   The subject making the request needs to be assigned to a policy containing:  &#x60;&#x60;&#x60; { \&quot;resources\&quot;: [\&quot;rn:hydra:keys:&lt;set&gt;:&lt;kid&gt;\&quot;], \&quot;actions\&quot;: [\&quot;update\&quot;], \&quot;effect\&quot;: \&quot;allow\&quot; } &#x60;&#x60;&#x60;
     * @param {String} kid The kid of the desired key
     * @param {String} set The set
     * @param {Object} opts Optional parameters
     * @param {module:model/JsonWebKey} opts.body 
     * @param {module:api/JsonWebKeyApi~updateJsonWebKeyCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/JsonWebKey}
     */
    this.updateJsonWebKey = function(kid, set, opts, callback) {
      opts = opts || {}
      var postBody = opts['body']

      // verify the required parameter 'kid' is set
      if (kid === undefined || kid === null) {
        throw new Error(
          "Missing the required parameter 'kid' when calling updateJsonWebKey"
        )
      }

      // verify the required parameter 'set' is set
      if (set === undefined || set === null) {
        throw new Error(
          "Missing the required parameter 'set' when calling updateJsonWebKey"
        )
      }

      var pathParams = {
        kid: kid,
        set: set
      }
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = ['oauth2']
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = JsonWebKey

      return this.apiClient.callApi(
        '/keys/{set}/{kid}',
        'PUT',
        pathParams,
        queryParams,
        headerParams,
        formParams,
        postBody,
        authNames,
        contentTypes,
        accepts,
        returnType,
        callback
      )
    }

    /**
     * Callback function to receive the result of the updateJsonWebKeySet operation.
     * @callback module:api/JsonWebKeyApi~updateJsonWebKeySetCallback
     * @param {String} error Error message, if any.
     * @param {module:model/JsonWebKeySet} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update a JSON Web Key Set
     * Use this method if you do not want to let Hydra generate the JWKs for you, but instead save your own.   The subject making the request needs to be assigned to a policy containing:  &#x60;&#x60;&#x60; { \&quot;resources\&quot;: [\&quot;rn:hydra:keys:&lt;set&gt;\&quot;], \&quot;actions\&quot;: [\&quot;update\&quot;], \&quot;effect\&quot;: \&quot;allow\&quot; } &#x60;&#x60;&#x60;
     * @param {String} set The set
     * @param {Object} opts Optional parameters
     * @param {module:model/JsonWebKeySet} opts.body 
     * @param {module:api/JsonWebKeyApi~updateJsonWebKeySetCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/JsonWebKeySet}
     */
    this.updateJsonWebKeySet = function(set, opts, callback) {
      opts = opts || {}
      var postBody = opts['body']

      // verify the required parameter 'set' is set
      if (set === undefined || set === null) {
        throw new Error(
          "Missing the required parameter 'set' when calling updateJsonWebKeySet"
        )
      }

      var pathParams = {
        set: set
      }
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = ['oauth2']
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = JsonWebKeySet

      return this.apiClient.callApi(
        '/keys/{set}',
        'PUT',
        pathParams,
        queryParams,
        headerParams,
        formParams,
        postBody,
        authNames,
        contentTypes,
        accepts,
        returnType,
        callback
      )
    }
  }

  return exports
})
