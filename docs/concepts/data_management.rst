.. _divedeep-data-management:

#################################
Understand How Nebula Handles Data
#################################

.. tags:: Basic, Glossary, Design

Types of Data
=============

There are two parts to the data in Nebula:

1. Metadata

* It consists of data about inputs to a task, and other artifacts.
* It is configured globally for NebulaPropeller, NebulaAdmin etc., and the running pods/jobs need access to this bucket to get the data.

2. Raw data

* It is the actual data (such as the Pandas DataFrame, Spark DataFrame, etc.).
* Raw data paths are unique for every execution, and the prefixes can be modified per execution.
* None of the Nebula control plane components would access the raw data. This provides great separation of data between the control plane and the data plane.

.. note:
  Metadata and raw data can be present in entirely separate buckets.


Let us consider a simple Python task:

.. code-block:: python

   @task
   def my_task(m: int, n: str, o: NebulaFile) -> pd.DataFrame:
      ...

In the above code sample, ``m``, ``n``, ``o`` are inputs to the task.
``m`` of type ``int`` and ``n`` of type ``str`` are simple primitive types, while ``o`` is an arbitrarily sized file.
All of them from Nebula's point of view are ``data``.
The difference lies in how Nebula stores and passes each of these data items.

For every task that receives input, Nebula sends an **Inputs Metadata** object, which contains all the primitive or simple scalar values inlined, but in the case of
complex, large objects, they are offloaded and the `Metadata` simply stores a reference to the object. In our example, ``m`` and ``n`` are inlined while
``o`` and the output ``pd.DataFrame`` are offloaded to an object store, and their reference is captured in the metadata.

`Nebulakit TypeTransformers` make it possible to use complex objects as if they are available locally - just like persistent filehandles. But Nebula backend only deals with
the references.

Thus, primitive data types and references to large objects fall under Metadata - `Meta input` or `Meta output`, and the actual large object is known as **Raw data**.
A unique property of this separation is that all `meta values` are read by NebulaPropeller engine and available on the NebulaConsole or CLI from the control plane.
`Raw` data is not read by any of the Nebula components and hence it is possible to store it in a completely separate blob storage or alternate stores, which can't be accessed by Nebula control plane components
but can be accessed by users's container/tasks.

Raw Data Prefix
~~~~~~~~~~~~~~~

Every task can read/write its own data files. If ``NebulaFile`` or any natively supported type like ``pandas.DataFrame`` is used, Nebula will automatically offload and download
data from the configured object-store paths. These paths are completely customizable per `LaunchPlan` or `Execution`.

- The default Rawoutput path (prefix in an object store like S3/GCS) can be configured during registration as shown in :std:ref:`nebulactl_register_files`.
  The argument ``--outputLocationPrefix`` allows us to set the destination directory for all the raw data produced. Nebula will create randomized folders in this path to store the data.
- To override the ``RawOutput`` path (prefix in an object store like S3/GCS), you can specify an alternate location when invoking a Nebula execution, as shown in the following screenshot of the LaunchForm in NebulaConsole:

  .. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/concepts/data_movement/launch_raw_output.png

- In the sandbox, the default Rawoutput-prefix is configured to be the root of the local bucket. Hence Nebula will write all the raw data (reference types like blob, file, df/schema/parquet, etc.) under a path defined by the execution.


Metadata
~~~~~~~~

Metadata in Nebula is critical to enable the passing of data between tasks. It allows to perform in-memory computations for branches or send partial outputs from one task to another or compose outputs from multiple tasks into one input to be sent to a task.

Thus, metadata is restricted due to its omnipresence. Each `meta output`/`input` cannot be larger than 1MB. If you have `List[int]`, it cannot be larger than 1MB, considering other input entities. In scenarios where large lists or strings need to be sent between tasks, file abstraction is preferred.

``LiteralType`` & Literals
~~~~~~~~~~~~~~~~~~~~~~~~~~

SERIALIZATION TIME
^^^^^^^^^^^^^^^^^^

When a task is declared with inputs and outputs, Nebula extracts the interface of the task and converts it to an internal representation called a :std:ref:`ref_nebulaidl.core.typedinterface`.
For each variable, a corresponding :std:ref:`ref_nebulaidl.core.literaltype` is created.

For example, the following Python function's interface is transformed as follows:

.. code-block:: python

    @task
    def my_task(a: int, b: str) -> NebulaFile:
        """
        Description of my function

        :param a: My input integer
        :param b: My input string
        :return: My output file
        """
        ...

.. code-block::

    interface {
    inputs {
      variables {
        key: "a"
        value {
          type {
            simple: INTEGER
          }
          description: "My input Integer"
        }
      }
      variables {
        key: "b"
        value {
          type {
            simple: STRING
          }
          description: "My input string"
        }
      }
    }
    outputs {
      variables {
        key: "o0"
        value {
          type {
            blob {
            }
          }
          description: "My output File"
        }
      }
    }
  }


RUNTIME
^^^^^^^

At runtime, data passes through Nebula using :std:ref:`ref_nebulaidl.core.literal` where the values are set.
For files, the corresponding ``Literal`` is called ``LiteralBlob`` (:std:ref:`ref_nebulaidl.core.blob`) which is a binary large object.
Many different objects can be mapped to the underlying `Blob` or `Struct` types. For example, an image is a Blob, a ``pandas.DataFrame`` is a Blob of type parquet, etc.

Data Movement
=============

Nebula is primarily a **DataFlow Engine**. It enables movement of data and provides an abstraction to enable movement of data between different languages.

One implementation of Nebula is the current workflow engine.

The workflow engine is responsible for moving data from a previous task to the next task. As explained previously, Nebula only deals with Metadata and not the actual Raw data.
The illustration below explains how data flows from engine to the task and how that is transferred between tasks. The medium to transfer the data can change, and will change in the future.
We could use fast metadata stores to speed up data movement or exploit locality.

Between Nebulapropeller and Tasks
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/concepts/data_movement/nebula_data_movement.png


Between Tasks
~~~~~~~~~~~~~~

.. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/concepts/data_movement/nebula_data_transfer.png


Bringing in Your Own Datastores for Raw Data
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Nebulakit has a pluggable `data persistence layer <https://docs.nebula.org/projects/nebulakit/en/latest/data.extend.html>`__.
This is driven by PROTOCOL.
For example, it is theoretically possible to use S3 ``s3://`` for metadata and GCS ``gcs://`` for raw data. It is also possible to create your own protocol ``my_fs://``, to change how data is stored and accessed.
But for Metadata, the data should be accessible to Nebula control plane.

Data persistence is also pluggable. By default, it supports all major blob stores and uses an interface defined in Nebulastdlib.
