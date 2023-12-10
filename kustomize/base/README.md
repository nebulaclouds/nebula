[Back to main menu](../)
# Base Components for Nebula
These deployments provide individual deployment units of the Nebula Backend.

As a user it might be preferable to use the `single_cluster` deployment base to create an overlay on top of, or directly edit on top of one of the existing overlays.

## To create a new nebula overlay for one K8s cluster
 Start here 
- [Single Cluster Nebula Deployment configuration](./single_cluster)

## To create a completely custom overlay refer to components
1. NebulaAdmin [Deployment](./admindeployment) | [ServiceAccount](./adminserviceaccount)
1. [Core Nebula namespace creation](./namespace)
1. [NebulaPropeller](./propeller) & its [CRD](./wf_crd)
1. [DataCatalog](./datacatalog)
1. [NebulaConsole](./console)
1. [Overall Ingress for Nebula (optional)](./ingress)
1. [Additional plugin components for Nebula using K8s operators](./operators)

