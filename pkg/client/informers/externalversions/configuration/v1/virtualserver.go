// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	configurationv1 "github.com/nginxinc/kubernetes-ingress/v3/pkg/apis/configuration/v1"
	versioned "github.com/nginxinc/kubernetes-ingress/v3/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nginxinc/kubernetes-ingress/v3/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/nginxinc/kubernetes-ingress/v3/pkg/client/listers/configuration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VirtualServerInformer provides access to a shared informer and lister for
// VirtualServers.
type VirtualServerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.VirtualServerLister
}

type virtualServerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualServerInformer constructs a new informer for VirtualServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualServerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualServerInformer constructs a new informer for VirtualServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1().VirtualServers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.K8sV1().VirtualServers(namespace).Watch(context.TODO(), options)
			},
		},
		&configurationv1.VirtualServer{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualServerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualServerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualServerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configurationv1.VirtualServer{}, f.defaultInformer)
}

func (f *virtualServerInformer) Lister() v1.VirtualServerLister {
	return v1.NewVirtualServerLister(f.Informer().GetIndexer())
}
