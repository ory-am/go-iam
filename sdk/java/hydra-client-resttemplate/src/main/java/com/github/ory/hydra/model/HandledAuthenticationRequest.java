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
 * HandledAuthenticationRequest
 */
@javax.annotation.Generated(value = "io.swagger.codegen.languages.JavaClientCodegen", date = "2019-04-20T10:32:55.035+02:00")
public class HandledAuthenticationRequest {
  @JsonProperty("acr")
  private String acr = null;

  @JsonProperty("force_subject_identifier")
  private String forceSubjectIdentifier = null;

  @JsonProperty("remember")
  private Boolean remember = null;

  @JsonProperty("remember_for")
  private Long rememberFor = null;

  @JsonProperty("subject")
  private String subject = null;

  public HandledAuthenticationRequest acr(String acr) {
    this.acr = acr;
    return this;
  }

   /**
   * ACR sets the Authentication AuthorizationContext Class Reference value for this authentication session. You can use it to express that, for example, a user authenticated using two factor authentication.
   * @return acr
  **/
  @ApiModelProperty(value = "ACR sets the Authentication AuthorizationContext Class Reference value for this authentication session. You can use it to express that, for example, a user authenticated using two factor authentication.")
  public String getAcr() {
    return acr;
  }

  public void setAcr(String acr) {
    this.acr = acr;
  }

  public HandledAuthenticationRequest forceSubjectIdentifier(String forceSubjectIdentifier) {
    this.forceSubjectIdentifier = forceSubjectIdentifier;
    return this;
  }

   /**
   * ForceSubjectIdentifier forces the \&quot;pairwise\&quot; user ID of the end-user that authenticated. The \&quot;pairwise\&quot; user ID refers to the (Pairwise Identifier Algorithm)[http://openid.net/specs/openid-connect-core-1_0.html#PairwiseAlg] of the OpenID Connect specification. It allows you to set an obfuscated subject (\&quot;user\&quot;) identifier that is unique to the client.  Please note that this changes the user ID on endpoint /userinfo and sub claim of the ID Token. It does not change the sub claim in the OAuth 2.0 Introspection.  Per default, ORY Hydra handles this value with its own algorithm. In case you want to set this yourself you can use this field. Please note that setting this field has no effect if &#x60;pairwise&#x60; is not configured in ORY Hydra or the OAuth 2.0 Client does not expect a pairwise identifier (set via &#x60;subject_type&#x60; key in the client&#39;s configuration).  Please also be aware that ORY Hydra is unable to properly compute this value during authentication. This implies that you have to compute this value on every authentication process (probably depending on the client ID or some other unique value).  If you fail to compute the proper value, then authentication processes which have id_token_hint set might fail.
   * @return forceSubjectIdentifier
  **/
  @ApiModelProperty(value = "ForceSubjectIdentifier forces the \"pairwise\" user ID of the end-user that authenticated. The \"pairwise\" user ID refers to the (Pairwise Identifier Algorithm)[http://openid.net/specs/openid-connect-core-1_0.html#PairwiseAlg] of the OpenID Connect specification. It allows you to set an obfuscated subject (\"user\") identifier that is unique to the client.  Please note that this changes the user ID on endpoint /userinfo and sub claim of the ID Token. It does not change the sub claim in the OAuth 2.0 Introspection.  Per default, ORY Hydra handles this value with its own algorithm. In case you want to set this yourself you can use this field. Please note that setting this field has no effect if `pairwise` is not configured in ORY Hydra or the OAuth 2.0 Client does not expect a pairwise identifier (set via `subject_type` key in the client's configuration).  Please also be aware that ORY Hydra is unable to properly compute this value during authentication. This implies that you have to compute this value on every authentication process (probably depending on the client ID or some other unique value).  If you fail to compute the proper value, then authentication processes which have id_token_hint set might fail.")
  public String getForceSubjectIdentifier() {
    return forceSubjectIdentifier;
  }

  public void setForceSubjectIdentifier(String forceSubjectIdentifier) {
    this.forceSubjectIdentifier = forceSubjectIdentifier;
  }

  public HandledAuthenticationRequest remember(Boolean remember) {
    this.remember = remember;
    return this;
  }

   /**
   * Remember, if set to true, tells ORY Hydra to remember this user by telling the user agent (browser) to store a cookie with authentication data. If the same user performs another OAuth 2.0 Authorization Request, he/she will not be asked to log in again.
   * @return remember
  **/
  @ApiModelProperty(value = "Remember, if set to true, tells ORY Hydra to remember this user by telling the user agent (browser) to store a cookie with authentication data. If the same user performs another OAuth 2.0 Authorization Request, he/she will not be asked to log in again.")
  public Boolean getRemember() {
    return remember;
  }

  public void setRemember(Boolean remember) {
    this.remember = remember;
  }

  public HandledAuthenticationRequest rememberFor(Long rememberFor) {
    this.rememberFor = rememberFor;
    return this;
  }

   /**
   * RememberFor sets how long the authentication should be remembered for in seconds. If set to &#x60;0&#x60;, the authorization will be remembered indefinitely.
   * @return rememberFor
  **/
  @ApiModelProperty(value = "RememberFor sets how long the authentication should be remembered for in seconds. If set to `0`, the authorization will be remembered indefinitely.")
  public Long getRememberFor() {
    return rememberFor;
  }

  public void setRememberFor(Long rememberFor) {
    this.rememberFor = rememberFor;
  }

  public HandledAuthenticationRequest subject(String subject) {
    this.subject = subject;
    return this;
  }

   /**
   * Subject is the user ID of the end-user that authenticated.
   * @return subject
  **/
  @ApiModelProperty(required = true, value = "Subject is the user ID of the end-user that authenticated.")
  public String getSubject() {
    return subject;
  }

  public void setSubject(String subject) {
    this.subject = subject;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HandledAuthenticationRequest handledAuthenticationRequest = (HandledAuthenticationRequest) o;
    return Objects.equals(this.acr, handledAuthenticationRequest.acr) &&
        Objects.equals(this.forceSubjectIdentifier, handledAuthenticationRequest.forceSubjectIdentifier) &&
        Objects.equals(this.remember, handledAuthenticationRequest.remember) &&
        Objects.equals(this.rememberFor, handledAuthenticationRequest.rememberFor) &&
        Objects.equals(this.subject, handledAuthenticationRequest.subject);
  }

  @Override
  public int hashCode() {
    return Objects.hash(acr, forceSubjectIdentifier, remember, rememberFor, subject);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HandledAuthenticationRequest {\n");
    
    sb.append("    acr: ").append(toIndentedString(acr)).append("\n");
    sb.append("    forceSubjectIdentifier: ").append(toIndentedString(forceSubjectIdentifier)).append("\n");
    sb.append("    remember: ").append(toIndentedString(remember)).append("\n");
    sb.append("    rememberFor: ").append(toIndentedString(rememberFor)).append("\n");
    sb.append("    subject: ").append(toIndentedString(subject)).append("\n");
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

