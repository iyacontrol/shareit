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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/iyacontrol/shareit/pkg/apis/confighpa/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConfigHpaLister helps list ConfigHpas.
type ConfigHpaLister interface {
	// List lists all ConfigHpas in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.ConfigHpa, err error)
	// ConfigHpas returns an object that can list and get ConfigHpas.
	ConfigHpas(namespace string) ConfigHpaNamespaceLister
	ConfigHpaListerExpansion
}

// configHpaLister implements the ConfigHpaLister interface.
type configHpaLister struct {
	indexer cache.Indexer
}

// NewConfigHpaLister returns a new ConfigHpaLister.
func NewConfigHpaLister(indexer cache.Indexer) ConfigHpaLister {
	return &configHpaLister{indexer: indexer}
}

// List lists all ConfigHpas in the indexer.
func (s *configHpaLister) List(selector labels.Selector) (ret []*v1beta1.ConfigHpa, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ConfigHpa))
	})
	return ret, err
}

// ConfigHpas returns an object that can list and get ConfigHpas.
func (s *configHpaLister) ConfigHpas(namespace string) ConfigHpaNamespaceLister {
	return configHpaNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConfigHpaNamespaceLister helps list and get ConfigHpas.
type ConfigHpaNamespaceLister interface {
	// List lists all ConfigHpas in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.ConfigHpa, err error)
	// Get retrieves the ConfigHpa from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.ConfigHpa, error)
	ConfigHpaNamespaceListerExpansion
}

// configHpaNamespaceLister implements the ConfigHpaNamespaceLister
// interface.
type configHpaNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConfigHpas in the indexer for a given namespace.
func (s configHpaNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ConfigHpa, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ConfigHpa))
	})
	return ret, err
}

// Get retrieves the ConfigHpa from the indexer for a given namespace and name.
func (s configHpaNamespaceLister) Get(name string) (*v1beta1.ConfigHpa, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("confighpa"), name)
	}
	return obj.(*v1beta1.ConfigHpa), nil
}
