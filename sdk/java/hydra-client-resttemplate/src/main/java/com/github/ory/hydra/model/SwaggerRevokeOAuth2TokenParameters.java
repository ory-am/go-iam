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
 * SwaggerRevokeOAuth2TokenParameters
 */
@javax.annotation.Generated(value = "io.swagger.codegen.languages.JavaClientCodegen", date = "2019-05-08T10:14:06.817+02:00")
public class SwaggerRevokeOAuth2TokenParameters {
  @JsonProperty("token")
  private String token = null;

  public SwaggerRevokeOAuth2TokenParameters token(String token) {
    this.token = token;
    return this;
  }

   /**
   * in: formData
   * @return token
  **/
  @ApiModelProperty(required = true, value = "in: formData")
  public String getToken() {
    return token;
  }

  public void setToken(String token) {
    this.token = token;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SwaggerRevokeOAuth2TokenParameters swaggerRevokeOAuth2TokenParameters = (SwaggerRevokeOAuth2TokenParameters) o;
    return Objects.equals(this.token, swaggerRevokeOAuth2TokenParameters.token);
  }

  @Override
  public int hashCode() {
    return Objects.hash(token);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SwaggerRevokeOAuth2TokenParameters {\n");
    
    sb.append("    token: ").append(toIndentedString(token)).append("\n");
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

