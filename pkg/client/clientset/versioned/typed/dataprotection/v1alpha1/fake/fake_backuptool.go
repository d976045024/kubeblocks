/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

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

	v1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackupTools implements BackupToolInterface
type FakeBackupTools struct {
	Fake *FakeDataprotectionV1alpha1
}

var backuptoolsResource = v1alpha1.SchemeGroupVersion.WithResource("backuptools")

var backuptoolsKind = v1alpha1.SchemeGroupVersion.WithKind("BackupTool")

// Get takes name of the backupTool, and returns the corresponding backupTool object, and an error if there is any.
func (c *FakeBackupTools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackupTool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(backuptoolsResource, name), &v1alpha1.BackupTool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupTool), err
}

// List takes label and field selectors, and returns the list of BackupTools that match those selectors.
func (c *FakeBackupTools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackupToolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(backuptoolsResource, backuptoolsKind, opts), &v1alpha1.BackupToolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackupToolList{ListMeta: obj.(*v1alpha1.BackupToolList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackupToolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupTools.
func (c *FakeBackupTools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(backuptoolsResource, opts))
}

// Create takes the representation of a backupTool and creates it.  Returns the server's representation of the backupTool, and an error, if there is any.
func (c *FakeBackupTools) Create(ctx context.Context, backupTool *v1alpha1.BackupTool, opts v1.CreateOptions) (result *v1alpha1.BackupTool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(backuptoolsResource, backupTool), &v1alpha1.BackupTool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupTool), err
}

// Update takes the representation of a backupTool and updates it. Returns the server's representation of the backupTool, and an error, if there is any.
func (c *FakeBackupTools) Update(ctx context.Context, backupTool *v1alpha1.BackupTool, opts v1.UpdateOptions) (result *v1alpha1.BackupTool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(backuptoolsResource, backupTool), &v1alpha1.BackupTool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupTool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupTools) UpdateStatus(ctx context.Context, backupTool *v1alpha1.BackupTool, opts v1.UpdateOptions) (*v1alpha1.BackupTool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(backuptoolsResource, "status", backupTool), &v1alpha1.BackupTool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupTool), err
}

// Delete takes name of the backupTool and deletes it. Returns an error if one occurs.
func (c *FakeBackupTools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(backuptoolsResource, name, opts), &v1alpha1.BackupTool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupTools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(backuptoolsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackupToolList{})
	return err
}

// Patch applies the patch and returns the patched backupTool.
func (c *FakeBackupTools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupTool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(backuptoolsResource, name, pt, data, subresources...), &v1alpha1.BackupTool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackupTool), err
}
