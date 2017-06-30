<h1 align="center"><img src="docs/images/logo.png" alt="ORY Hydra"></h1>

<p align="center">
    <a href="https://gitter.im/ory-am/hydra"><img src="https://img.shields.io/badge/join-ory%20chat-00cc99.svg" alt="Join the chat at https://gitter.im/ory-am/hydra"></a>
    <a href="https://community.ory.am/"><img src="https://img.shields.io/badge/join-ory%20community-00cc99.svg" alt="Join the community at https://community.ory.am/"></a>
    <a href="http://eepurl.com/bKT3N9"><img src="https://img.shields.io/badge/join-ory%20newsletter-00cc99.svg" alt="Join newsletter"></a>
    <a href="https://patreon.com/user?u=4298803"><img src="https://img.shields.io/badge/support%20us-on%20patreon-green.svg" alt="Become a patron!"></a>
</p>
<p align="center">
    <a href="https://travis-ci.org/ory/hydra"><img src="https://travis-ci.org/ory/hydra.svg?branch=master" alt="Build Status"></a>
    <a href="https://coveralls.io/github/ory/hydra?branch=master"><img src="https://coveralls.io/repos/ory/hydra/badge.svg?branch=master&service=github" alt="Coverage Status"></a>
    <a href="https://codeclimate.com/github/ory/hydra"><img src="https://codeclimate.com/github/ory/hydra/badges/gpa.svg" alt="Code Climate"></a>
    <a href="https://goreportcard.com/report/github.com/ory/hydra"><img src="https://goreportcard.com/badge/github.com/ory/hydra" alt="Go Report Card"></a>
    <a href="https://bestpractices.coreinfrastructure.org/projects/364"><img src="https://bestpractices.coreinfrastructure.org/projects/364/badge" alt="CII Best Practices"></a>
</p>

<p align="center">
    <a href="https://ory.gitbooks.io/hydra/content/"><img src="https://img.shields.io/badge/docs-guide-blue.svg" alt="Docs Guide"></a>
    <a href="http://docs.hydra13.apiary.io/"><img src="https://img.shields.io/badge/docs-http%20api-blue.svg" alt="HTTP API Documentation"></a>
    <a href="https://godoc.org/github.com/ory/hydra"><img src="https://img.shields.io/badge/docs-godoc-blue.svg" alt="Code Documentation"></a>
</p>

