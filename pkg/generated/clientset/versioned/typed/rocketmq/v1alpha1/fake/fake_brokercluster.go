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

package fake

import (
	v1alpha1 "github.com/huanwei/rocketmq-operator/pkg/apis/rocketmq/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBrokerClusters implements BrokerClusterInterface
type FakeBrokerClusters struct {
	Fake *FakeRocketmqV1alpha1
	ns   string
}

var brokerclustersResource = schema.GroupVersionResource{Group: "rocketmq", Version: "v1alpha1", Resource: "brokerclusters"}

var brokerclustersKind = schema.GroupVersionKind{Group: "rocketmq", Version: "v1alpha1", Kind: "BrokerCluster"}

// Get takes name of the brokerCluster, and returns the corresponding brokerCluster object, and an error if there is any.
func (c *FakeBrokerClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.BrokerCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(brokerclustersResource, c.ns, name), &v1alpha1.BrokerCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BrokerCluster), err
}

// List takes label and field selectors, and returns the list of BrokerClusters that match those selectors.
func (c *FakeBrokerClusters) List(opts v1.ListOptions) (result *v1alpha1.BrokerClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(brokerclustersResource, brokerclustersKind, c.ns, opts), &v1alpha1.BrokerClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BrokerClusterList{ListMeta: obj.(*v1alpha1.BrokerClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.BrokerClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested brokerClusters.
func (c *FakeBrokerClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(brokerclustersResource, c.ns, opts))

}

// Create takes the representation of a brokerCluster and creates it.  Returns the server's representation of the brokerCluster, and an error, if there is any.
func (c *FakeBrokerClusters) Create(brokerCluster *v1alpha1.BrokerCluster) (result *v1alpha1.BrokerCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(brokerclustersResource, c.ns, brokerCluster), &v1alpha1.BrokerCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BrokerCluster), err
}

// Update takes the representation of a brokerCluster and updates it. Returns the server's representation of the brokerCluster, and an error, if there is any.
func (c *FakeBrokerClusters) Update(brokerCluster *v1alpha1.BrokerCluster) (result *v1alpha1.BrokerCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(brokerclustersResource, c.ns, brokerCluster), &v1alpha1.BrokerCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BrokerCluster), err
}

// Delete takes name of the brokerCluster and deletes it. Returns an error if one occurs.
func (c *FakeBrokerClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(brokerclustersResource, c.ns, name), &v1alpha1.BrokerCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBrokerClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(brokerclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BrokerClusterList{})
	return err
}

// Patch applies the patch and returns the patched brokerCluster.
func (c *FakeBrokerClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BrokerCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(brokerclustersResource, c.ns, name, data, subresources...), &v1alpha1.BrokerCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BrokerCluster), err
}
