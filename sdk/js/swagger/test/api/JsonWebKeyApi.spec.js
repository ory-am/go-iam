/**
 * Hydra OAuth2 & OpenID Connect Server
 * Please refer to the user guide for in-depth documentation: https://ory.gitbooks.io/hydra/content/   Hydra offers OAuth 2.0 and OpenID Connect Core 1.0 capabilities as a service. Hydra is different, because it works with any existing authentication infrastructure, not just LDAP or SAML. By implementing a consent app (works with any programming language) you build a bridge between Hydra and your authentication infrastructure. Hydra is able to securely manage JSON Web Keys, and has a sophisticated policy-based access control you can use if you want to. Hydra is suitable for green- (new) and brownfield (existing) projects. If you are not familiar with OAuth 2.0 and are working on a greenfield project, we recommend evaluating if OAuth 2.0 really serves your purpose. Knowledge of OAuth 2.0 is imperative in understanding what Hydra does and how it works.   The official repository is located at https://github.com/ory/hydra   ### ATTENTION - IMPORTANT NOTE   The swagger generator used to create this documentation does currently not support example responses. To see request and response payloads click on **\"Show JSON schema\"**: ![Enable JSON Schema on Apiary](https://storage.googleapis.com/ory.am/hydra/json-schema.png)
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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.HydraOAuth2OpenIdConnectServer);
  }
}(this, function(expect, HydraOAuth2OpenIdConnectServer) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new HydraOAuth2OpenIdConnectServer.JsonWebKeyApi();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('JsonWebKeyApi', function() {
    describe('createJsonWebKeySet', function() {
      it('should call createJsonWebKeySet successfully', function(done) {
        //uncomment below and update the code to test createJsonWebKeySet
        //instance.createJsonWebKeySet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteJsonWebKey', function() {
      it('should call deleteJsonWebKey successfully', function(done) {
        //uncomment below and update the code to test deleteJsonWebKey
        //instance.deleteJsonWebKey(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteJsonWebKeySet', function() {
      it('should call deleteJsonWebKeySet successfully', function(done) {
        //uncomment below and update the code to test deleteJsonWebKeySet
        //instance.deleteJsonWebKeySet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getJsonWebKey', function() {
      it('should call getJsonWebKey successfully', function(done) {
        //uncomment below and update the code to test getJsonWebKey
        //instance.getJsonWebKey(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getJsonWebKeySet', function() {
      it('should call getJsonWebKeySet successfully', function(done) {
        //uncomment below and update the code to test getJsonWebKeySet
        //instance.getJsonWebKeySet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateJsonWebKey', function() {
      it('should call updateJsonWebKey successfully', function(done) {
        //uncomment below and update the code to test updateJsonWebKey
        //instance.updateJsonWebKey(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateJsonWebKeySet', function() {
      it('should call updateJsonWebKeySet successfully', function(done) {
        //uncomment below and update the code to test updateJsonWebKeySet
        //instance.updateJsonWebKeySet(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
