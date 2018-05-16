package client

const (
	RegistryStatusType     = "registryStatus"
	RegistryStatusFieldABC = "abc"
)

type RegistryStatus struct {
	ABC string `json:"abc,omitempty" yaml:"abc,omitempty"`
}
