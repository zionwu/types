package client

const (
	ClusterRegistrySpecType             = "clusterRegistrySpec"
	ClusterRegistrySpecFieldClusterId   = "clusterId"
	ClusterRegistrySpecFieldDisplayName = "displayName"
	ClusterRegistrySpecFieldHost        = "host"
)

type ClusterRegistrySpec struct {
	ClusterId   string `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Host        string `json:"host,omitempty" yaml:"host,omitempty"`
}
