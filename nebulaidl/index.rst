.. nebulaidl documentation master file, created by

Nebulaidl: Nebula's Core Language Specification
==============================================

``Nebulaidl`` contains the core language specification of Nebula, using `Google's Protocol Buffers <https://developers.google.com/protocol-buffers>`_.
The Specification contains:

#. The core specification for Nebula workflows, tasks, and the type system
#. The specification for NebulaAdmin's `gRPC <https://grpc.io/>`_ and ``REST`` endpoints
#. Some of the core plugin APIs like - Spark, Sagemaker, etc.

This specification is used to generate client stubs for `Nebulakit <https://nebula.readthedocs.io/projects/nebulakit>`_, `Nebulakit Java <https://github.com/spotify/nebulakit-java>`_, `Nebulactl <https://github.com/nebulaclouds/nebulactl>`_ and the `NebulaAdmin Service <https://pkg.go.dev/github.com/lyft/nebulaadmin>`_.


.. toctree::
   :maxdepth: 1
   :hidden:

   |plane| Getting Started <https://docs.nebula.org/en/latest/getting_started.html>
   |book-reader| User Guide <https://docs.nebula.org/projects/cookbook/en/latest/user_guide.html>
   |chalkboard| Tutorials <https://docs.nebula.org/projects/cookbook/en/latest/tutorials.html>
   |project-diagram| Concepts <https://docs.nebula.org/en/latest/concepts/basics.html>
   |rocket| Deployment <https://docs.nebula.org/en/latest/deployment/index.html>
   |book| API Reference <https://docs.nebula.org/en/latest/reference/index.html>
   |hands-helping| Community <https://docs.nebula.org/en/latest/community/index.html>

.. NOTE: the caption text is important for the sphinx theme to correctly render the nav header
.. https://github.com/nebulaclouds/furo
.. toctree::
   :maxdepth: -1
   :caption: NebulaIDL
   :hidden:

   Nebulaidl <self>
   protos/index
   Contributing Guide <README>
