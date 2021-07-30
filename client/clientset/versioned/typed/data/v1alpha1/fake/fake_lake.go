/*
Copyright AppsCode Inc. and Contributors

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

	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/data/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLakes implements LakeInterface
type FakeLakes struct {
	Fake *FakeDataV1alpha1
	ns   string
}

var lakesResource = schema.GroupVersionResource{Group: "data.mongodbatlas.kubeform.com", Version: "v1alpha1", Resource: "lakes"}

var lakesKind = schema.GroupVersionKind{Group: "data.mongodbatlas.kubeform.com", Version: "v1alpha1", Kind: "Lake"}

// Get takes name of the lake, and returns the corresponding lake object, and an error if there is any.
func (c *FakeLakes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Lake, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lakesResource, c.ns, name), &v1alpha1.Lake{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lake), err
}

// List takes label and field selectors, and returns the list of Lakes that match those selectors.
func (c *FakeLakes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LakeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lakesResource, lakesKind, c.ns, opts), &v1alpha1.LakeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LakeList{ListMeta: obj.(*v1alpha1.LakeList).ListMeta}
	for _, item := range obj.(*v1alpha1.LakeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lakes.
func (c *FakeLakes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lakesResource, c.ns, opts))

}

// Create takes the representation of a lake and creates it.  Returns the server's representation of the lake, and an error, if there is any.
func (c *FakeLakes) Create(ctx context.Context, lake *v1alpha1.Lake, opts v1.CreateOptions) (result *v1alpha1.Lake, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lakesResource, c.ns, lake), &v1alpha1.Lake{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lake), err
}

// Update takes the representation of a lake and updates it. Returns the server's representation of the lake, and an error, if there is any.
func (c *FakeLakes) Update(ctx context.Context, lake *v1alpha1.Lake, opts v1.UpdateOptions) (result *v1alpha1.Lake, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lakesResource, c.ns, lake), &v1alpha1.Lake{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lake), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLakes) UpdateStatus(ctx context.Context, lake *v1alpha1.Lake, opts v1.UpdateOptions) (*v1alpha1.Lake, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lakesResource, "status", c.ns, lake), &v1alpha1.Lake{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lake), err
}

// Delete takes name of the lake and deletes it. Returns an error if one occurs.
func (c *FakeLakes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lakesResource, c.ns, name), &v1alpha1.Lake{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLakes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lakesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LakeList{})
	return err
}

// Patch applies the patch and returns the patched lake.
func (c *FakeLakes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Lake, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lakesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Lake{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lake), err
}