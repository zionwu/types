package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterRegistryLifecycle interface {
	Create(obj *ClusterRegistry) (*ClusterRegistry, error)
	Remove(obj *ClusterRegistry) (*ClusterRegistry, error)
	Updated(obj *ClusterRegistry) (*ClusterRegistry, error)
}

type clusterRegistryLifecycleAdapter struct {
	lifecycle ClusterRegistryLifecycle
}

func (w *clusterRegistryLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*ClusterRegistry))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterRegistryLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*ClusterRegistry))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterRegistryLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*ClusterRegistry))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewClusterRegistryLifecycleAdapter(name string, clusterScoped bool, client ClusterRegistryInterface, l ClusterRegistryLifecycle) ClusterRegistryHandlerFunc {
	adapter := &clusterRegistryLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *ClusterRegistry) error {
		if obj == nil {
			return syncFn(key, nil)
		}
		return syncFn(key, obj)
	}
}
