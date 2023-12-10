Nebula Golang Compile
~~~~~~~~~~~~~~~~~~~~

Common compile script for Nebula golang services.

**To Enable:**

Add ``nebulaclouds/nebula_golang_compile`` to your ``boilerplate/update.cfg`` file.

Add the following to your Makefile

::

  .PHONY: compile_linux
  compile_linux:
    PACKAGES={{ *your packages }} OUTPUT={{ /path/to/output }} ./boilerplate/nebula/nebula_golang_compile.sh
