// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	operatorv1 "github.com/openshift/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeSchedulers implements KubeSchedulerInterface
type FakeKubeSchedulers struct {
	Fake *FakeOperatorV1
}

var kubeschedulersResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "kubeschedulers"}

var kubeschedulersKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "KubeScheduler"}

// Get takes name of the kubeScheduler, and returns the corresponding kubeScheduler object, and an error if there is any.
func (c *FakeKubeSchedulers) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.KubeScheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubeschedulersResource, name), &operatorv1.KubeScheduler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeScheduler), err
}

// List takes label and field selectors, and returns the list of KubeSchedulers that match those selectors.
func (c *FakeKubeSchedulers) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.KubeSchedulerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubeschedulersResource, kubeschedulersKind, opts), &operatorv1.KubeSchedulerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.KubeSchedulerList{ListMeta: obj.(*operatorv1.KubeSchedulerList).ListMeta}
	for _, item := range obj.(*operatorv1.KubeSchedulerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeSchedulers.
func (c *FakeKubeSchedulers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubeschedulersResource, opts))
}

// Create takes the representation of a kubeScheduler and creates it.  Returns the server's representation of the kubeScheduler, and an error, if there is any.
func (c *FakeKubeSchedulers) Create(ctx context.Context, kubeScheduler *operatorv1.KubeScheduler, opts v1.CreateOptions) (result *operatorv1.KubeScheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubeschedulersResource, kubeScheduler), &operatorv1.KubeScheduler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeScheduler), err
}

// Update takes the representation of a kubeScheduler and updates it. Returns the server's representation of the kubeScheduler, and an error, if there is any.
func (c *FakeKubeSchedulers) Update(ctx context.Context, kubeScheduler *operatorv1.KubeScheduler, opts v1.UpdateOptions) (result *operatorv1.KubeScheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubeschedulersResource, kubeScheduler), &operatorv1.KubeScheduler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeScheduler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeSchedulers) UpdateStatus(ctx context.Context, kubeScheduler *operatorv1.KubeScheduler, opts v1.UpdateOptions) (*operatorv1.KubeScheduler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kubeschedulersResource, "status", kubeScheduler), &operatorv1.KubeScheduler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeScheduler), err
}

// Delete takes name of the kubeScheduler and deletes it. Returns an error if one occurs.
func (c *FakeKubeSchedulers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(kubeschedulersResource, name, opts), &operatorv1.KubeScheduler{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeSchedulers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubeschedulersResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.KubeSchedulerList{})
	return err
}

// Patch applies the patch and returns the patched kubeScheduler.
func (c *FakeKubeSchedulers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.KubeScheduler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeschedulersResource, name, pt, data, subresources...), &operatorv1.KubeScheduler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.KubeScheduler), err
}
