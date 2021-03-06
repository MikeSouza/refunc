/*
Copyright 2018 The refunc Authors

TODO: choose a opensource licence.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta3 "github.com/refunc/refunc/pkg/apis/refunc/v1beta3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeXenvs implements XenvInterface
type FakeXenvs struct {
	Fake *FakeRefuncV1beta3
	ns   string
}

var xenvsResource = schema.GroupVersionResource{Group: "refunc.refunc.io", Version: "v1beta3", Resource: "xenvs"}

var xenvsKind = schema.GroupVersionKind{Group: "refunc.refunc.io", Version: "v1beta3", Kind: "Xenv"}

// Get takes name of the xenv, and returns the corresponding xenv object, and an error if there is any.
func (c *FakeXenvs) Get(name string, options v1.GetOptions) (result *v1beta3.Xenv, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(xenvsResource, c.ns, name), &v1beta3.Xenv{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.Xenv), err
}

// List takes label and field selectors, and returns the list of Xenvs that match those selectors.
func (c *FakeXenvs) List(opts v1.ListOptions) (result *v1beta3.XenvList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(xenvsResource, xenvsKind, c.ns, opts), &v1beta3.XenvList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta3.XenvList{ListMeta: obj.(*v1beta3.XenvList).ListMeta}
	for _, item := range obj.(*v1beta3.XenvList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested xenvs.
func (c *FakeXenvs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(xenvsResource, c.ns, opts))

}

// Create takes the representation of a xenv and creates it.  Returns the server's representation of the xenv, and an error, if there is any.
func (c *FakeXenvs) Create(xenv *v1beta3.Xenv) (result *v1beta3.Xenv, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(xenvsResource, c.ns, xenv), &v1beta3.Xenv{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.Xenv), err
}

// Update takes the representation of a xenv and updates it. Returns the server's representation of the xenv, and an error, if there is any.
func (c *FakeXenvs) Update(xenv *v1beta3.Xenv) (result *v1beta3.Xenv, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(xenvsResource, c.ns, xenv), &v1beta3.Xenv{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.Xenv), err
}

// Delete takes name of the xenv and deletes it. Returns an error if one occurs.
func (c *FakeXenvs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(xenvsResource, c.ns, name), &v1beta3.Xenv{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeXenvs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(xenvsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta3.XenvList{})
	return err
}

// Patch applies the patch and returns the patched xenv.
func (c *FakeXenvs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta3.Xenv, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(xenvsResource, c.ns, name, data, subresources...), &v1beta3.Xenv{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta3.Xenv), err
}
