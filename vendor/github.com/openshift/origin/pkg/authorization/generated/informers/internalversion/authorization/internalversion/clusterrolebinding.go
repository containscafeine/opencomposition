// This file was automatically generated by informer-gen

package internalversion

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	internalinterfaces "github.com/openshift/origin/pkg/authorization/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/authorization/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/authorization/generated/listers/authorization/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterRoleBindingInformer provides access to a shared informer and lister for
// ClusterRoleBindings.
type ClusterRoleBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ClusterRoleBindingLister
}

type clusterRoleBindingInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newClusterRoleBindingInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Authorization().ClusterRoleBindings().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Authorization().ClusterRoleBindings().Watch(options)
			},
		},
		&authorization.ClusterRoleBinding{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *clusterRoleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authorization.ClusterRoleBinding{}, newClusterRoleBindingInformer)
}

func (f *clusterRoleBindingInformer) Lister() internalversion.ClusterRoleBindingLister {
	return internalversion.NewClusterRoleBindingLister(f.Informer().GetIndexer())
}