/*
Copyright 2018 Giant Swarm GmbH.

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

package v1alpha1

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AzureClusterConfigsGetter has a method to return a AzureClusterConfigInterface.
// A group's client should implement this interface.
type AzureClusterConfigsGetter interface {
	AzureClusterConfigs(namespace string) AzureClusterConfigInterface
}

// AzureClusterConfigInterface has methods to work with AzureClusterConfig resources.
type AzureClusterConfigInterface interface {
	Create(*v1alpha1.AzureClusterConfig) (*v1alpha1.AzureClusterConfig, error)
	Update(*v1alpha1.AzureClusterConfig) (*v1alpha1.AzureClusterConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AzureClusterConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.AzureClusterConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureClusterConfig, err error)
	AzureClusterConfigExpansion
}

// azureClusterConfigs implements AzureClusterConfigInterface
type azureClusterConfigs struct {
	client rest.Interface
	ns     string
}

// newAzureClusterConfigs returns a AzureClusterConfigs
func newAzureClusterConfigs(c *CoreV1alpha1Client, namespace string) *azureClusterConfigs {
	return &azureClusterConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azureClusterConfig, and returns the corresponding azureClusterConfig object, and an error if there is any.
func (c *azureClusterConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureClusterConfig, err error) {
	result = &v1alpha1.AzureClusterConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureclusterconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureClusterConfigs that match those selectors.
func (c *azureClusterConfigs) List(opts v1.ListOptions) (result *v1alpha1.AzureClusterConfigList, err error) {
	result = &v1alpha1.AzureClusterConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureclusterconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureClusterConfigs.
func (c *azureClusterConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azureclusterconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a azureClusterConfig and creates it.  Returns the server's representation of the azureClusterConfig, and an error, if there is any.
func (c *azureClusterConfigs) Create(azureClusterConfig *v1alpha1.AzureClusterConfig) (result *v1alpha1.AzureClusterConfig, err error) {
	result = &v1alpha1.AzureClusterConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azureclusterconfigs").
		Body(azureClusterConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a azureClusterConfig and updates it. Returns the server's representation of the azureClusterConfig, and an error, if there is any.
func (c *azureClusterConfigs) Update(azureClusterConfig *v1alpha1.AzureClusterConfig) (result *v1alpha1.AzureClusterConfig, err error) {
	result = &v1alpha1.AzureClusterConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureclusterconfigs").
		Name(azureClusterConfig.Name).
		Body(azureClusterConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the azureClusterConfig and deletes it. Returns an error if one occurs.
func (c *azureClusterConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureclusterconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureClusterConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureclusterconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched azureClusterConfig.
func (c *azureClusterConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureClusterConfig, err error) {
	result = &v1alpha1.AzureClusterConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azureclusterconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
