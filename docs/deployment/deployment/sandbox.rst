.. _deployment-deployment-sandbox:

#########################
Sandbox Deployment
#########################

.. tags:: Kubernetes, Infrastructure, Basic

A sandbox deployment of Nebula bundles together portable versions of Nebula's
dependencies such as a relational database and durable object store.

For the blob store requirements, Nebula Sandbox uses `Minio <https://min.io/>`__,
which offers an S3 compatible interface, and for Postgres, it uses the stock
Postgres Docker image and Helm chart.

.. important::

    The sandbox deployment is not suitable for production environments. For instructions on how to create a
    production-ready Nebula deployment, checkout the :ref:`Deployment Paths <deployment-deployment>` guide.

*******************************************
Nebula Sandbox as a Single Docker Container
*******************************************

Nebula provides a way for creating a Nebula cluster as a self-contained Docker image. This is mini-replica of an
entire Nebula deployment, without the scalability and with minimal extensions.

The Nebula Sandbox can be run on any environment that supports containers and makes it extremely easy for users of Nebula
to try out the platform and get a feel for the user experience, all without having to understand Kubernetes or dabble
with configuration.

.. note::

   The Nebula single container sandbox is also used by the team to run continuous integration tests and used by the
   :ref:`cookbook:userguide`, :ref:`cookbook:tutorials` and :ref:`cookbook:integrations` documentation.

Requirements
============

- Install `kubectl <https://kubernetes.io/docs/tasks/tools/install-kubectl/>`__.
- Install `docker <https://docs.docker.com/engine/install/>`__ or any other OCI-compatible tool, like Podman or LXD.
- Install `nebulactl <https://github.com/nebulaclouds/nebulactl>`__, the official CLI for Nebula.

While Nebula can run any OCI-compatible task image using the default Kubernetes container runtime (``containerd``), the Nebula
core maintainers typically use Docker. Note that the ``nebulactl demo`` command does rely on Docker APIs, but as this
demo environment is just one self-contained image, you can also run the image directly using another run time.

Within the single container environment, a mini Kubernetes cluster is installed using `k3s <https://k3s.io/>`__. K3s
uses an in-container Docker daemon, run using `docker-in-docker configuration <https://www.docker.com/blog/docker-can-now-run-within-docker/>`__
to orchestrate user containers.

Start the Sandbox
==================

To spin up a Nebula Sandbox, run:

.. prompt:: bash $

    nebulactl demo start

This command runs a Docker container, which itself comes with a Docker registry
on ``localhost:30000`` so you can build images outside of the docker-in-docker
container by tagging your containers with ``localhost:30000/imgname:tag`` and
pushing the image.

The local Postgres installation is also available on port ``30001`` for users
who wish to dig deeper into the storage layer.

.. div:: shadow p-3 mb-8 rounded

   **Expected Output:**

   .. code-block::

      👨‍💻 Nebula is ready! Nebula UI is available at http://localhost:30080/console 🚀 🚀 🎉
      ❇️ Run the following command to export sandbox environment variables for accessing nebulactl
      	export NEBULACTL_CONFIG=~/.nebula/config-sandbox.yaml
      🐋 Nebula sandbox ships with a Docker registry. Tag and push custom workflow images to localhost:30000
      📂 The Minio API is hosted on localhost:30002. Use http://localhost:30080/minio/login for Minio console


Configuration
______________

The ``config-sandbox.yaml`` file contains configuration for **NebulaAdmin**,
which is the Nebula cluster backend component that processes all client requests
such as workflow executions. The default values are enough to let you connect and use Nebula:


.. code-block:: yaml
   
   admin:
     # For GRPC endpoints you might want to use dns:///nebula.myexample.com
     endpoint: localhost:30080
     authType: Pkce
     insecure: true
     console:
       endpoint: http://localhost:30080
   logger:
     show-source: true
   level: 0

.. note:: 
   
   You can also create your own config file with `nebulactl config init`, which
   will create a config file at `~/.nebula/config.yaml`.

   Learn more about the configuration settings in the
   {ref}`Deployment Guide <nebula:nebulaadmin-config-specification>`



Now that you have the sandbox cluster running, you can now go to the :ref:`User Guide <cookbook:userguide>` or
:ref:`Tutorials <cookbook:tutorials>` to run tasks and workflows written in ``nebulakit``, the Python SDK for Nebula.
