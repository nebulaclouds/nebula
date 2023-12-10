/*
Package manager introduces a NebulaPropeller Manager implementation that enables horizontal scaling of NebulaPropeller by sharding NebulaWorkflows.

The NebulaPropeller Manager manages a collection of NebulaPropeller instances to effectively distribute load. Each managed NebulaPropller instance is created as a k8s pod using a configurable k8s PodTemplate resource. The NebulaPropeller Manager use a control loop to periodically check the status of managed NebulaPropeller instances and creates, updates, or deletes pods as required. It is important to note that if the NebulaPropeller Manager fails, managed instances are left running. This is in effort to ensure progress continues in evaluating NebulaWorkflow CRDs.

NebulaPropeller Manager is configured at the root of the NebulaPropeller configuration. Below is an example of the variety of configuration options along with succinct associated descriptions for each field:

	manager:
	  pod-application: "nebulapropeller"             # application name for managed pods
	  pod-template-container-name: "nebulapropeller" # the container name within the K8s PodTemplate name used to set NebulaWorkflow CRD labels selectors
	  pod-template-name: "nebulapropeller-template"  # k8s PodTemplate name to use for starting NebulaPropeller pods
	  pod-template-namespace: "nebula"               # namespace where the k8s PodTemplate is located
	  scan-interval: 10s                            # frequency to scan NebulaPropeller pods and start / restart if necessary
	  shard:                                        # configure sharding strategy
	    # shard configuration redacted

NebulaPropeller Manager handles dynamic updates to both the k8s PodTemplate and shard configuration. The k8s PodTemplate resource has an associated resource version which uniquely identifies changes. Additionally, shard configuration modifications may be tracked using a simple hash. Nebula stores these values as annotations on managed NebulaPropeller instances. Therefore, if either of there values change the NebulaPropeller Manager instance will detect it and perform the necessary deployment updates.

# Shard Strategies

Nebula defines a variety of Shard Strategies for configuring how NebulaWorkflows are sharded. These options may include the shard type (ex. hash, project, or domain) along with the number of shards or the distribution of project / domain IDs over shards.

Internally, NebulaWorkflow CRDs are initialized with k8s labels for project, domain, and a shard-key. The project and domain label values are associated with the environment of the registered workflow. The shard-key value is a range-bounded hash over various components of the NebulaWorkflow metadata, currently the keyspace range is defined as [0,32). A sharded Nebula deployment ensures deterministic NebulaWorkflow evaluations by setting disjoint k8s label selectors, based on the aforementioned labels, on each managed NebulaPropeller instance. This ensures that only a single NebulaPropeller instance is responsible for processing each NebulaWorkflow.

The Hash Shard Strategy, denoted by "type: hash" in the configuration below, uses consistent hashing to evenly distribute NebulaWorkflows over managed NebulaPropeller instances. This is achieved by partitioning the keyspace (i.e. [0,32)) into a collection of disjoint ranges and using label selectors to assign those ranges to managed NebulaPropeller instances. For example, with "shard-count: 4" the first instance is responsible for NebulaWorkflows with "shard-keys" in the range [0,8), the second [8,16), the third [16,24), and the fourth [24,32). It may be useful to note that the default shard type is "hash", so it will be implicitly defined if otherwise left out of the configuration. An example configuration for the Hash Shard Strategy is provided below:

	# a configuration example using the "hash" shard type
	manager:
	  # pod and scanning configuration redacted
	  shard:
	    type: hash     # use the "hash" shard strategy
	    shard-count: 4 # the total number of shards

The Project and Domain Shard Strategies, denoted by "type: project" and "type: domain" respectively, use the NebulaWorkflow project and domain metadata to distributed NebulaWorkflows over managed NebulaPropeller instances. These Shard Strategies are configured using a "per-shard-mapping" option, which is a list of ID lists. Each element in the "per-shard-mapping" list defines a new shard and the ID list assigns responsibility for the specified IDs to that shard. The assignment is performed using k8s label selectors, where each managed NebulaPropeller instance includes NebulaWorkflows with the specified project or domain labels.

A shard configured as a single wildcard ID (i.e. "*") is responsible for all IDs that are not covered by other shards. Only a single shard may be configured with a wildcard ID and on that shard their must be only one ID, namely the wildcard. In this case, the managed NebulaPropeller instance uses k8s label selectors to exclude NebulaWorkflows with project or domain IDs from other shards.

	# a configuration example using the "project" shard type
	manager:
	  # pod and scanning configuration redacted
	  shard:
	    type: project       # use the "project" shard strategy
	    per-shard-mapping:  # a list of per shard mappings - one shard is created for each element
	      - ids:            # the list of ids to be managed by the first shard
	        - nebulasnacks
	      - ids:            # the list of ids to be managed by the second shard
	        - nebulaexamples
	        - nebulalabs
	      - ids:            # the list of ids to be managed by the third shard
	        - "*"           # use the wildcard to manage all ids not managed by other shards

	# a configuration example using the "domain" shard type
	manager:
	  # pod and scanning configuration redacted
	  shard:
	    type: domain        # use the "domain" shard strategy
	    per-shard-mapping:  # a list of per shard mappings - one shard is created for each element
	      - ids:            # the list of ids to be managed by the first shard
	        - production
	      - ids:            # the list of ids to be managed by the second shard
	        - "*"           # use the wildcard to manage all ids not managed by other shards
*/
package manager
