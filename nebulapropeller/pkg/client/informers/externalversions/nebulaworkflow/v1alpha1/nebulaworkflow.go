// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	nebulaworkflowv1alpha1 "github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	versioned "github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/listers/nebulaworkflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NebulaWorkflowInformer provides access to a shared informer and lister for
// NebulaWorkflows.
type NebulaWorkflowInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NebulaWorkflowLister
}

type nebulaWorkflowInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNebulaWorkflowInformer constructs a new informer for NebulaWorkflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNebulaWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNebulaWorkflowInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNebulaWorkflowInformer constructs a new informer for NebulaWorkflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNebulaWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NebulaworkflowV1alpha1().NebulaWorkflows(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NebulaworkflowV1alpha1().NebulaWorkflows(namespace).Watch(context.TODO(), options)
			},
		},
		&nebulaworkflowv1alpha1.NebulaWorkflow{},
		resyncPeriod,
		indexers,
	)
}

func (f *nebulaWorkflowInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNebulaWorkflowInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nebulaWorkflowInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&nebulaworkflowv1alpha1.NebulaWorkflow{}, f.defaultInformer)
}

func (f *nebulaWorkflowInformer) Lister() v1alpha1.NebulaWorkflowLister {
	return v1alpha1.NewNebulaWorkflowLister(f.Informer().GetIndexer())
}