ORY Hydra is a security-first OAuth 2.0 and OpenID Connect server optimized for low-latency, high throughput,
and low resource consumption. ORY Hydra *is not* an identity provider (user sign up, user log in, password reset flow),
but connects to your existing identity provider through a [consent app](https://ory.gitbooks.io/hydra/content/oauth2.html#consent-app-flow).
Implementing the consent app in a different language is easy, and exemplary consent apps
([Go](https://github.com/ory/hydra-consent-app-go), [Node](https://github.com/ory/hydra-consent-app-express)) and
SDKs ([Go](https://github.com/ory/hydra/tree/update-docs/sdk), [Node](https://github.com/ory/hydra-js)) are provided.

Besides mitigating various attack vectors, such as database compromisation and OAuth 2.0 weaknesses, ORY Hydra is
able to securely manage JSON Web Keys, and has a sophisticated policy-based access control you can use if you want to.
[Click here](https://ory.gitbooks.io/hydra/content/security.html#security-overview) to read more about security.

---

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [What is ORY Hydra?](#what-is-ory-hydra)
  - [ORY Hydra implements open standards](#ory-hydra-implements-open-standards)
- [ORY Cloud Security](#ory-cloud-security)
  - [Products](#products)
  - [Consulting](#consulting)
- [Quickstart](#quickstart)
  - [5 minutes tutorial: Run your very own OAuth2 environment](#5-minutes-tutorial-run-your-very-own-oauth2-environment)
  - [Installation](#installation)
    - [Download binaries](#download-binaries)
    - [Using Docker](#using-docker)
    - [Building from source](#building-from-source)
- [Security](#security)
- [Telemetry](#telemetry)
- [Sponsors & Adopters](#sponsors-&-adopters)
  - [Sponsorship](#sponsorship)
  - [Sponsors](#sponsors)
  - [Adopters](#adopters)
- [Documentation](#documentation)
  - [Guide](#guide)
  - [HTTP API documentation](#http-api-documentation)
  - [Command line documentation](#command-line-documentation)
  - [Develop](#develop)
- [Reception](#reception)
- [Libraries and third-party projects](#libraries-and-third-party-projects)
- [Blog posts & articles](#blog-posts-&-articles)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## What is ORY Hydra?

ORY Hydra is a server implementation of the OAuth 2.0 authorization framework and the OpenID Connect Core 1.0. Existing OAuth2
implementations usually ship as libraries or SDKs such as [node-oauth2-server](https://github.com/oauthjs/node-oauth2-server)
or [fosite](https://github.com/ory/fosite/issues), or as fully featured identity solutions with user
management and user interfaces, such as [Dex](https://github.com/coreos/dex).

Implementing and using OAuth2 without understanding the whole specification is challenging and prone to errors, even when
SDKs are being used. The primary goal of ORY Hydra is to make OAuth 2.0 and OpenID Connect 1.0 better accessible.

ORY Hydra implements the flows described in OAuth2 and OpenID Connect 1.0 without forcing you to use a "Hydra User Management"
or some template engine or a predefined front-end. Instead it relies on HTTP redirection and cryptographic methods
to verify user consent allowing you to use ORY Hydra with any authentication endpoint, be it [authboss](https://github.com/go-authboss/authboss),
[auth0.com](https://auth0.com/) or your proprietary PHP authentication.

### ORY Hydra implements open standards

ORY Hydra implements Open Standards set by the IETF:

* [The OAuth 2.0 Authorization Framework](https://tools.ietf.org/html/rfc6749)
* [OAuth 2.0 Threat Model and Security Considerations](https://tools.ietf.org/html/rfc6819)
* [OAuth 2.0 Token Revocation](https://tools.ietf.org/html/rfc7009)
* [OAuth 2.0 Token Introspection](https://tools.ietf.org/html/rfc7662)
* [OAuth 2.0 Dynamic Client Registration Protocol](https://tools.ietf.org/html/rfc7591)
* [OAuth 2.0 Dynamic Client Registration Management Protocol](https://tools.ietf.org/html/rfc7592)
* [OAuth 2.0 for Native Apps](https://tools.ietf.org/html/draft-ietf-oauth-native-apps-10)

and the OpenID Foundation:

* [OpenID Connect Core 1.0](http://openid.net/specs/openid-connect-core-1_0.html)
* [OpenID Connect Discovery 1.0](https://openid.net/specs/openid-connect-discovery-1_0.html)

## ORY Cloud Security

ORY Cloud Security (OCS) is a popular, hardened toolbox of open and closed source services for securing your APIs
and web services. Our products are capable of handling high throughput with low latency, work flawlessly in the
cloud and on premise, and are securing millions of requests each day already.

### Products

* **ORY Hydra+** is a closed source modification of ORY Hydra with fault-tolerance, low latency, increased throughput,
reduced database load, stateless authorization, and other features. Pricing starts at 399€ per month. For more information [send us an inquiry](https://docs.google.com/forms/d/e/1FAIpQLSfRV34gUWHF1CnOa5oiWj_AMhh7SSGwwb_6DeMNhH-EQqnAwQ/viewform).
* **ORY Security Console** is an admin user interface for ORY Hydra and ORY Gatekeeper. Available on-premise and as a service.
Pricing starts at 49€ / month. For more information [send us an inquiry](https://docs.google.com/forms/d/e/1FAIpQLSfRV34gUWHF1CnOa5oiWj_AMhh7SSGwwb_6DeMNhH-EQqnAwQ/viewform).
* **ORY Gatekeeper** is a configurable firewall reverse proxy making access control easy and secure. ORY Gatekeeper is compatible with ORY Hydra
Access Control Policies and OAuth 2.0 Access Tokens. Pricing starts at 99€ / month. For more information [send us an inquiry](https://docs.google.com/forms/d/e/1FAIpQLSfRV34gUWHF1CnOa5oiWj_AMhh7SSGwwb_6DeMNhH-EQqnAwQ/viewform).

* **ORY Hydra** - this repository.
* **ORY Fosite** is a popular open source OAuth2 and OpenID Connect SDK for Go.
* **ORY Ladon** is an acclaimed open source SDK for access control policies, solving authorization in the microservice and IoT age.

### Consulting

Do you need help with setting up ORY Hydra in your environment, do you require custom features, or have other consulting
needs? Then [contact us now!](https://docs.google.com/forms/d/e/1FAIpQLScy_mUDYCF9BTul3TPqYh-Fw-TCFkHHyhlzLOubrGY899hrqA/viewform).

## Quickstart

This section is a quickstart guide to working with ORY Hydra. In-depth docs are available as well:

* The documentation is available on [GitBook](https://ory.gitbooks.io/hydra/content/).
* The REST API documentation is available at [Apiary](http://docs.hydra13.apiary.io/).

### 5 minutes tutorial: Run your very own OAuth2 environment

The **[tutorial](https://ory.gitbooks.io/hydra/content/tutorial.html)** teaches you to set up ORY Hydra,
a Postgres instance and an exemplary identity provider written in React using docker compose.
It will take you about 5 minutes to complete the **[tutorial](https://ory.gitbooks.io/hydra/content/tutorial.html)**.

<img src="docs/images/oauth2-flow.gif" alt="OAuth2 Flow">

<br clear="all">

### Installation

There are various ways of installing ORY Hydra on your system.

#### Download binaries

The client and server **binaries are downloadable at [releases](https://github.com/ory/hydra/releases)**.
There is currently no installer available. You have to add the ORY Hydra binary to the PATH environment variable yourself or put
the binary in a location that is already in your path (`/usr/bin`, ...). 
If you do not understand what that all of this means, ask in our [chat channel](https://gitter.im/ory-am/hydra). We are happy to help.

#### Using Docker

**Starting the host** is easiest with docker. The host process handles HTTP requests and is backed by a database.
Read how to install docker on [Linux](https://docs.docker.com/linux/), [OSX](https://docs.docker.com/mac/) or
[Windows](https://docs.docker.com/windows/). ORY Hydra is available on [Docker Hub](https://hub.docker.com/r/oryd/hydra/).

You can use ORY Hydra without a database, but be aware that restarting, scaling
or stopping the container will **lose all data**:

```
$ docker run -e "DATABASE_URL=memory" -e "ISSUER=https://localhost:4444/" -d --name my-hydra -p 4444:4444 oryd/hydra
ec91228cb105db315553499c81918258f52cee9636ea2a4821bdb8226872f54b
```

*Note: We had to create a new docker hub repository. Tags prior to 0.7.5 are available [here](https://hub.docker.com/r/ory-am/hydra/).*

**Using the client command line interface:** You can ssh into the ORY Hydra container
and execute the ORY Hydra command from there:

```
$ docker exec -i -t <hydra-container-id> /bin/bash
# e.g. docker exec -i -t ec91228 /bin/bash

root@ec91228cb105:/go/src/github.com/ory/hydra# hydra
Hydra is a twelve factor OAuth2 and OpenID Connect provider

[...]
```

#### Building from source

If you wish to compile ORY Hydra yourself, you need to install and set up [Go 1.8+](https://golang.org/) and add `$GOPATH/bin`
to your `$PATH`. To do so, run the following commands in a shell (bash, sh, cmd.exe, ...):

```
go get -d -u github.com/ory/hydra
go get github.com/Masterminds/glide
cd $GOPATH/src/github.com/ory/hydra
glide install
go install github.com/ory/hydra
hydra
```

**Notes**

* We changed organization name from `ory-am` to `ory`. In order to keep backwards compatibility, we did not rename Go packages.
* You can ignore warnings similar to `package github.com/ory/hydra/cmd/server: case-insensitive import collision: "github.com/Sirupsen/logrus" and "github.com/sirupsen/logrus"`.

## Security

*Why should I use ORY Hydra? It's not that hard to implement two OAuth2 endpoints and there are numerous SDKs out there!*

OAuth2 and OAuth2 related specifications are over 400 written pages. Implementing OAuth2 is easy, getting it right is hard.
ORY Hydra is trusted by companies all around the world, has a vibrant community and faces millions of requests in production
each day. Of course, we also compiled a security guide with more details on cryptography and security concepts.
Read [the security guide now](https://ory.gitbooks.io/hydra/content/security.html).

## Telemetry

ORY Hydra collects summarized, anonymized telemetry which can optionally be turned off. Click [here](https://ory.gitbooks.io/hydra/content/telemetry.html)
to learn more.

## Sponsors & Adopters

ORY Hydra is open source with a permissive license. We are a dedicated, young but also experienced team of developers
in the area of cloud computing and internet security. Please support our mission for a safer web and become a sponsor, buy an
engineer membership, or [tell us](mailto:hi@ory.am) why and how you are using ORY Hydra.

### Sponsorship

**[Become a ORY Hydra sponsor now](https://www.patreon.com/bePatron?c=581435&rid=1746199)** and help keep the project
open source, free, and maintained. Additionally, sponsors unlock

* Access to ORY's private slack community.
* Featured logo & link in the readme and on the website.
* placement in release announcements.

Sponsorship starts at 600$ per month. We accept smaller contributions as well, but without the benefits listed above.

### Sponsors

<img src="docs/images/sponsors/auth0.png" align="left" width="30%" alt="Auth0.com" />

We are proud to have [Auth0](https://auth0.com) as a **gold sponsor** for ORY Hydra. [Auth0](https://auth0.com) solves
the most complex identity use cases with an extensible and easy to integrate platform that secures billions of logins
every year. At ORY, we use [Auth0](https://auth0.com) in conjunction with ORY Hydra for various internal projects.

<br clear="all"/>

### Adopters

ORY Hydra is battle-tested in production systems. This is a curated list of ORY Hydra adopters. [Tell us](mailto:hi@ory.am) if
you think that your company should be listed here.

<img src="https://www.arduino.cc/en/uploads/Trademark/ArduinoCommunityLogo.png" align="left" width="30%" alt="arduino.cc"/>

<p>Arduino is an open-source electronics platform based on easy-to-use hardware and software. It's intended
for anyone making interactive projects. ORY Hydra secures Arduino's developer platform.</p>

<br clear="all"/>

## Documentation

### Guide

The Guide is available on [GitBook](https://ory.gitbooks.io/hydra/content/).

### HTTP API documentation

The HTTP API is documented at [Apiary](http://docs.hydra13.apiary.io/).

### Command line documentation

Run `hydra -h` or `hydra help`.

### Develop

Developing with ORY Hydra is as easy as:

```
go get -d -u github.com/ory/hydra
go get github.com/Masterminds/glide
cd $GOPATH/src/github.com/ory/hydra
glide install
go test $(glide novendor)
```

Then run it with in-memory database:

```
DATABASE_URL=memory go run main.go host
```

**Notes**

* We changed organization name from `ory-am` to `ory`. In order to keep backwards compatibility, we did not rename Go packages.
* You can ignore warnings similar to `package github.com/ory/hydra/cmd/server: case-insensitive import collision: "github.com/Sirupsen/logrus" and "github.com/sirupsen/logrus"`.

## Reception

Hydra has received a lot of positive feedback. Let's see what the community is saying:

> Nice! Lowering barriers to the use of technologies like these is important.

[Pyxl101](https://news.ycombinator.com/item?id=11798641)

> OAuth is a framework not a protocol. The security it provides can vary greatly between implementations.
Fosite (which is what this is based on) is a very good implementation from a security perspective: https://github.com/ory/fosite#a-word-on-security

[abritishguy](https://news.ycombinator.com/item?id=11800515)

> [...] Thanks for releasing this by the way, looks really well engineered. [...]

## Libraries and third-party projects

Official:
* [Consent App SDK For NodeJS](https://github.com/ory/hydra-js)
* [Consent App Example written in Go](https://github.com/ory/hydra-consent-app-go)
* [Exemplary Consent App with Express and NodeJS](https://github.com/ory/hydra-consent-app-express)

Community:
* [Consent App SDK for Go](https://github.com/janekolszak/idp)
* [ORY Hydra middleware for Gin](https://github.com/janekolszak/gin-hydra)
* [Kubernetes helm chart](https://github.com/kubernetes/charts/pull/1022)

## Blog posts & articles

* [Creating an oauth2 custom lamda authorizer for use with Amazons (AWS) API Gateway using Hydra](https://blogs.edwardwilde.com/2017/01/12/creating-an-oauth2-custom-lamda-authorizer-for-use-with-amazons-aws-api-gateway-using-hydra/)
* Warning, ORY Hydra has changed almost everything since writing this
article: [Hydra: Run your own Identity and Access Management service in <5 Minutes](https://blog.gopheracademy.com/advent-2015/hydra-auth/)
