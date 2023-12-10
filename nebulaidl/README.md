# Nebulaidl

This is one of the core repositories of Nebula. It contains the Specification of the Nebula Language using protobuf messages, the Backend API specification in gRPC, and Swagger REST. The repo contains the generated clients and protocol message structures in multiple languages. Along with the generated code, the repository also contains the Golang clients for Nebula's backend APIs (the services grouped under NebulaAdmin).


[![Slack](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://slack.nebula.org)

* [nebula.org](https://nebula.org)
* [Nebula Docs](http://docs.nebula.org)
* [Nebulaidl API reference documentation](https://docs.nebula.org/projects/nebulaidl/en/stable/index.html)

## Contributing to Nebulaidl

## Tooling for Nebulaidl

1. Run ``make download_tooling`` to install generator dependencies.

```bash
   make download_tooling
```

2. Ensure Docker is installed locally.
3. Run ``make generate`` to generate all the code, mock client, and docs for NebulaAdmin Service.

```bash
    make generate
```

4. To add new dependencies for documentation generation, modify ``doc-requirements.in`` and run

```bash
   make doc-requirements.txt
```

## Docs structure

The index.rst files for protos are arranged in parallel under the ``docs`` folder.
All the proto definitions are within ``protos/nebulaidl`` and their corresponding docs are in ``protos/docs``.

```
docs
├── admin
│   ├── admin.rst
│   └── index.rst
├── core
│   ├── core.rst
│   └── index.rst
├── datacatalog
│   ├── datacatalog.rst
│   └── index.rst
├── event
│   ├── event.rst
│   └── index.rst
├── plugins
│   ├── index.rst
│   └── plugins.rst
├── service
│   ├── index.rst
│   └── service.rst
```

Each module in protos has a module in docs with the same name.
For example: ``protos/nebulaidl/core`` has a module ``protos/docs/core`` under the ``docs`` folder which has the corresponding index and documentation files.


## Generating Documentation

* If a new module is to be introduced, follow the structure for core files in `generate_protos.sh` file which helps generate the core documentation from its proto files.
```
     core_proto_files=`ls protos/nebulaidl/core/*.proto |xargs`
     # Remove any currently generated file
     ls -d protos/docs/core/* | grep -v index.rst | xargs rm
     protoc --doc_out=protos/docs/core --doc_opt=restructuredtext,core.rst -I=protos `echo $core_proto_files`
```

* ``make generate`` generates the modified rst files.

* ``make html`` generates the Sphinx documentation from the docs folder that uses the modified rst files.

