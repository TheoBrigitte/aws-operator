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

// ChartConfigsGetter has a method to return a ChartConfigInterface.
// A group's client should implement this interface.
type ChartConfigsGetter interface {
	ChartConfigs(namespace string) ChartConfigInterface
}

// ChartConfigInterface has methods to work with ChartConfig resources.
type ChartConfigInterface interface {
	Create(*v1alpha1.ChartConfig) (*v1alpha1.ChartConfig, error)
	Update(*v1alpha1.ChartConfig) (*v1alpha1.ChartConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ChartConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.ChartConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ChartConfig, err error)
	ChartConfigExpansion
}

// chartConfigs implements ChartConfigInterface
type chartConfigs struct {
	client rest.Interface
	ns     string
}

// newChartConfigs returns a ChartConfigs
func newChartConfigs(c *CoreV1alpha1Client, namespace string) *chartConfigs {
	return &chartConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the chartConfig, and returns the corresponding chartConfig object, and an error if there is any.
func (c *chartConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.ChartConfig, err error) {
	result = &v1alpha1.ChartConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chartconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ChartConfigs that match those selectors.
func (c *chartConfigs) List(opts v1.ListOptions) (result *v1alpha1.ChartConfigList, err error) {
	result = &v1alpha1.ChartConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chartconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested chartConfigs.
func (c *chartConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("chartconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a chartConfig and creates it.  Returns the server's representation of the chartConfig, and an error, if there is any.
func (c *chartConfigs) Create(chartConfig *v1alpha1.ChartConfig) (result *v1alpha1.ChartConfig, err error) {
	result = &v1alpha1.ChartConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("chartconfigs").
		Body(chartConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a chartConfig and updates it. Returns the server's representation of the chartConfig, and an error, if there is any.
func (c *chartConfigs) Update(chartConfig *v1alpha1.ChartConfig) (result *v1alpha1.ChartConfig, err error) {
	result = &v1alpha1.ChartConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chartconfigs").
		Name(chartConfig.Name).
		Body(chartConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the chartConfig and deletes it. Returns an error if one occurs.
func (c *chartConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chartconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *chartConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chartconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched chartConfig.
func (c *chartConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ChartConfig, err error) {
	result = &v1alpha1.ChartConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("chartconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
