package v3

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ClusterRegistryGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ClusterRegistry",
	}
	ClusterRegistryResource = metav1.APIResource{
		Name:         "clusterregistries",
		SingularName: "clusterregistry",
		Namespaced:   true,

		Kind: ClusterRegistryGroupVersionKind.Kind,
	}
)

type ClusterRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRegistry
}

type ClusterRegistryHandlerFunc func(key string, obj *ClusterRegistry) error

type ClusterRegistryLister interface {
	List(namespace string, selector labels.Selector) (ret []*ClusterRegistry, err error)
	Get(namespace, name string) (*ClusterRegistry, error)
}

type ClusterRegistryController interface {
	Informer() cache.SharedIndexInformer
	Lister() ClusterRegistryLister
	AddHandler(name string, handler ClusterRegistryHandlerFunc)
	AddClusterScopedHandler(name, clusterName string, handler ClusterRegistryHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ClusterRegistryInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*ClusterRegistry) (*ClusterRegistry, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ClusterRegistry, error)
	Get(name string, opts metav1.GetOptions) (*ClusterRegistry, error)
	Update(*ClusterRegistry) (*ClusterRegistry, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ClusterRegistryList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ClusterRegistryController
	AddHandler(name string, sync ClusterRegistryHandlerFunc)
	AddLifecycle(name string, lifecycle ClusterRegistryLifecycle)
	AddClusterScopedHandler(name, clusterName string, sync ClusterRegistryHandlerFunc)
	AddClusterScopedLifecycle(name, clusterName string, lifecycle ClusterRegistryLifecycle)
}

type clusterRegistryLister struct {
	controller *clusterRegistryController
}

func (l *clusterRegistryLister) List(namespace string, selector labels.Selector) (ret []*ClusterRegistry, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*ClusterRegistry))
	})
	return
}

func (l *clusterRegistryLister) Get(namespace, name string) (*ClusterRegistry, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    ClusterRegistryGroupVersionKind.Group,
			Resource: "clusterRegistry",
		}, name)
	}
	return obj.(*ClusterRegistry), nil
}

type clusterRegistryController struct {
	controller.GenericController
}

func (c *clusterRegistryController) Lister() ClusterRegistryLister {
	return &clusterRegistryLister{
		controller: c,
	}
}

func (c *clusterRegistryController) AddHandler(name string, handler ClusterRegistryHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}
		return handler(key, obj.(*ClusterRegistry))
	})
}

func (c *clusterRegistryController) AddClusterScopedHandler(name, cluster string, handler ClusterRegistryHandlerFunc) {
	c.GenericController.AddHandler(name, func(key string) error {
		obj, exists, err := c.Informer().GetStore().GetByKey(key)
		if err != nil {
			return err
		}
		if !exists {
			return handler(key, nil)
		}

		if !controller.ObjectInCluster(cluster, obj) {
			return nil
		}

		return handler(key, obj.(*ClusterRegistry))
	})
}

type clusterRegistryFactory struct {
}

func (c clusterRegistryFactory) Object() runtime.Object {
	return &ClusterRegistry{}
}

func (c clusterRegistryFactory) List() runtime.Object {
	return &ClusterRegistryList{}
}

func (s *clusterRegistryClient) Controller() ClusterRegistryController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.clusterRegistryControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ClusterRegistryGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &clusterRegistryController{
		GenericController: genericController,
	}

	s.client.clusterRegistryControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type clusterRegistryClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ClusterRegistryController
}

func (s *clusterRegistryClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *clusterRegistryClient) Create(o *ClusterRegistry) (*ClusterRegistry, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*ClusterRegistry), err
}

func (s *clusterRegistryClient) Get(name string, opts metav1.GetOptions) (*ClusterRegistry, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*ClusterRegistry), err
}

func (s *clusterRegistryClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*ClusterRegistry, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*ClusterRegistry), err
}

func (s *clusterRegistryClient) Update(o *ClusterRegistry) (*ClusterRegistry, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*ClusterRegistry), err
}

func (s *clusterRegistryClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *clusterRegistryClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *clusterRegistryClient) List(opts metav1.ListOptions) (*ClusterRegistryList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ClusterRegistryList), err
}

func (s *clusterRegistryClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *clusterRegistryClient) Patch(o *ClusterRegistry, data []byte, subresources ...string) (*ClusterRegistry, error) {
	obj, err := s.objectClient.Patch(o.Name, o, data, subresources...)
	return obj.(*ClusterRegistry), err
}

func (s *clusterRegistryClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *clusterRegistryClient) AddHandler(name string, sync ClusterRegistryHandlerFunc) {
	s.Controller().AddHandler(name, sync)
}

func (s *clusterRegistryClient) AddLifecycle(name string, lifecycle ClusterRegistryLifecycle) {
	sync := NewClusterRegistryLifecycleAdapter(name, false, s, lifecycle)
	s.AddHandler(name, sync)
}

func (s *clusterRegistryClient) AddClusterScopedHandler(name, clusterName string, sync ClusterRegistryHandlerFunc) {
	s.Controller().AddClusterScopedHandler(name, clusterName, sync)
}

func (s *clusterRegistryClient) AddClusterScopedLifecycle(name, clusterName string, lifecycle ClusterRegistryLifecycle) {
	sync := NewClusterRegistryLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.AddClusterScopedHandler(name, clusterName, sync)
}
