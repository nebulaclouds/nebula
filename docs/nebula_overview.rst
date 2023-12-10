.. nebulactl doc

####################################################
Nebulactl: The Official Nebula Command-line Interface
####################################################

Overview
=========
This video will take you on a tour of Nebulactl - how to install and configure it, as well as how to use the Verbs and Nouns sections on the left hand side menu. Detailed information can be found in the sections below the video.

.. youtube:: cV8ezYnBANE


Installation
============

Nebulactl is a Golang binary that can be installed on any platform supported by Golang.

.. tabbed:: OSX

  .. prompt:: bash $

      brew install nebulaclouds/homebrew-tap/nebulactl

  *Upgrade* existing installation using the following command:

  .. prompt:: bash $

      nebulactl upgrade

.. tabbed:: Other Operating systems

  .. prompt:: bash $

      curl -sL https://ctl.nebula.org/install | bash

  *Upgrade* existing installation using the following command:

  .. prompt:: bash $

      nebulactl upgrade

**Test** if Nebulactl is installed correctly (your Nebulactl version should be > 0.2.0) using the following command:

.. prompt:: bash $

  nebulactl version

Configuration
=============

Nebulactl allows you to communicate with NebulaAdmin using a YAML file or by passing every configuration value
on the command-line. The following configuration can be used for the setup:

Basic Configuration
--------------------

The full list of available configurable options can be found by running ``nebulactl --help``, or `here <https://docs.nebula.org/projects/nebulactl/en/stable/gen/nebulactl.html#synopsis>`__.

.. NOTE::

    Currently, the Project ``-p``, Domain ``-d``, and Output ``-o`` flags cannot be used in the config file.

.. tabbed:: Local Nebula Sandbox

    Automatically configured for you by ``nebulactl sandbox`` command.

    .. code-block:: yaml

        admin:
          # For GRPC endpoints you might want to use dns:///nebula.myexample.com
          endpoint: dns:///localhost:30081
          insecure: true # Set to false to enable TLS/SSL connection (not recommended except on local sandbox deployment).
          authType: Pkce # authType: Pkce # if using authentication or just drop this.

.. tabbed:: AWS Configuration

    .. code-block:: yaml

        admin:
          # For GRPC endpoints you might want to use dns:///nebula.myexample.com
          endpoint: dns:///<replace-me>
          authType: Pkce # authType: Pkce # if using authentication or just drop this.
          insecure: true # insecure: True # Set to true if the endpoint isn't accessible through TLS/SSL connection (not recommended except on local sandbox deployment)

.. tabbed:: GCS Configuration

    .. code-block:: yaml

        admin:
          # For GRPC endpoints you might want to use dns:///nebula.myexample.com
          endpoint: dns:///<replace-me>
          authType: Pkce # authType: Pkce # if using authentication or just drop this.
          insecure: false # insecure: True # Set to true if the endpoint isn't accessible through TLS/SSL connection (not recommended except on local sandbox deployment)

.. tabbed:: Others

    For other supported storage backends like Oracle, Azure, etc., refer to the configuration structure `here <https://pkg.go.dev/github.com/nebulaclouds/nebula/nebulastdlib/storage#Config>`__.

    Place the config file in ``$HOME/.nebula`` directory with the name config.yaml.
    This file is typically searched in:

    * ``$HOME/.nebula``
    * currDir from where you run nebulactl
    * ``/etc/nebula/config``
    
    You can also pass the file name in the command line using ``--config <config-file-path>``.
