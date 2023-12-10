# Welcome to Nebula!

```{eval-rst}
.. raw:: html

   <p style="color: #808080; font-weight: 350; font-size: 25px; padding-top: 10px; padding-bottom: 10px;">
   The highly scalable and flexible workflow orchestrator that unifies data, ML and analytics.
   </p>

.. image:: https://img.shields.io/badge/Graduate%20Project-Linux%20Foundation-purple?style=for-the-badge
    :target: https://lfaidata.foundation/projects/nebula/
    :alt: Linux Foundation

.. image:: https://img.shields.io/github/stars/nebulaclouds/nebula?label=github&logo=github&style=for-the-badge
   :target: https://github.com/nebulaclouds/nebula
   :alt: GitHub Repo stars

.. image:: https://img.shields.io/github/release/nebulaclouds/nebula.svg?style=for-the-badge&color=blue
   :target: https://github.com/nebulaclouds/nebula/releases/latest
   :alt: Nebula Release

.. image:: https://img.shields.io/github/actions/workflow/status/nebulaclouds/nebula/tests.yml?label=tests&style=for-the-badge
   :target: https://github.com/nebulaclouds/nebula/actions/workflows/tests.yml
   :alt: GitHub Test Status

.. image:: https://img.shields.io/github/actions/workflow/status/nebulaclouds/nebula/sandbox.yml?label=Sandbox%20docker%20image&style=for-the-badge
   :target: https://github.com/nebulaclouds/nebula/actions/workflows/sandbox.yml
   :alt: GitHub Sandbox Status

.. image:: https://img.shields.io/github/milestones/closed/nebulaclouds/nebula?style=for-the-badge
    :target: https://github.com/nebulaclouds/nebula/milestones?state=closed
    :alt: Completed Milestones

.. image:: https://img.shields.io/pypi/dm/nebulakit?color=blue&label=nebulakit%20downloads&style=for-the-badge&logo=pypi&logoColor=white
   :target: https://github.com/nebulaclouds/nebulakit
   :alt: Nebulakit Downloads

.. image:: https://img.shields.io/badge/Slack-Chat-pink?style=for-the-badge&logo=slack
    :target: https://slack.nebula.org
    :alt: Nebula Slack

.. image:: https://img.shields.io/badge/LICENSE-Apache2.0-ff69b4.svg?style=for-the-badge
    :target: http://www.apache.org/licenses/LICENSE-2.0.html
    :alt: License

.. |br| raw:: html

   <br>
   <br>

```

