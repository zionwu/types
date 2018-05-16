package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterRegistryType                      = "clusterRegistry"
	ClusterRegistryFieldAnnotations          = "annotations"
	ClusterRegistryFieldClusterId            = "clusterId"
	ClusterRegistryFieldCreated              = "created"
	ClusterRegistryFieldCreatorID            = "creatorId"
	ClusterRegistryFieldHost                 = "host"
	ClusterRegistryFieldLabels               = "labels"
	ClusterRegistryFieldName                 = "name"
	ClusterRegistryFieldNamespaceId          = "namespaceId"
	ClusterRegistryFieldOwnerReferences      = "ownerReferences"
	ClusterRegistryFieldRemoved              = "removed"
	ClusterRegistryFieldState                = "state"
	ClusterRegistryFieldStatus               = "status"
	ClusterRegistryFieldTransitioning        = "transitioning"
	ClusterRegistryFieldTransitioningMessage = "transitioningMessage"
	ClusterRegistryFieldUuid                 = "uuid"
)

type ClusterRegistry struct {
	types.Resource
	Annotations          map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	ClusterId            string            `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Created              string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Host                 string            `json:"host,omitempty" yaml:"host,omitempty"`
	Labels               map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                 string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed              string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                string            `json:"state,omitempty" yaml:"state,omitempty"`
	Status               *RegistryStatus   `json:"status,omitempty" yaml:"status,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	Uuid                 string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}
type ClusterRegistryCollection struct {
	types.Collection
	Data   []ClusterRegistry `json:"data,omitempty"`
	client *ClusterRegistryClient
}

type ClusterRegistryClient struct {
	apiClient *Client
}

type ClusterRegistryOperations interface {
	List(opts *types.ListOpts) (*ClusterRegistryCollection, error)
	Create(opts *ClusterRegistry) (*ClusterRegistry, error)
	Update(existing *ClusterRegistry, updates interface{}) (*ClusterRegistry, error)
	ByID(id string) (*ClusterRegistry, error)
	Delete(container *ClusterRegistry) error
}

func newClusterRegistryClient(apiClient *Client) *ClusterRegistryClient {
	return &ClusterRegistryClient{
		apiClient: apiClient,
	}
}

func (c *ClusterRegistryClient) Create(container *ClusterRegistry) (*ClusterRegistry, error) {
	resp := &ClusterRegistry{}
	err := c.apiClient.Ops.DoCreate(ClusterRegistryType, container, resp)
	return resp, err
}

func (c *ClusterRegistryClient) Update(existing *ClusterRegistry, updates interface{}) (*ClusterRegistry, error) {
	resp := &ClusterRegistry{}
	err := c.apiClient.Ops.DoUpdate(ClusterRegistryType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterRegistryClient) List(opts *types.ListOpts) (*ClusterRegistryCollection, error) {
	resp := &ClusterRegistryCollection{}
	err := c.apiClient.Ops.DoList(ClusterRegistryType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterRegistryCollection) Next() (*ClusterRegistryCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterRegistryCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterRegistryClient) ByID(id string) (*ClusterRegistry, error) {
	resp := &ClusterRegistry{}
	err := c.apiClient.Ops.DoByID(ClusterRegistryType, id, resp)
	return resp, err
}

func (c *ClusterRegistryClient) Delete(container *ClusterRegistry) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterRegistryType, &container.Resource)
}
