/*
Copyright The Kubernetes Authors.

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

package v1

import (
	time "time"

	cronhpav1 "github.com/iyacontrol/shareit/pkg/apis/cronhpa/v1"
	versioned "github.com/iyacontrol/shareit/pkg/client/cronhpa/clientset/versioned"
	internalinterfaces "github.com/iyacontrol/shareit/pkg/client/cronhpa/informers/externalversions/internalinterfaces"
	v1 "github.com/iyacontrol/shareit/pkg/client/cronhpa/listers/cronhpa/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CronHpaInformer provides access to a shared informer and lister for
// CronHpas.
type CronHpaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CronHpaLister
}

type cronHpaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCronHpaInformer constructs a new informer for CronHpa type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCronHpaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCronHpaInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCronHpaInformer constructs a new informer for CronHpa type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCronHpaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CronhpaV1().CronHpas(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CronhpaV1().CronHpas(namespace).Watch(options)
			},
		},
		&cronhpav1.CronHpa{},
		resyncPeriod,
		indexers,
	)
}

func (f *cronHpaInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCronHpaInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cronHpaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cronhpav1.CronHpa{}, f.defaultInformer)
}

func (f *cronHpaInformer) Lister() v1.CronHpaLister {
	return v1.NewCronHpaLister(f.Informer().GetIndexer())
}
