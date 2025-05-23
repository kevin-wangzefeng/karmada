/*
Copyright The Karmada Authors.

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

package v1alpha1

import (
	policyv1alpha1 "github.com/karmada-io/karmada/pkg/apis/policy/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterOverridePolicyLister helps list ClusterOverridePolicies.
// All objects returned here must be treated as read-only.
type ClusterOverridePolicyLister interface {
	// List lists all ClusterOverridePolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*policyv1alpha1.ClusterOverridePolicy, err error)
	// Get retrieves the ClusterOverridePolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*policyv1alpha1.ClusterOverridePolicy, error)
	ClusterOverridePolicyListerExpansion
}

// clusterOverridePolicyLister implements the ClusterOverridePolicyLister interface.
type clusterOverridePolicyLister struct {
	listers.ResourceIndexer[*policyv1alpha1.ClusterOverridePolicy]
}

// NewClusterOverridePolicyLister returns a new ClusterOverridePolicyLister.
func NewClusterOverridePolicyLister(indexer cache.Indexer) ClusterOverridePolicyLister {
	return &clusterOverridePolicyLister{listers.New[*policyv1alpha1.ClusterOverridePolicy](indexer, policyv1alpha1.Resource("clusteroverridepolicy"))}
}
