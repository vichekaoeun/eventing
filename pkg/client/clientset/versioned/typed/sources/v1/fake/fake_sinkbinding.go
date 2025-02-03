/*
Copyright 2021 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1 "knative.dev/eventing/pkg/apis/sources/v1"
)

// FakeSinkBindings implements SinkBindingInterface
type FakeSinkBindings struct {
	Fake *FakeSourcesV1
	ns   string
}

var sinkbindingsResource = v1.SchemeGroupVersion.WithResource("sinkbindings")

var sinkbindingsKind = v1.SchemeGroupVersion.WithKind("SinkBinding")

// Get takes name of the sinkBinding, and returns the corresponding sinkBinding object, and an error if there is any.
func (c *FakeSinkBindings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SinkBinding, err error) {
	emptyResult := &v1.SinkBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(sinkbindingsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.SinkBinding), err
}

// List takes label and field selectors, and returns the list of SinkBindings that match those selectors.
func (c *FakeSinkBindings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SinkBindingList, err error) {
	emptyResult := &v1.SinkBindingList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(sinkbindingsResource, sinkbindingsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.SinkBindingList{ListMeta: obj.(*v1.SinkBindingList).ListMeta}
	for _, item := range obj.(*v1.SinkBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sinkBindings.
func (c *FakeSinkBindings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(sinkbindingsResource, c.ns, opts))

}

// Create takes the representation of a sinkBinding and creates it.  Returns the server's representation of the sinkBinding, and an error, if there is any.
func (c *FakeSinkBindings) Create(ctx context.Context, sinkBinding *v1.SinkBinding, opts metav1.CreateOptions) (result *v1.SinkBinding, err error) {
	emptyResult := &v1.SinkBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(sinkbindingsResource, c.ns, sinkBinding, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.SinkBinding), err
}

// Update takes the representation of a sinkBinding and updates it. Returns the server's representation of the sinkBinding, and an error, if there is any.
func (c *FakeSinkBindings) Update(ctx context.Context, sinkBinding *v1.SinkBinding, opts metav1.UpdateOptions) (result *v1.SinkBinding, err error) {
	emptyResult := &v1.SinkBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(sinkbindingsResource, c.ns, sinkBinding, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.SinkBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSinkBindings) UpdateStatus(ctx context.Context, sinkBinding *v1.SinkBinding, opts metav1.UpdateOptions) (result *v1.SinkBinding, err error) {
	emptyResult := &v1.SinkBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(sinkbindingsResource, "status", c.ns, sinkBinding, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.SinkBinding), err
}

// Delete takes name of the sinkBinding and deletes it. Returns an error if one occurs.
func (c *FakeSinkBindings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(sinkbindingsResource, c.ns, name, opts), &v1.SinkBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSinkBindings) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(sinkbindingsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.SinkBindingList{})
	return err
}

// Patch applies the patch and returns the patched sinkBinding.
func (c *FakeSinkBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SinkBinding, err error) {
	emptyResult := &v1.SinkBinding{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(sinkbindingsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.SinkBinding), err
}
