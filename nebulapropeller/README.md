Nebula Propeller
===============

[![Current Release](https://img.shields.io/github/release/nebulaclouds/nebulapropeller.svg)](https://github.com/nebulaclouds/nebulapropeller/releases/latest)
![Master](https://github.com/nebulaclouds/nebulapropeller/workflows/Master/badge.svg)
[![GoDoc](https://godoc.org/github.com/nebulaclouds/nebulapropeller?status.svg)](https://pkg.go.dev/mod/github.com/nebulaclouds/nebulapropeller)
[![License](https://img.shields.io/badge/LICENSE-Apache2.0-ff69b4.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)
[![CodeCoverage](https://img.shields.io/codecov/c/github/nebulaclouds/nebulapropeller.svg)](https://codecov.io/gh/nebulaclouds/nebulapropeller)
[![Go Report Card](https://goreportcard.com/badge/github.com/nebulaclouds/nebulapropeller)](https://goreportcard.com/report/github.com/nebulaclouds/nebulapropeller)
![Commit activity](https://img.shields.io/github/commit-activity/w/nebulaclouds/nebulapropeller.svg?style=plastic)
![Commit since last release](https://img.shields.io/github/commits-since/nebulaclouds/nebulapropeller/latest.svg?style=plastic)
[![Slack](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://slack.nebula.org)

Kubernetes operator to executes Nebula graphs natively on kubernetes

Components
==========

Propeller
---------
Propeller is a K8s native operator that executes Nebula workflows. Workflow Spec is written in Protobuf for
cross-compatibility.

Propeller Webhook
-----------------
A Mutating Webhook that can be optionally deployed to extend Nebula Propeller's functionality. It currently supports
enables injecting secrets into pods launched directly or indirectly through Nebula backend plugins.

kubectl-nebula
-------------
A Kubectl-plugin to interact with Nebula Workflow CRDs. It enables retrieving and rendering Nebula Workflows in CLI as
well as safely aborting running workflows.

Getting Started
===============
kubectl-nebula tool
------------------
kubectl-nebula is an command line tool that can be used as an extension to kubectl. It is a separate binary that is built
from the propeller repo.

Install
-------
This command will install kubectl-nebula and nebulapropeller to `~/go/bin`

```
   $ make compile
```

You can also use [Krew](https://github.com/kubernetes-sigs/krew) to install the kubectl-nebula CLI:

```
   $ kubectl krew install nebula
```

Use
---
Two ways to execute the command, either standalone *kubectl-nebula* or as a subcommand of *kubectl*

```
    $ kubectl-nebula --help
    OR
    $ kubectl nebula --help
    Nebula is a serverless workflow processing platform built for native execution on K8s.
          It is extensible and flexible to allow adding new operators and comes with many operators built in

    Usage:
      kubectl-nebula [flags]
      kubectl-nebula [command]

    Available Commands:
      compile     Compile a workflow from core proto-buffer files and output a closure.
      config      Runs various config commands, look at the help of this command to get a list of available commands..
      create      Creates a new workflow from proto-buffer files.
      delete      delete a workflow
      get         Gets a single workflow or lists all workflows currently in execution
      help        Help about any command
      visualize   Get GraphViz dot-formatted output.
```

Observing running workflows
---------------------------

To retrieve all workflows in a namespace use the --namespace option, --namespace = "" implies all namespaces.

```
   $ kubectl-nebula get --namespace nebulakit-development
    workflows
    ├── nebulakit-development/nebulakit-development-f01c74085110840b8827 [ExecId: ... ] (2m34s Succeeded) - Time SinceCreation(30h1m39.683602s)
    ...
    Found 19 workflows
    Success: 19, Failed: 0, Running: 0, Waiting: 0
```

To retrieve a specific workflow, namespace can either be provided in the format namespace/name or using the --namespace
argument

```
   $ kubectl-nebula get nebulakit-development/nebulakit-development-ff806e973581f4508bf1
    Workflow
    └── nebulakit-development/nebulakit-development-ff806e973581f4508bf1 [ExecId: project:"nebulakit" domain:"development" name:"ff806e973581f4508bf1" ] (2m32s Succeeded )
        ├── start-node start 0s Succeeded
        ├── c task 0s Succeeded
        ├── b task 0s Succeeded
        ├── a task 0s Succeeded
        └── end-node end 0s Succeeded
```

Deleting workflows
------------------
To delete a specific workflow

```
   $ kubectl-nebula delete --namespace nebulakit-development nebulakit-development-ff806e973581f4508bf1
```

To delete all completed workflows - they have to be either success/failed with a special isCompleted label set on them.
The Label is set [here](https://github.com/nebulaclouds/nebulapropeller/blob/master/pkg/controller/controller.go#L247)

```
   $ kubectl-nebula delete --namespace nebulakit-development --all-completed
```

Running propeller locally
-------------------------
use the config.yaml in root found [here](https://github.com/nebulaclouds/nebulapropeller/blob/master/config.yaml). Cd into
this folder and then run

```
   $ nebulapropeller
```

Following dependencies need to be met

1. Blob store (you can forward minio port to localhost)
2. Admin Service endpoint (can be forwarded) OR *Disable* events to admin and launchplans
3. access to kubeconfig and kubeapi

Running webhook
---------------

API Server requires the webhook to serve traffic over SSL. To issue self-signed certs to be used for serving traffic,
use:

```
    $ nebulapropeller webhook init-certs
```

This will create a ca.crt, tls.crt and key.crt and store them to nebula-pod-webhook secret. If a secret of the same name
already exist, it'll not override it.

Starting the webhook can be done by running:

```
    $ nebulapropeller webhook
```

The secret should be mounted and accessible to this command. It'll then create a MutatingWebhookConfiguration object
with the details of the webhook and that registers the webhook with ApiServer.

Making changes to CRD
=====================
*Remember* changes to CRD should be carefully done, they should be backwards compatible or else you should use proper
operator versioning system. Once you do the changes, you have to follow the following steps.

- Ensure the propeller code is checked out in `$GOPATH/github.com/nebulaclouds/nebulapropeller`
- Uncomment https://github.com/nebulaclouds/nebulapropeller/blob/master/hack/tools.go#L5

```bash
    $ go mod vendor
```

- Now generate the code

```bash
    $ make op_code_generate
```

**Why do we have to do this?**
Nebulapropeller uses old way of writing Custom controllers for K8s. The k8s.io/code-generator only works in the GOPATH
relative code path (sadly). So you have checkout the code in the right place. Also, `go mod vendor` is needed to get
code-generator in a discoverable path.

**TODO**

1. We may be able to avoid needing the old style go-path
2. Migrate to using controller runtime

