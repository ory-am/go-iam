module.exports = {
  "Introduction": [
    "index",
    "5min-tutorial",
    "install",
  ],
  Concepts: [
    "concepts/before-oauth2",
    "concepts/oauth2",
    "concepts/openid-connect-oidc",
    "concepts/login",
    "concepts/consent",
    "concepts/logout",
    "jwks",
    "limitations",
  ],
  "Guides": [
    {
      type: "category",
      label: "Implementing the User Interface",
      items: [
        "guides/login",
        "guides/consent",
        "guides/logout",
      ]
    },
    {
      type: "category",
      label: "Operations",
      items: [
        "configure-deploy",
        "dependencies-environment",
        "production",
        "guides/tracing",
        "guides/secrets-key-rotation",
        "guides/kubernetes-helm-chart",
        "guides/ssl-https-tls",
        "guides/cookies",
        "guides/scaling-hydra",
        "guides/cors",
      ]
    },
    {
      type: "category",
      label: "OAuth2 & OpenID Connect",
      items: [
        "advanced",
        "guides/oauth2-clients",
        "guides/common-oauth2-openid-connect-flows",
        "guides/using-oauth2",
        "guides/token-expiration",
        "guides/oauth2-token-introspection",
        "guides/oauth2-public-spa-mobile"
      ]
    },
  ],
  "Reference": [
    "reference/configuration",
    "reference/api",
  ],
  "Debug & Help": [
    "debug",
    "debug/csrf",
    "debug/token-endpoint-auth-method",
    "debug/logout",
  ],
  "SDKs": [
    "sdk",
    "sdk/go",
    "sdk/js",
    "sdk/php",
  ],
  "Further Reading": [
    "case-study",
    "benchmark",
    "security-architecture",
    "faq"
  ],
};
