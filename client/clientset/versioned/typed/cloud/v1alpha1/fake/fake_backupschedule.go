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

	v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/cloud/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackupSchedules implements BackupScheduleInterface
type FakeBackupSchedules struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var backupschedulesResource = schema.GroupVersionResource{Group: "cloud.mongodbatlas.kubeform.com", Version: "v1alpha1", Resource: "backupschedules"}

var backupschedulesKind = schema.GroupVersionKind{Group: "cloud.mongodbatlas.kubeform.com", Version: "v1alpha1", Kind: "BackupSchedule"}

// Get takes name of the backupSchedule, and returns the corresponding backupSchedule object, and an error if there is any.
func (c *FakeBackupSchedules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupschedulesResource, c.ns, name), &v1alpha1.BackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSchedule), err
}

// List takes label and field selectors, and returns the list of BackupSchedules that match those selectors.
func (c *FakeBackupSchedules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackupScheduleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupschedulesResource, backupschedulesKind, c.ns, opts), &v1alpha1.BackupScheduleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackupScheduleList{ListMeta: obj.(*v1alpha1.BackupScheduleList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackupScheduleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupSchedules.
func (c *FakeBackupSchedules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupschedulesResource, c.ns, opts))

}

// Create takes the representation of a backupSchedule and creates it.  Returns the server's representation of the backupSchedule, and an error, if there is any.
func (c *FakeBackupSchedules) Create(ctx context.Context, backupSchedule *v1alpha1.BackupSchedule, opts v1.CreateOptions) (result *v1alpha1.BackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupschedulesResource, c.ns, backupSchedule), &v1alpha1.BackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSchedule), err
}

// Update takes the representation of a backupSchedule and updates it. Returns the server's representation of the backupSchedule, and an error, if there is any.
func (c *FakeBackupSchedules) Update(ctx context.Context, backupSchedule *v1alpha1.BackupSchedule, opts v1.UpdateOptions) (result *v1alpha1.BackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupschedulesResource, c.ns, backupSchedule), &v1alpha1.BackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSchedule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupSchedules) UpdateStatus(ctx context.Context, backupSchedule *v1alpha1.BackupSchedule, opts v1.UpdateOptions) (*v1alpha1.BackupSchedule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupschedulesResource, "status", c.ns, backupSchedule), &v1alpha1.BackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSchedule), err
}

// Delete takes name of the backupSchedule and deletes it. Returns an error if one occurs.
func (c *FakeBackupSchedules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupschedulesResource, c.ns, name), &v1alpha1.BackupSchedule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupSchedules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupschedulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackupScheduleList{})
	return err
}

// Patch applies the patch and returns the patched backupSchedule.
func (c *FakeBackupSchedules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupschedulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupSchedule), err
}
