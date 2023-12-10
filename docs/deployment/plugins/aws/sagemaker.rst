.. _deployment-plugin-setup-aws-sagemaker:

Sagemaker Plugin Setup
----------------------

This guide gives an overview of how to set up Sagemaker in your Nebula deployment.

.. note::
   
   The Sagemaker plugin needs Nebula deployment in AWS cloud; sandbox/GCP/Azure
   won't work.

Setup the AWS Nebula cluster
===========================

.. tabs::

   .. tab:: AWS cluster setup
   
     * Make sure you have up and running nebula cluster in `AWS <https://docs.nebula.org/en/latest/deployment/aws/index.html#deployment-aws>`__
     * You have your `AWS role set up correctly for SageMaker <https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html>`_
     * `AWS SageMaker k8s operator <https://github.com/aws/amazon-sagemaker-operator-for-k8s>`_ is installed in your k8s cluster
     * Make sure you have correct kubeconfig and selected the correct kubernetes context
     * make sure you have the correct NebulaCTL config at ~/.nebula/config.yaml

Specify Plugin Configuration
======================================

Create a file named ``values-override.yaml`` and add the following config to it.
Please make sure that the propeller has the correct service account for Sagemaker.

.. code-block:: yaml

    configmap:
      enabled_plugins:
        # -- Tasks specific configuration [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#GetConfig)
        tasks:
          # -- Plugins configuration, [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#TaskPluginConfig)
          task-plugins:
            # -- [Enabled Plugins](https://pkg.go.dev/github.com/nebulaclouds/nebulaplugins/go/tasks/config#Config).
            # plugins
            enabled-plugins:
              - container
              - sidecar
              - k8s-array
              - sagemaker_training
              - sagemaker_hyperparameter_tuning
            default-for-task-types:
              container: container
              sidecar: sidecar
              container_array: k8s-array

Upgrade the Nebula Helm release
==============================

.. prompt:: bash $

  helm upgrade -n nebula -f values-override.yaml nebulaclouds/nebula-core


Register the Sagemaker plugin example
=====================================

.. prompt:: bash $

  nebulactl register files https://github.com/nebulaclouds/nebulasnacks/releases/download/v0.3.0/snacks-cookbook-integrations-aws-sagemaker_training.tar.gz --archive -p nebulasnacks -d development


Launch an execution
===================

.. tabs::

   .. tab:: Nebula Console

     * Navigate to Nebula Console's UI (e.g. `sandbox <http://localhost:30081/console>`_) and find the workflow.
     * Click on `Launch` to open up the launch form.
     * Submit the form.

   .. tab:: Nebulactl
   
     Retrieve an execution form in the form of a YAML file:
   
     .. code-block:: bash
   
        nebulactl get launchplan --config ~/.nebula/nebulactl.yaml \
           --project nebulasnacks \
           --domain development \
           sagemaker_training.sagemaker_custom_training.mnist_trainer \
           --latest \
           --execFile exec_spec.yaml
   
     Launch! 🚀
   
     .. code-block:: bash
   
        nebulactl --config ~/.nebula/nebulactl.yaml create execution \
            -p <project> -d <domain> --execFile ~/exec_spec.yaml
