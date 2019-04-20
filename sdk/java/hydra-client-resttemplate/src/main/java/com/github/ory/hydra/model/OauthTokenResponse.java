/*
 * ORY Hydra
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package com.github.ory.hydra.model;

import java.util.Objects;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

/**
 * The token response
 */
@ApiModel(description = "The token response")
@javax.annotation.Generated(value = "io.swagger.codegen.languages.JavaClientCodegen", date = "2019-04-20T10:32:55.035+02:00")
public class OauthTokenResponse {
  @JsonProperty("access_token")
  private String accessToken = null;

  @JsonProperty("expires_in")
  private Long expiresIn = null;

  @JsonProperty("id_token")
  private Long idToken = null;

  @JsonProperty("refresh_token")
  private String refreshToken = null;

  @JsonProperty("scope")
  private Long scope = null;

  @JsonProperty("token_type")
  private String tokenType = null;

  public OauthTokenResponse accessToken(String accessToken) {
    this.accessToken = accessToken;
    return this;
  }

   /**
   * The access token issued by the authorization server.
   * @return accessToken
  **/
  @ApiModelProperty(value = "The access token issued by the authorization server.")
  public String getAccessToken() {
    return accessToken;
  }

  public void setAccessToken(String accessToken) {
    this.accessToken = accessToken;
  }

  public OauthTokenResponse expiresIn(Long expiresIn) {
    this.expiresIn = expiresIn;
    return this;
  }

   /**
   * The lifetime in seconds of the access token.  For example, the value \&quot;3600\&quot; denotes that the access token will expire in one hour from the time the response was generated.
   * @return expiresIn
  **/
  @ApiModelProperty(value = "The lifetime in seconds of the access token.  For example, the value \"3600\" denotes that the access token will expire in one hour from the time the response was generated.")
  public Long getExpiresIn() {
    return expiresIn;
  }

  public void setExpiresIn(Long expiresIn) {
    this.expiresIn = expiresIn;
  }

  public OauthTokenResponse idToken(Long idToken) {
    this.idToken = idToken;
    return this;
  }

   /**
   * To retrieve a refresh token request the id_token scope.
   * @return idToken
  **/
  @ApiModelProperty(value = "To retrieve a refresh token request the id_token scope.")
  public Long getIdToken() {
    return idToken;
  }

  public void setIdToken(Long idToken) {
    this.idToken = idToken;
  }

  public OauthTokenResponse refreshToken(String refreshToken) {
    this.refreshToken = refreshToken;
    return this;
  }

   /**
   * The refresh token, which can be used to obtain new access tokens. To retrieve it add the scope \&quot;offline\&quot; to your access token request.
   * @return refreshToken
  **/
  @ApiModelProperty(value = "The refresh token, which can be used to obtain new access tokens. To retrieve it add the scope \"offline\" to your access token request.")
  public String getRefreshToken() {
    return refreshToken;
  }

  public void setRefreshToken(String refreshToken) {
    this.refreshToken = refreshToken;
  }

  public OauthTokenResponse scope(Long scope) {
    this.scope = scope;
    return this;
  }

   /**
   * The scope of the access token
   * @return scope
  **/
  @ApiModelProperty(value = "The scope of the access token")
  public Long getScope() {
    return scope;
  }

  public void setScope(Long scope) {
    this.scope = scope;
  }

  public OauthTokenResponse tokenType(String tokenType) {
    this.tokenType = tokenType;
    return this;
  }

   /**
   * The type of the token issued
   * @return tokenType
  **/
  @ApiModelProperty(value = "The type of the token issued")
  public String getTokenType() {
    return tokenType;
  }

  public void setTokenType(String tokenType) {
    this.tokenType = tokenType;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    OauthTokenResponse oauthTokenResponse = (OauthTokenResponse) o;
    return Objects.equals(this.accessToken, oauthTokenResponse.accessToken) &&
        Objects.equals(this.expiresIn, oauthTokenResponse.expiresIn) &&
        Objects.equals(this.idToken, oauthTokenResponse.idToken) &&
        Objects.equals(this.refreshToken, oauthTokenResponse.refreshToken) &&
        Objects.equals(this.scope, oauthTokenResponse.scope) &&
        Objects.equals(this.tokenType, oauthTokenResponse.tokenType);
  }

  @Override
  public int hashCode() {
    return Objects.hash(accessToken, expiresIn, idToken, refreshToken, scope, tokenType);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class OauthTokenResponse {\n");
    
    sb.append("    accessToken: ").append(toIndentedString(accessToken)).append("\n");
    sb.append("    expiresIn: ").append(toIndentedString(expiresIn)).append("\n");
    sb.append("    idToken: ").append(toIndentedString(idToken)).append("\n");
    sb.append("    refreshToken: ").append(toIndentedString(refreshToken)).append("\n");
    sb.append("    scope: ").append(toIndentedString(scope)).append("\n");
    sb.append("    tokenType: ").append(toIndentedString(tokenType)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
  
}

