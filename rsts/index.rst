.. toctree::
   :maxdepth: 1
   :name: mainsections
   :titlesonly:
   :hidden:

   |plane| Getting Started <https://docs.nebula.org/projects/cookbook/en/latest/index.html>
   |book-reader| User Guide <https://docs.nebula.org/projects/cookbook/en/latest/user_guide.html>
   |chalkboard| Tutorials <https://docs.nebula.org/projects/cookbook/en/latest/tutorials.html>
   |project-diagram| Concepts <concepts/basics>
   |rocket| Deployment and Administration <deployment/index>
   |book| API Reference <reference/index>
   |hands-helping| Community <community/index>

.. toctree::
   :caption: Concepts
   :maxdepth: -1
   :name: divedeeptoc
   :hidden:

   concepts/basics
   concepts/control_plane
   concepts/architecture

.. toctree::
   :caption: Deployment
   :maxdepth: -1
   :name: deploymenttoc
   :hidden:

   deployment/index
   deployment/deployment/index
   deployment/plugins/index
   deployment/agents/index
   deployment/configuration/index
   deployment/configuration/generated/index
   deployment/security/index


.. toctree::
   :caption: Community
   :maxdepth: -1
   :name: roadmaptoc
   :hidden:

   Community <community/index>
   community/contribute
   community/roadmap
   Frequently Asked Questions <https://github.com/nebulaclouds/nebula/discussions/categories/q-a>
   community/troubleshoot

.. toctree::
   :caption: API Reference
   :maxdepth: -1
   :name: apireference
   :hidden:

   API Reference <reference/index>


*************************************************
Production-grade Data and ML Workflows Made Easy
*************************************************

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

.. raw:: html

   <p style="color: #808080; font-weight: 350; font-size: 25px; padding-top: 10px; padding-bottom: 10px;">
   Highly scalable and flexible workflow orchestration for prototyping and production
   </p>

`Nebula Tags <_tags/tagsindex.html>`__

`Nebula <https://github.com/nebulaclouds/nebula>`__ is an open-source, Kubernetes-native
workflow orchestrator implemented in `Go <https://go.dev/>`__. It enables highly
concurrent, scalable and reproducible workflows for data processing, machine
learning and analytics.

Created at `Lyft <https://www.lyft.com/>`__ in collaboration with Spotify,
Freenome, and many others, Nebula provides first-class support for
`Python <https://docs.nebula.org/projects/nebulakit/en/latest/>`__,
`Java, and Scala <https://github.com/nebulaclouds/nebulakit-java>`__. Data Scientists
and ML Engineers in the industry use Nebula to create:

- ETL pipelines for petabyte-scale data processing.
- Analytics workflows for business and finance use cases.
- Machine learning pipelines for logistics, image processing, and cancer diagnostics.

Explore Nebula
=============

Get a birds-eye view 🦅 of Nebula at the `official website <https://nebula.org/>`__:

.. panels::
    :header: text-center
    :column: col-lg-12 p-2

    .. link-button:: https://nebula.org/features
       :type: url
       :text: ⭐️ Core features
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    From strongly typed interfaces to container-native DAGs, Nebula mitigates the
    trade-off between scalability and usability.

    ---

    .. link-button:: https://nebula.org/integrations
       :type: url
       :text: 🤝 Integrations
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    From strongly typed interfaces to container-native DAGs, Nebula mitigates the
    trade-off between scalability and usability.

    ---

    .. link-button:: https://nebula.org/airflow-alternative
       :type: url
       :text: 💨 Nebula vs Airflow
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Say goodbye to Airflow versioning pain and stepping over your teammate's toes
    when you change your package versions. Ouch!

    ---

    .. link-button:: https://nebula.org/kubeflow-alternative
       :type: url
       :text: 🔁 Nebula vs Kubeflow
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Unintuitive Python DSL boilerplate got you down? With ``nebulakit`` you just
    write Python code and Nebula compiles down to type-safe execution graphs.

    ---

    .. link-button:: https://nebula.org/blog
       :type: url
       :text: 📝 Blog
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Learn more about orchestration, Nebula, and everything in between.

    ---

    .. link-button:: https://nebula.org/events
       :type: url
       :text: 🗓 Events
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Keep up-to-date with Nebula's upcoming talks, conferences, and more.


Learn Nebula
===========

The following main sections in the documentation will guide you through your
Nebula journey, whether you want to write Nebula workflows, deploy the Nebula
platform to your K8s cluster, or extend and contribute its architecture and
design.

.. panels::
    :header: text-center
    :column: col-lg-12 p-2

    .. link-button:: cookbook:getting_started_index
       :type: ref
       :text: 🔤 Getting Started
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Get your first workflow running, learn about the Nebula development lifecycle,
    and see the core use cases that Nebula enables.

    ---

    .. link-button:: cookbook:userguide
       :type: ref
       :text: 📖 User Guide
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    A comprehensive view of Nebula's functionality for data scientists, ML engineers,
    data engineers, and data analysts.

    ---

    .. link-button:: cookbook:tutorials
       :type: ref
       :text: 📚 Tutorials
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    End-to-end examples of Nebula for data/feature engineering, machine learning,
    bioinformatics, and more.

    ---

    .. link-button:: cookbook:integrations
       :type: ref
       :text: 🤝 Integrations
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Learn how to leverage a rich ecosystem of third-party tools and libraries
    to make your Nebula workflows even more effective.

    ---

    .. link-button:: deployment
       :type: ref
       :text: 🚀 Deployment Guide
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Guides for platform engineers to deploy and maintain a Nebula cluster on your
    own infrastructure.

    ---

    .. link-button:: reference
       :type: ref
       :text: 📒 API Reference
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Reference for all of Nebula's component libraries.

    ---

    .. link-button:: divedeep
       :type: ref
       :text: 🧠 Concepts
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Dive deep into all of Nebula's concepts, from tasks and workflows to the underlying Nebula scheduler.

    ---

    .. link-button:: community
       :type: ref
       :text: 🤗 Community
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Join the fast-growing Nebula community to get help, ask questions, and contribute!

Get Help
========

Have questions or need support? The best way to reach us is through Slack:

.. panels::
    :header: text-center
    :column: col-lg-12 p-2

    .. link-button:: https://nebula-org.slack.com/archives/CP2HDHKE1
       :type: url
       :text: 🤔 Ask the Community
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Ask anything related to Nebula and get a response within a few hours.

    ---

    .. link-button:: https://nebula-org.slack.com/archives/C01RXBFV1M5
       :type: url
       :text: 👋 Introduce yourself
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Tell us about yourself. We'd love to know about you and what brings you to Nebula.

    ---

    .. link-button:: https://nebula-org.slack.com/archives/CPQ3ZFQ84
       :type: url
       :text: 💭 Share ideas
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    Share any suggestions or feedback you have on how to make Nebula better.

    ---

    .. link-button:: https://nebula-org.slack.com/archives/C01P3B761A6
       :type: url
       :text: 🛠 Get help with deploment
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    If you need any help with Nebula deployment, hit us up.
