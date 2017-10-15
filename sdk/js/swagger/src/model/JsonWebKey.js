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
    define(['ApiClient'], factory)
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'))
  } else {
    // Browser globals (root is window)
    if (!root.HydraOAuth2OpenIdConnectServer) {
      root.HydraOAuth2OpenIdConnectServer = {}
    }
    root.HydraOAuth2OpenIdConnectServer.JsonWebKey = factory(
      root.HydraOAuth2OpenIdConnectServer.ApiClient
    )
  }
})(this, function(ApiClient) {
  'use strict'

  /**
   * The JsonWebKey model module.
   * @module model/JsonWebKey
   * @version Latest
   */

  /**
   * Constructs a new <code>JsonWebKey</code>.
   * @alias module:model/JsonWebKey
   * @class
   */
  var exports = function() {
    var _this = this
  }

  /**
   * Constructs a <code>JsonWebKey</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/JsonWebKey} obj Optional instance to populate.
   * @return {module:model/JsonWebKey} The populated <code>JsonWebKey</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports()

      if (data.hasOwnProperty('alg')) {
        obj['alg'] = ApiClient.convertToType(data['alg'], 'String')
      }
      if (data.hasOwnProperty('crv')) {
        obj['crv'] = ApiClient.convertToType(data['crv'], 'String')
      }
      if (data.hasOwnProperty('d')) {
        obj['d'] = ApiClient.convertToType(data['d'], 'String')
      }
      if (data.hasOwnProperty('dp')) {
        obj['dp'] = ApiClient.convertToType(data['dp'], 'String')
      }
      if (data.hasOwnProperty('dq')) {
        obj['dq'] = ApiClient.convertToType(data['dq'], 'String')
      }
      if (data.hasOwnProperty('e')) {
        obj['e'] = ApiClient.convertToType(data['e'], 'String')
      }
      if (data.hasOwnProperty('k')) {
        obj['k'] = ApiClient.convertToType(data['k'], 'String')
      }
      if (data.hasOwnProperty('kid')) {
        obj['kid'] = ApiClient.convertToType(data['kid'], 'String')
      }
      if (data.hasOwnProperty('kty')) {
        obj['kty'] = ApiClient.convertToType(data['kty'], 'String')
      }
      if (data.hasOwnProperty('n')) {
        obj['n'] = ApiClient.convertToType(data['n'], 'String')
      }
      if (data.hasOwnProperty('p')) {
        obj['p'] = ApiClient.convertToType(data['p'], 'String')
      }
      if (data.hasOwnProperty('q')) {
        obj['q'] = ApiClient.convertToType(data['q'], 'String')
      }
      if (data.hasOwnProperty('qi')) {
        obj['qi'] = ApiClient.convertToType(data['qi'], 'String')
      }
      if (data.hasOwnProperty('use')) {
        obj['use'] = ApiClient.convertToType(data['use'], 'String')
      }
      if (data.hasOwnProperty('x')) {
        obj['x'] = ApiClient.convertToType(data['x'], 'String')
      }
      if (data.hasOwnProperty('x5c')) {
        obj['x5c'] = ApiClient.convertToType(data['x5c'], ['String'])
      }
      if (data.hasOwnProperty('y')) {
        obj['y'] = ApiClient.convertToType(data['y'], 'String')
      }
    }
    return obj
  }

  /**
   * The \"alg\" (algorithm) parameter identifies the algorithm intended for use with the key.  The values used should either be registered in the IANA \"JSON Web Signature and Encryption Algorithms\" registry established by [JWA] or be a value that contains a Collision- Resistant Name.
   * @member {String} alg
   */
  exports.prototype['alg'] = undefined
  /**
   * @member {String} crv
   */
  exports.prototype['crv'] = undefined
  /**
   * @member {String} d
   */
  exports.prototype['d'] = undefined
  /**
   * @member {String} dp
   */
  exports.prototype['dp'] = undefined
  /**
   * @member {String} dq
   */
  exports.prototype['dq'] = undefined
  /**
   * @member {String} e
   */
  exports.prototype['e'] = undefined
  /**
   * @member {String} k
   */
  exports.prototype['k'] = undefined
  /**
   * The \"kid\" (key ID) parameter is used to match a specific key.  This is used, for instance, to choose among a set of keys within a JWK Set during key rollover.  The structure of the \"kid\" value is unspecified.  When \"kid\" values are used within a JWK Set, different keys within the JWK Set SHOULD use distinct \"kid\" values.  (One example in which different keys might use the same \"kid\" value is if they have different \"kty\" (key type) values but are considered to be equivalent alternatives by the application using them.)  The \"kid\" value is a case-sensitive string.
   * @member {String} kid
   */
  exports.prototype['kid'] = undefined
  /**
   * The \"kty\" (key type) parameter identifies the cryptographic algorithm family used with the key, such as \"RSA\" or \"EC\". \"kty\" values should either be registered in the IANA \"JSON Web Key Types\" registry established by [JWA] or be a value that contains a Collision- Resistant Name.  The \"kty\" value is a case-sensitive string.
   * @member {String} kty
   */
  exports.prototype['kty'] = undefined
  /**
   * @member {String} n
   */
  exports.prototype['n'] = undefined
  /**
   * @member {String} p
   */
  exports.prototype['p'] = undefined
  /**
   * @member {String} q
   */
  exports.prototype['q'] = undefined
  /**
   * @member {String} qi
   */
  exports.prototype['qi'] = undefined
  /**
   * The \"use\" (public key use) parameter identifies the intended use of the public key. The \"use\" parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Values are commonly \"sig\" (signature) or \"enc\" (encryption).
   * @member {String} use
   */
  exports.prototype['use'] = undefined
  /**
   * @member {String} x
   */
  exports.prototype['x'] = undefined
  /**
   * The \"x5c\" (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates [RFC5280].  The certificate chain is represented as a JSON array of certificate value strings.  Each string in the array is a base64-encoded (Section 4 of [RFC4648] -- not base64url-encoded) DER [ITU.X690.1994] PKIX certificate value. The PKIX certificate containing the key value MUST be the first certificate.
   * @member {Array.<String>} x5c
   */
  exports.prototype['x5c'] = undefined
  /**
   * @member {String} y
   */
  exports.prototype['y'] = undefined

  return exports
})
