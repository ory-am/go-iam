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
 * Swaggeroauth2TokenResponse The Access Token Response
 */
@ApiModel(description = "Swaggeroauth2TokenResponse The Access Token Response")
@javax.annotation.Generated(value = "io.swagger.codegen.languages.JavaClientCodegen", date = "2019-04-11T18:32:06.000+02:00")
public class Swaggeroauth2TokenResponse {
  @JsonProperty("access_token")
  private String accessToken = null;

  @JsonProperty("client_id")
  private String clientId = null;

  @JsonProperty("code")
  private String code = null;

  @JsonProperty("redirect_uri")
  private String redirectUri = null;

  public Swaggeroauth2TokenResponse accessToken(String accessToken) {
    this.accessToken = accessToken;
    return this;
  }

   /**
   * access token
   * @return accessToken
  **/
  @ApiModelProperty(value = "access token")
  public String getAccessToken() {
    return accessToken;
  }

  public void setAccessToken(String accessToken) {
    this.accessToken = accessToken;
  }

  public Swaggeroauth2TokenResponse clientId(String clientId) {
    this.clientId = clientId;
    return this;
  }

   /**
   * refresh token
   * @return clientId
  **/
  @ApiModelProperty(value = "refresh token")
  public String getClientId() {
    return clientId;
  }

  public void setClientId(String clientId) {
    this.clientId = clientId;
  }

  public Swaggeroauth2TokenResponse code(String code) {
    this.code = code;
    return this;
  }

   /**
   * token type
   * @return code
  **/
  @ApiModelProperty(value = "token type")
  public String getCode() {
    return code;
  }

  public void setCode(String code) {
    this.code = code;
  }

  public Swaggeroauth2TokenResponse redirectUri(String redirectUri) {
    this.redirectUri = redirectUri;
    return this;
  }

   /**
   * expires in
   * @return redirectUri
  **/
  @ApiModelProperty(value = "expires in")
  public String getRedirectUri() {
    return redirectUri;
  }

  public void setRedirectUri(String redirectUri) {
    this.redirectUri = redirectUri;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Swaggeroauth2TokenResponse swaggeroauth2TokenResponse = (Swaggeroauth2TokenResponse) o;
    return Objects.equals(this.accessToken, swaggeroauth2TokenResponse.accessToken) &&
        Objects.equals(this.clientId, swaggeroauth2TokenResponse.clientId) &&
        Objects.equals(this.code, swaggeroauth2TokenResponse.code) &&
        Objects.equals(this.redirectUri, swaggeroauth2TokenResponse.redirectUri);
  }

  @Override
  public int hashCode() {
    return Objects.hash(accessToken, clientId, code, redirectUri);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Swaggeroauth2TokenResponse {\n");
    
    sb.append("    accessToken: ").append(toIndentedString(accessToken)).append("\n");
    sb.append("    clientId: ").append(toIndentedString(clientId)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    redirectUri: ").append(toIndentedString(redirectUri)).append("\n");
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

