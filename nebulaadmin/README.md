NebulaAdmin
==========

[![Current Release](https://img.shields.io/github/release/nebulaclouds/nebulaadmin.svg)](https://github.com/nebulaclouds/nebulaadmin/releases/latest)
![Master](https://github.com/nebulaclouds/nebulaadmin/workflows/Master/badge.svg)
[![GoDoc](https://godoc.org/github.com/nebulaclouds/nebulaadmin?status.svg)](https://pkg.go.dev/mod/github.com/nebulaclouds/nebulaadmin)
[![License](https://img.shields.io/badge/LICENSE-Apache2.0-ff69b4.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)
[![CodeCoverage](https://img.shields.io/codecov/c/github/nebulaclouds/nebulaadmin.svg)](https://codecov.io/gh/nebulaclouds/nebulaadmin)
[![Go Report Card](https://goreportcard.com/badge/github.com/nebulaclouds/nebulaadmin)](https://goreportcard.com/report/github.com/nebulaclouds/nebulaadmin)
![Commit activity](https://img.shields.io/github/commit-activity/w/nebulaclouds/nebulaadmin.svg?style=plastic)
![Commit since last release](https://img.shields.io/github/commits-since/nebulaclouds/nebulaadmin/latest.svg?style=plastic)
[![Slack](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://slack.nebula.org)

NebulaAdmin is the control plane for Nebula responsible for managing entities (task, workflows, launch plans) and
administering workflow executions. NebulaAdmin implements the
[AdminService](https://github.com/nebulaclouds/nebulaidl/blob/master/protos/nebulaidl/service/admin.proto) which
defines a stateless REST/gRPC service for interacting with registered Nebula entities and executions.
NebulaAdmin uses a relational style Metadata Store abstracted by [GORM](http://gorm.io/) ORM library.

For more background on Nebula, check out the official [website](https://nebula.org/) and the [docs](https://docs.nebula.org/en/latest/index.html)

Before Check-In
---------------

Nebula Admin has a few useful make targets for linting and testing. Please use these before checking in to help suss out
minor bugs and linting errors.

```
  # Please make sure you have all the dependencies installed:
  $ make install
  
  # In case you are only missing goimports:
  $ go install golang.org/x/tools/cmd/goimports@latest
  $ make goimports
```

```
  $ make test_unit
```

```
  $ make lint
```
