# ![Ory/Hydra](docs/dist/images/logo.png)

[![Join the chat at https://gitter.im/ory-am/hydra](https://img.shields.io/badge/join-chat-00cc99.svg)](https://gitter.im/ory-am/hydra?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Join mailinglist](https://img.shields.io/badge/join-mailinglist-00cc99.svg)](https://groups.google.com/forum/#!forum/ory-hydra/new)
[![Join newsletter](https://img.shields.io/badge/join-newsletter-00cc99.svg)](http://eepurl.com/bKT3N9)
[![Follow newsletter](https://img.shields.io/badge/follow-twitter-00cc99.svg)](https://twitter.com/_aeneasr)
[![Follow GitHub](https://img.shields.io/badge/follow-github-00cc99.svg)](https://github.com/arekkas)

[![Build Status](https://travis-ci.org/ory-am/hydra.svg?branch=master)](https://travis-ci.org/ory-am/hydra)
[![Coverage Status](https://coveralls.io/repos/ory-am/hydra/badge.svg?branch=master&service=github)](https://coveralls.io/github/ory-am/hydra?branch=master)
[![Code Climate](https://codeclimate.com/github/ory-am/hydra/badges/gpa.svg)](https://codeclimate.com/github/ory-am/hydra)
[![Go Report Card](https://goreportcard.com/badge/github.com/ory-am/hydra)](https://goreportcard.com/report/github.com/ory-am/hydra)


[![Docs Guide](https://img.shields.io/badge/docs-guide-blue.svg)](https://ory-am.gitbooks.io/hydra/content/)
[![HTTP API Documentation](https://img.shields.io/badge/docs-http%20api-blue.svg)](http://docs.hdyra.apiary.io/)
[![Code Documentation](https://img.shields.io/badge/docs-godoc-blue.svg)](https://godoc.org/github.com/ory-am/hydra)

Hydra is being developed by german-based company [Ory](https://ory.am).
Join our [newsletter](http://eepurl.com/bKT3N9) to stay on top of new developments.
We offer basic support requests on [Google Groups](https://groups.google.com/forum/#!forum/ory-hydra/new) and [Gitter](https://gitter.im/ory-am/hydra)
as well as [consulting](mailto:hi@ory.am) around integrating Hydra into
your particular environment and [premium support](mailto:hi@ory.am).

Hydra uses the security first OAuth2 and OpenID Connect SDK [Fosite](https://github.com/ory-am/fosite) and [Ladon](https://github.com/ory-am/ladon) for policy-based access control.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [What is Hydra?](#what-is-hydra)
- [Feature Overview](#feature-overview)
- [Quickstart](#quickstart)
  - [Installation](#installation)
    - [Download binaries](#download-binaries)
    - [Using Docker](#using-docker)
    - [Building from source](#building-from-source)
  - [5 minutes tutorial: Run your very own OAuth2 environment](#5-minutes-tutorial-run-your-very-own-oauth2-environment)
- [Security](#security)
- [Reception](#reception)
- [Documentation](#documentation)
  - [Guide](#guide)
  - [HTTP API Documentation](#http-api-documentation)
  - [Command Line Documentation](#command-line-documentation)
  - [Develop](#develop)
- [Hall of Fame](#hall-of-fame)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## What is Hydra?

At first, there was the monolith. The monolith worked well with the bespoke authentication module.
Then, the web evolved into an elastic cloud that serves thousands of different user agents
in every part of the world.

Hydra is driven by the need for a **scalable, low-latency, in memory
Access Control, OAuth2, and OpenID Connect layer** that integrates with every identity provider you can imagine.

Hydra is available through [Docker](https://hub.docker.com/r/oryam/hydra/) and relies on RethinkDB for persistence.
Database drivers are extensible in case you want to use RabbitMQ, MySQL, MongoDB, or some other database instead.

Hydra is built for high throughput environments. Using 10.000 simultaneous connections on a Macbook Pro Late 2013,
the OAuth2 token validation endpoint served on average **37500 requests / sec**. Other endpoints like the JSON Web Key endpoint
serve up to 4700 requests / sec. Read [this issue](https://github.com/ory-am/hydra/issues/161) for information on reproducing these benchmarks yourself.

## Feature Overview

1. **Availability:** Hydra uses pub/sub to have the latest data available in memory. The in-memory architecture allows for heavy duty workloads.
2. **Scalability:** Hydra scales effortlessly on every platform you can imagine, including Heroku, Cloud Foundry, Docker,
Google Container Engine and many more.
3. **Integration:** Hydra wraps your existing stack like a blanket and keeps it safe. Hydra uses cryptographic tokens to authenticate users and request their consent, no APIs required.
The deprecated php-3.0 authentication service your intern wrote? It works with that too, don't worry.
We wrote an example with React to show you what this could look like: [React.js Identity Provider Example App](https://github.com/ory-am/hydra-idp-react).
4. **Security:** Hydra leverages the security first OAuth2 framework **[Fosite](https://github.com/ory-am/fosite)**,
encrypts important data at rest, and supports HTTP over TLS (https) out of the box.
5. **Ease of use:** Developers and operators are human. Therefore, Hydra is easy to install and manage. Hydra does not care if you use React, Angular, or Cocoa for your user interface.
To support you even further, there are APIs available for *cryptographic key management, social log on, policy based access control, policy management, and two factor authentication (tbd).*
Hydra is packaged using [Docker](https://hub.docker.com/r/oryam/hydra/).
6. **Open Source:** Hydra is licensed under Apache Version 2.0
7. **Professional:** Hydra implements peer reviewed open standards published by [The Internet Engineering Task Force (IETF®)](https://www.ietf.org/) and the [OpenID Foundation](https://openid.net/)
and under supervision of the [LMU Teaching and Research Unit Programming and Modelling Languages](http://www.en.pms.ifi.lmu.de). No funny business.
8.  <img src="docs/dist/images/monitoring.gif" width="45%" align="right"> **Real Time:** Operation is a lot easier with real time. There are no caches,
    no invalidation strategies and no magic - just simple, cloud native pub-sub. Hydra leverages RethinkDB, so check out their real time database monitoring too!

<br clear="all">

## Quickstart

This section is a quickstart guide to working with Hydra. In-depth docs are available as well:

* The documentation is available on [GitBook](https://ory-am.gitbooks.io/hydra/content/).
* The REST API documentation is available at [Apiary](http://docs.hdyra.apiary.io).

### Installation

There are various ways of installing hydra on your system.

#### Download binaries

The client and server **binaries are downloadable at [releases](https://github.com/ory-am/hydra/releases)**.
There is currently no installer available. You have to add the hydra binary to the PATH environment variable yourself or put
the binary in a location that is already in your path (`/usr/bin`, ...). 
If you do not understand what that all of this means, ask in our [chat channel](https://gitter.im/ory-am/hydra). We are happy to help.

#### Using Docker

**Starting the host** is easiest with docker. The host process handles HTTP requests and is backed by a database.
Read how to install docker on [Linux](https://docs.docker.com/linux/), [OSX](https://docs.docker.com/mac/) or
[Windows](https://docs.docker.com/windows/). Hydra is available on [Docker Hub](https://hub.docker.com/r/oryam/hydra/).

You can use Hydra without a database, but be aware that restarting, scaling
or stopping the container will **lose all data**:

```
$ docker run -d -p 4444:4444 oryam/hydra --name my-hydra
ec91228cb105db315553499c81918258f52cee9636ea2a4821bdb8226872f54b
```

**Using the client command line interface:** You can ssh into the hydra container
and execute the hydra command from there:

```
$ docker exec -i -t <hydra-container-id> /bin/bash
# e.g. docker exec -i -t ec91228 /bin/bash

root@ec91228cb105:/go/src/github.com/ory-am/hydra# hydra
Hydra is a twelve factor OAuth2 and OpenID Connect provider

[...]
```

#### Building from source

If you wish to compile hydra yourself, you need to install and set up [Go 1.5+](https://golang.org/) and add `$GOPATH/bin`
to your `$PATH`. To do so, run the following commands in a shell (bash, sh, cmd.exe, ...):

```
go get github.com/ory-am/hydra
go get github.com/Masterminds/glide
cd $GOPATH/src/github.com/ory-am/hydra
glide install
go install github.com/ory-am/hydra
hydra
```

### 5 minutes tutorial: Run your very own OAuth2 environment

The **[tutorial](https://ory-am.gitbooks.io/hydra/content/demo.md)** teaches you to set up Hydra,
a RethinkDB instance and an exemplary identity provider written in React using docker compose.
It will take you about 5 minutes to get complete the **[tutorial](https://ory-am.gitbooks.io/hydra/content/demo.md)**.

<img src="docs/dist/oauth2-flow.gif" alt="OAuth2 Flow">

<img alt="Running the example" align="right" width="35%" src="docs/dist/run-the-example.gif">

<br clear="all">

## Security

*Why should I use Hydra? It's not that hard to implement two OAuth2 endpoints and there are numerous SDKs out there!*

OAuth2 and OAuth2 related specifications are over 200 written pages. Implementing OAuth2 is easy, getting it right is hard.
Even if you use a secure SDK (there are numerous SDKs not secure by design in the wild), messing up the implementation
is a real threat - no matter how good you or your team is. To err is human.

An in-depth list of security features is listed [in the security guide]().

## Reception

Hydra has received a lot of positive feedback. Let's see what the community is saying:

> Nice! Lowering barriers to the use of technologies like these is important.

[Pyxl101](https://news.ycombinator.com/item?id=11798641)

> OAuth is a framework not a protocol. The security it provides can vary greatly between implementations.
Fosite (which is what this is based on) is a very good implementation from a security perspective: https://github.com/ory-am/fosite#a-word-on-security

[abritishguy](https://news.ycombinator.com/item?id=11800515)

> [...] Thanks for releasing this by the way, looks really well engineered. [...]

[olalonde](https://news.ycombinator.com/item?id=11798831)

## Documentation

### Guide

The Guide is available on [GitBook](https://ory-am.gitbooks.io/hydra/content/).

### HTTP API Documentation

The HTTP API is documented at [Apiary](http://docs.hdyra.apiary.io).

### Command Line Documentation

Run `hydra -h` or `hydra help`.

### Develop

Developing with Hydra is as easy as:

```
go get github.com/ory-am/hydra
go get github.com/Masterminds/glide
cd $GOPATH/src/github.com/ory-am/hydra
glide install
go test $(glide novendor)
```

If you want to run a Hydra instance, there are two possibilities:

Run without Database:
```
go run main.go host
```
 
Run against RethinkDB using Docker:
```
docker run --name some-rethink -d -p 8080:8080 -p 28015:28015 rethinkdb
DATABASE_URL=rethinkdb://localhost:28015/hydra go run main.go host
```

## Hall of Fame

A list of extraordinary contributors and [bug hunters](https://github.com/ory-am/hydra/issues/84).

* [Alexander Widerberg (leetal)](https://github.com/leetal) for implementing the prototype RethinkDB adapters.
* The active Community on Gitter.