package client

const (
	ZookeeperType          = "zookeeper"
	ZookeeperFieldEndpoint = "endpoint"
)

type Zookeeper struct {
	Endpoint string `json:"endpoint,omitempty"`
}
