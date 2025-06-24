package selector

import (
	"context"
	"errors"
	"sync/atomic"
)

var ErrNoAvailableNode = errors.New("no available node")

var globalSelector Selector

func GetGlobalSelector() Selector {
	return globalSelector
}
func SetGlobalSelector(s Selector) {
	globalSelector = s
}

type Config struct {
	NodesFilter []Filter
}
type Option func(*Config)

func WithNodesFilter(filters ...Filter) Option {
	return func(config *Config) {
		config.NodesFilter = filters
	}
}

type Selector interface {
	Rebalancer
	Select(ctx context.Context, opts ...Option) (node Node, err error)
}

type Rebalancer interface {
	Appy(nodes []Node)
}

type NodeSelector struct {
	Balancer Balancer
	nodes    atomic.Value
}

func (s *NodeSelector) Select(ctx context.Context, opts ...Option) (node Node, err error) {
	nodes, ok := s.nodes.Load().([]Node)
	newNodes := make([]Node, len(nodes))
	for i, node := range nodes {
		newNodes[i] = node
	}
	if !ok {
		return nil, ErrNoAvailableNode
	}
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}
	if len(config.NodesFilter) > 0 {
		for _, filter := range config.NodesFilter {
			newNodes = filter(ctx, newNodes)
		}
	}
	node, err = s.Balancer.Pick(ctx, newNodes)
	if err != nil {
		return nil, err
	}
	peer, ok := GetPeerFromContext(ctx)
	if ok {
		peer.Node = node
	}
	return node, nil
}

func (s *NodeSelector) Appy(nodes []Node) {
	s.nodes.Store(nodes)
}

type NodeSelectBuilder struct {
	BalancerBuilder BalancerBuilder
}

func (b *NodeSelectBuilder) Builder() Selector {
	return &NodeSelector{
		Balancer: b.BalancerBuilder.Build(),
	}
}

type Node interface {
	Scheme() string
	Address() string
	ServiceName() string
	InitialWeight() *int64
	Version() string
	Metadata() map[string]string
}

type Builder interface {
	Builder() Selector
}