[Nebula](https://github.com/nebulaclouds/nebula) is an open-source, Kubernetes-native
workflow orchestrator implemented in [Go](https://go.dev/). It enables highly
concurrent, scalable and reproducible workflows for data processing, machine
learning and analytics.

Created at [Lyft](https://www.lyft.com/) in collaboration with Spotify,
Freenome, and many others, Nebula provides first-class support for
{doc}`Python <reference_nebulakit>`,
[Java, and Scala](https://github.com/nebulaclouds/nebulakit-java). Data Scientists
and ML Engineers in the industry use Nebula to create:

- Data pipelines for processing petabyte-scale data.
- Analytics workflows for business and finance use cases.
- Machine learning pipelines for logistics, image processing, and cancer diagnostics.

## Learn Nebula

The following guides will take you through Nebula, whether you want to write
workflows, deploy the Nebula platform to your K8s cluster, or extend and
contribute its architecture and design. You can also access the
{ref}`docs pages by tag <tagoverview>`.

```{list-table}
:header-rows: 0
:widths: 20 30

* - {doc}`🔤 Intro to Nebula <introduction>`
  - Get your first workflow running, learn about the Nebula development lifecycle
    and core use cases.
* - {doc}`📖 User Guide <userguide>`
  - A comprehensive view of Nebula's functionality for data and ML practitioners.
* - {doc}`📚 Tutorials <tutorials>`
  - End-to-end examples of Nebula for data/feature engineering, machine learning,
    bioinformatics, and more.
* - {doc}`🔌 Integrations <integrations>`
  - Leverage a rich ecosystem of third-party tools and libraries to make your
    Nebula workflows even more effective.
* - {ref}`🚀 Deployment Guide <deployment>`
  - Guides for platform engineers to deploy and maintain a Nebula cluster on your
    own infrastructure.
* - {ref}`🧠 Concepts <divedeep>`
  - Dive deep into all of Nebula's concepts, from tasks and workflows to the underlying Nebula scheduler.
```

## API Reference

Below are the API reference to the different components of Nebula:

```{list-table}
:header-rows: 0
:widths: 20 30

* - {doc}`Nebulakit <reference_nebulakit>`
  - Nebula's official Python SDK.
* - {doc}`NebulaCTL <reference_nebulactl>`
  - Nebula's command-line interface for interacting with a Nebula cluster.
* - {doc}`NebulaIDL <reference_nebulaidl>`
  - Nebula's core specification language.
```

## Get Help

Have questions or need support? The best way to reach us is through Slack:

```{list-table}
:header-rows: 0
:widths: 20 30

* - {ref}`🗓️ Resources <community>`
  - Find resources for office hours, newsletter, and slack.
* - [🤔 Ask the Community](https://nebula-org.slack.com/archives/CP2HDHKE1)
  - Ask anything related to Nebula and get a response within a few hours.
* - [👋 Introduce yourself](https://nebula-org.slack.com/archives/C01RXBFV1M5)
  - Tell us about yourself. We'd love to know about you and what brings you to Nebula.
* - [💭 Share ideas](https://nebula-org.slack.com/archives/CPQ3ZFQ84>)
  - Share any suggestions or feedback you have on how to make Nebula better.
* - [🛠 Get help with deploment](https://nebula-org.slack.com/archives/C01P3B761A6>)
  - If you need any help with Nebula deployment, hit us up.
```

```{toctree}
:maxdepth: 1
:hidden:

Introduction <introduction>
Nebula Fundamentals <nebulasnacks/getting_started/nebula_fundamentals>
Core Use Cases <nebulasnacks/getting_started/core_use_cases>
```

```{toctree}
:maxdepth: 1
:caption: Examples
:name: examples-guides
:hidden:

User Guide <userguide>
Tutorials <tutorials>
Integrations <integrations>
```

```{toctree}
:caption: Cluster Deployment
:maxdepth: -1
:name: deploymenttoc
:hidden:

Getting Started <deployment/index>
deployment/deployment/index
deployment/plugins/index
deployment/agents/index
deployment/configuration/index
deployment/configuration/generated/index
deployment/security/index
reference/swagger
```

```{toctree}
:maxdepth: 1
:caption: API Reference
:name: apitoc
:hidden:

nebulakit <reference_nebulakit>
nebulactl <reference_nebulactl>
nebulaidl <reference_nebulaidl>
```

```{toctree}
:maxdepth: 1
:caption: Ecosystem
:name: ecosystem
:hidden:

nebulakit-java <https://github.com/spotify/nebulakit-java>
unionml <https://unionml.readthedocs.io/>
pterodactyl <https://github.com/NotMatthewGriffin/pterodactyl>
latch sdk <https://docs.latch.bio/>
```

```{toctree}
:caption: Community
:maxdepth: -1
:name: roadmaptoc
:hidden:

Community Resources <community/index>
community/contribute
Contributing Examples <nebulasnacks/contribute>
community/roadmap
Frequently Asked Questions <https://github.com/nebulaclouds/nebula/discussions/categories/q-a>
community/troubleshoot
```

```{toctree}
:caption: Glossary
:maxdepth: -1
:name: divedeeptoc
:hidden:

Main Concepts <concepts/basics>
concepts/control_plane
concepts/architecture
```
