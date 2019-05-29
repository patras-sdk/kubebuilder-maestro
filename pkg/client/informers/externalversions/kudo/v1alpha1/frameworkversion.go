/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kudov1alpha1 "github.com/kudobuilder/kudo/pkg/apis/kudo/v1alpha1"
	versioned "github.com/kudobuilder/kudo/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kudobuilder/kudo/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kudobuilder/kudo/pkg/client/listers/kudo/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FrameworkVersionInformer provides access to a shared informer and lister for
// FrameworkVersions.
type FrameworkVersionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FrameworkVersionLister
}

type frameworkVersionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFrameworkVersionInformer constructs a new informer for FrameworkVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFrameworkVersionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFrameworkVersionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFrameworkVersionInformer constructs a new informer for FrameworkVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFrameworkVersionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KudoV1alpha1().FrameworkVersions(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KudoV1alpha1().FrameworkVersions(namespace).Watch(options)
			},
		},
		&kudov1alpha1.FrameworkVersion{},
		resyncPeriod,
		indexers,
	)
}

func (f *frameworkVersionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFrameworkVersionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *frameworkVersionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kudov1alpha1.FrameworkVersion{}, f.defaultInformer)
}

func (f *frameworkVersionInformer) Lister() v1alpha1.FrameworkVersionLister {
	return v1alpha1.NewFrameworkVersionLister(f.Informer().GetIndexer())
}
