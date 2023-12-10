[All Overlays](./)
# :beta: Google Cloud Platform Overlay

This overlay serves as an example to bootstrap Nebula setup on Google Cloud Platform (GCP). It is not
designed to work out of the box due to the need of GCP resources. Please follow the instruction
below to further configure.

_Hint_: searching `TODO:` through this directory would help to understand what needs to be done.

## Cloud SQL

[Cloud SQL](https://cloud.google.com/sql) is used as persistence layer. To set it up, please
follow standard GCP documentation.

A few things are required for this overlay to function:

* Two databases named as `nebula` and `datacatalog`
* A database user named as `nebula`
* Password of the database user can be added to either to [kustomization.yaml](kustomization.yaml) or you can create a new file and change the secretGenerator tag to use files. (Refer to kustomize documentation)
* Service account(s) associated with `nebulaadmin` and `datacatalog` pods (either as GKE cluster
  service account or through workload identity) should have `Cloud SQL Editor` role

To securely access Cloud SQL instance, [Cloud SQL
Proxy](https://cloud.google.com/sql/docs/postgres/connect-admin-proxy) is launched as a pod sitting
in between Nebula and Cloud SQL instance.

The kustomization files can be found under [cloudsqlproxy](dependencies/cloudsqlproxy/). Please note that one
needs to replace `<project-id>` and `<region>` accordingly in
[dependencies/cloudsqlproxy/deployment.yaml](cloudsqlproxy/deployment.yaml).

## Create GCS Storage
1. Create a GCS bucket named as `nebula` in a GCP project.
1. Replace `<project-id>` in [config/common/storage.yaml](nebula/config/common/storage.yaml) with the GCP project ID and if using a bucket other than Nebula then replace the bucket name too

## nebulaadmin

nebulaadmin configuration is derived from the [single cluster](../../base/single_cluster) overlay, with only modification to [database configuration db.yaml](nebula/config/admin/db.yaml)

If one has followed [Cloud SQL](#cloud-sql) section, there is nothing to be done for database.

**Advanced / OPTIONAL**
1. The default CORS setting in nebulaAdmin allows cross origin requests. A more secure way would be to allow requests only from the expected domain. To do this, you will have to create a new *server.yaml*
similar to [base/single_cluster/headless/config](../../base/single_cluster/headless/config) under config/admin and then set
`server -> security -> allowedOrigins`.

1. nebulaadmin (including metrics endpoint) is exposed as a service using [internal load balancer](https://cloud.google.com/kubernetes-engine/docs/how-to/internal-load-balancing).

## nebulaconsole

[nebulaconsole configmap](console/config.yaml) needs to be updated with nebulaadmin internal load
balancer IP address or the DNS name associated with it if any.

nebulaconsole is exposed as a service using [internal load balancer](https://cloud.google.com/kubernetes-engine/docs/how-to/internal-load-balancing).

## nebulapropeller

nebulapropeller configuration is derived from the [single cluster](../../base/single_cluster) overlay, with only modification to the config for performance tuning and logs
For logs configuration Replace `<project-id>` in [config/propeller/plugins/task_logs.yaml](nebula/config/propeller/plugins/task_logs.yaml) with the GCP project ID

Some important points

* Storage configuration is shared with Admin and Catalog. Ideally in production Propeller should have its own configuration with real high cache size.

* By default, three plugins are enabled:
1. container
2. k8s-array
3. sidecar

* nebulapropeller metrics endpoint is exposed as a service using [internal load balancer](https://cloud.google.com/kubernetes-engine/docs/how-to/internal-load-balancing).

## datacatalog

datacatalog configuration is derived from the [single cluster](../../base/single_cluster) overlay, with only modification to [database configuration db.yaml](nebula/config/datacatalog/db.yaml)

If one has followed [Cloud SQL](#cloud-sql) section, there is nothing to be done for database.

datacatalog metrics endpoint is exposed as a service using [internal load
balancer](https://cloud.google.com/kubernetes-engine/docs/how-to/internal-load-balancing).


## How to build your overlay
To build your overlay there are 2 options
1. Build it in your own repo Example coming soon :construction:
1. hack it in your clone of Nebula repo in place of GCP overlay. In this case just navigate to the root of the repo and run
```bash
$ make kustomize
```
If all goes well a new overlay composite should be generated in [<root>/deployment/gcp/nebula_generated.yaml](../../../deployment/gcp/nebula_generated.yaml)

## Now ship it

``` shell
make
kubectl apply -f deployment/gcp/nebula_generated.yaml
```
