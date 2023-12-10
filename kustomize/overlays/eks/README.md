[All Overlays](./)
# :construction: Amazon EKS deployment

This overlay serves as an example to bootstrap Nebula setup on AWS. It is not
designed to work out of the box due to the need of AWS resources. Please follow the instruction
below to further configure.

_Hint_: searching `TODO:` through this directory would help to understand what needs to be done.

## Amazon RDS / Amazon Aurora

A few things are required for this overlay to function:

* Two databases named as `nebula` and `datacatalog`
* A database user named as `nebula`
* Password of the database user can be added to either to [kustomization.yaml](kustomization.yaml) or you can create a new file and change the secretGenerator tag to use files. (Refer to kustomize documentation)
* Service account(s) associated with `nebulaadmin` and `datacatalog` pods (either as GKE cluster
  service account or through workload identity) should have `Cloud SQL Editor` role

## Create S3 bucket
1. Create a S3 bucket named as `nebula` (if other name replace it next)
1. Replace in [config/common/storage.yaml](nebula/config/common/storage.yaml) if using a bucket other than Nebula then replace the bucket name too

## nebulaadmin

nebulaadmin configuration is derived from the [single cluster](../../base/single_cluster) overlay, with only modification to [database configuration db.yaml](nebula/config/admin/db.yaml)

**Advanced / OPTIONAL**
1. The default CORS setting in nebulaAdmin allows cross origin requests. A more secure way would be to allow requests only from the expected domain. To do this, you will have to create a new *server.yaml*
similar to [base/single_cluster/headless/config](../../base/single_cluster/headless/config) under config/admin and then set
`server -> security -> allowedOrigins`.

## nebulaconsole

[nebulaconsole configmap](console/config.yaml) needs to be updated with nebulaadmin internal load
balancer IP address or the DNS name associated with it if any.

nebulaconsole is exposed as a service using [internal load balancer](https://cloud.google.com/kubernetes-engine/docs/how-to/internal-load-balancing).

## nebulapropeller

nebulapropeller configuration is derived from the [single cluster](../../base/single_cluster) overlay, with only modification to the config for performance tuning and logs
For logs configuration Replace `<project-id>` in [config/propeller/plugins/task_logs.yaml](nebula/config/propeller/plugins/task_logs.yaml) to use CloudWatch

Some important points

* Storage configuration is shared with Admin and Catalog. Ideally in production Propeller should have its own configuration with real high cache size.

* By default, three plugins are enabled:
1. container
2. k8s-array
3. sidecar

## datacatalog

datacatalog configuration is derived from the [single cluster](../../base/single_cluster) overlay, with only modification to [database configuration db.yaml](nebula/config/datacatalog/db.yaml)


## How to build your overlay
To build your overlay there are 2 options
1. Build it in your own repo Example coming soon :construction:
1. hack it in your clone of Nebula repo in place of EKS overlay. In this case just navigate to the root of the repo and run
```bash
$ make kustomize
```
If all goes well a new overlay composite should be generated in [<root>/deployment/eks/nebula_generated.yaml](../../../deployment/eks/nebula_generated.yaml)

## Now ship it

``` shell
make
kubectl apply -f deployment/gcp/nebula_generated.yaml
```
