package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterAlertType                       = "clusterAlert"
	ClusterAlertFieldAnnotations           = "annotations"
	ClusterAlertFieldClusterId             = "clusterId"
	ClusterAlertFieldCreated               = "created"
	ClusterAlertFieldCreatorID             = "creatorId"
	ClusterAlertFieldDescription           = "description"
	ClusterAlertFieldDisplayName           = "displayName"
	ClusterAlertFieldInitialWaitSeconds    = "initialWaitSeconds"
	ClusterAlertFieldLabels                = "labels"
	ClusterAlertFieldName                  = "name"
	ClusterAlertFieldNamespaceId           = "namespaceId"
	ClusterAlertFieldOwnerReferences       = "ownerReferences"
	ClusterAlertFieldRecipients            = "recipients"
	ClusterAlertFieldRemoved               = "removed"
	ClusterAlertFieldRepeatIntervalSeconds = "repeatIntervalSeconds"
	ClusterAlertFieldSeverity              = "severity"
	ClusterAlertFieldState                 = "state"
	ClusterAlertFieldStatus                = "status"
	ClusterAlertFieldTargetEvent           = "targetEvent"
	ClusterAlertFieldTargetNode            = "targetNode"
	ClusterAlertFieldTargetSystemService   = "targetSystemService"
	ClusterAlertFieldTransitioning         = "transitioning"
	ClusterAlertFieldTransitioningMessage  = "transitioningMessage"
	ClusterAlertFieldUuid                  = "uuid"
)

type ClusterAlert struct {
	types.Resource
	Annotations           map[string]string    `json:"annotations,omitempty"`
	ClusterId             string               `json:"clusterId,omitempty"`
	Created               string               `json:"created,omitempty"`
	CreatorID             string               `json:"creatorId,omitempty"`
	Description           string               `json:"description,omitempty"`
	DisplayName           string               `json:"displayName,omitempty"`
	InitialWaitSeconds    *int64               `json:"initialWaitSeconds,omitempty"`
	Labels                map[string]string    `json:"labels,omitempty"`
	Name                  string               `json:"name,omitempty"`
	NamespaceId           string               `json:"namespaceId,omitempty"`
	OwnerReferences       []OwnerReference     `json:"ownerReferences,omitempty"`
	Recipients            []Recipient          `json:"recipients,omitempty"`
	Removed               string               `json:"removed,omitempty"`
	RepeatIntervalSeconds *int64               `json:"repeatIntervalSeconds,omitempty"`
	Severity              string               `json:"severity,omitempty"`
	State                 string               `json:"state,omitempty"`
	Status                *AlertStatus         `json:"status,omitempty"`
	TargetEvent           *TargetEvent         `json:"targetEvent,omitempty"`
	TargetNode            *TargetNode          `json:"targetNode,omitempty"`
	TargetSystemService   *TargetSystemService `json:"targetSystemService,omitempty"`
	Transitioning         string               `json:"transitioning,omitempty"`
	TransitioningMessage  string               `json:"transitioningMessage,omitempty"`
	Uuid                  string               `json:"uuid,omitempty"`
}
type ClusterAlertCollection struct {
	types.Collection
	Data   []ClusterAlert `json:"data,omitempty"`
	client *ClusterAlertClient
}

type ClusterAlertClient struct {
	apiClient *Client
}

type ClusterAlertOperations interface {
	List(opts *types.ListOpts) (*ClusterAlertCollection, error)
	Create(opts *ClusterAlert) (*ClusterAlert, error)
	Update(existing *ClusterAlert, updates interface{}) (*ClusterAlert, error)
	ByID(id string) (*ClusterAlert, error)
	Delete(container *ClusterAlert) error
}

func newClusterAlertClient(apiClient *Client) *ClusterAlertClient {
	return &ClusterAlertClient{
		apiClient: apiClient,
	}
}

func (c *ClusterAlertClient) Create(container *ClusterAlert) (*ClusterAlert, error) {
	resp := &ClusterAlert{}
	err := c.apiClient.Ops.DoCreate(ClusterAlertType, container, resp)
	return resp, err
}

func (c *ClusterAlertClient) Update(existing *ClusterAlert, updates interface{}) (*ClusterAlert, error) {
	resp := &ClusterAlert{}
	err := c.apiClient.Ops.DoUpdate(ClusterAlertType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterAlertClient) List(opts *types.ListOpts) (*ClusterAlertCollection, error) {
	resp := &ClusterAlertCollection{}
	err := c.apiClient.Ops.DoList(ClusterAlertType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterAlertCollection) Next() (*ClusterAlertCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterAlertCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterAlertClient) ByID(id string) (*ClusterAlert, error) {
	resp := &ClusterAlert{}
	err := c.apiClient.Ops.DoByID(ClusterAlertType, id, resp)
	return resp, err
}

func (c *ClusterAlertClient) Delete(container *ClusterAlert) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterAlertType, &container.Resource)
}
