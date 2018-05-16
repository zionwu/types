package v3

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterRegistry struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterRegistrySpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status RegistryStatus `json:"status"`
}

type ClusterRegistrySpec struct {
	DisplayName string `json:"displayName,omitempty" norman:"required"`

	ClusterName string `json:"clusterName" norman:"type=reference[cluster]"`
	Host        string `json:"host"`
}

type RegistryStatus struct {
	ABC string `json:"abc"`
}
