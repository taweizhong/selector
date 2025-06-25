package selector

type DefaultNode struct {
	scheme   string
	addr     string
	weight   float64
	version  string
	name     string
	metadata map[string]string
}

// Scheme is node scheme
func (n *DefaultNode) Scheme() string {
	return n.scheme
}

// Address is node address
func (n *DefaultNode) Address() string {
	return n.addr
}

// ServiceName is node serviceName
func (n *DefaultNode) ServiceName() string {
	return n.name
}

// Version is node version
func (n *DefaultNode) Version() string {
	return n.version
}

// Metadata is node metadata
func (n *DefaultNode) Metadata() map[string]string {
	return n.metadata
}
func (n *DefaultNode) Weight() float64 {
	return n.weight
}
func NewDefaultNode(scheme string, addr string, weight float64) *DefaultNode {
	return &DefaultNode{
		scheme:   scheme,
		addr:     addr,
		weight:   weight,
		version:  "0.0.1",
		metadata: make(map[string]string),
	}
}
