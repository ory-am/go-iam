/**
 * ORY Hydra - Cloud Native OAuth 2.0 and OpenID Connect Server
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here. Keep in mind that this document reflects the latest branch, always. Support for versioned documentation is coming in the future.
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
    if (!root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer) {
      root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer = {}
    }
    root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer.ConsentRequestManager = factory(
      root.OryHydraCloudNativeOAuth20AndOpenIdConnectServer.ApiClient
    )
  }
})(this, function(ApiClient) {
  'use strict'

  /**
   * The ConsentRequestManager model module.
   * @module model/ConsentRequestManager
   * @version Latest
   */

  /**
   * Constructs a new <code>ConsentRequestManager</code>.
   * @alias module:model/ConsentRequestManager
   * @class
   */
  var exports = function() {
    var _this = this
  }

  /**
   * Constructs a <code>ConsentRequestManager</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ConsentRequestManager} obj Optional instance to populate.
   * @return {module:model/ConsentRequestManager} The populated <code>ConsentRequestManager</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports()
    }
    return obj
  }

  return exports
})
