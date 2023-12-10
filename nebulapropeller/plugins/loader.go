// Package plugins facilitates all the plugins that should be loaded by NebulaPropeller
package plugins

import (
	// Common place to import all plugins, so that it can be imported by Singlebinary (nebulalite) or by propeller main
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/array/awsbatch"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/array/k8s"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/hive"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/k8s/dask"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/k8s/kfoperators/mpi"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/k8s/kfoperators/pytorch"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/k8s/kfoperators/tensorflow"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/k8s/pod"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/k8s/ray"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/k8s/spark"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/testing"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/webapi/athena"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/webapi/bigquery"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/webapi/databricks"
	_ "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/plugins/webapi/snowflake"
)
