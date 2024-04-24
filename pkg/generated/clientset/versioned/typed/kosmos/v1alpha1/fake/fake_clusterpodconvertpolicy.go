// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/kosmos.io/kosmos/pkg/apis/kosmos/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterPodConvertPolicies implements ClusterPodConvertPolicyInterface
type FakeClusterPodConvertPolicies struct {
	Fake *FakeKosmosV1alpha1
}

var clusterpodconvertpoliciesResource = schema.GroupVersionResource{Group: "kosmos.io", Version: "v1alpha1", Resource: "clusterpodconvertpolicies"}

var clusterpodconvertpoliciesKind = schema.GroupVersionKind{Group: "kosmos.io", Version: "v1alpha1", Kind: "ClusterPodConvertPolicy"}

// Get takes name of the clusterPodConvertPolicy, and returns the corresponding clusterPodConvertPolicy object, and an error if there is any.
func (c *FakeClusterPodConvertPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterPodConvertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterpodconvertpoliciesResource, name), &v1alpha1.ClusterPodConvertPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPodConvertPolicy), err
}

// List takes label and field selectors, and returns the list of ClusterPodConvertPolicies that match those selectors.
func (c *FakeClusterPodConvertPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterPodConvertPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterpodconvertpoliciesResource, clusterpodconvertpoliciesKind, opts), &v1alpha1.ClusterPodConvertPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterPodConvertPolicyList{ListMeta: obj.(*v1alpha1.ClusterPodConvertPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterPodConvertPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPodConvertPolicies.
func (c *FakeClusterPodConvertPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterpodconvertpoliciesResource, opts))
}

// Create takes the representation of a clusterPodConvertPolicy and creates it.  Returns the server's representation of the clusterPodConvertPolicy, and an error, if there is any.
func (c *FakeClusterPodConvertPolicies) Create(ctx context.Context, clusterPodConvertPolicy *v1alpha1.ClusterPodConvertPolicy, opts v1.CreateOptions) (result *v1alpha1.ClusterPodConvertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterpodconvertpoliciesResource, clusterPodConvertPolicy), &v1alpha1.ClusterPodConvertPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPodConvertPolicy), err
}

// Update takes the representation of a clusterPodConvertPolicy and updates it. Returns the server's representation of the clusterPodConvertPolicy, and an error, if there is any.
func (c *FakeClusterPodConvertPolicies) Update(ctx context.Context, clusterPodConvertPolicy *v1alpha1.ClusterPodConvertPolicy, opts v1.UpdateOptions) (result *v1alpha1.ClusterPodConvertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterpodconvertpoliciesResource, clusterPodConvertPolicy), &v1alpha1.ClusterPodConvertPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPodConvertPolicy), err
}

// Delete takes name of the clusterPodConvertPolicy and deletes it. Returns an error if one occurs.
func (c *FakeClusterPodConvertPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterpodconvertpoliciesResource, name, opts), &v1alpha1.ClusterPodConvertPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPodConvertPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterpodconvertpoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterPodConvertPolicyList{})
	return err
}

// Patch applies the patch and returns the patched clusterPodConvertPolicy.
func (c *FakeClusterPodConvertPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterPodConvertPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterpodconvertpoliciesResource, name, pt, data, subresources...), &v1alpha1.ClusterPodConvertPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPodConvertPolicy), err
}