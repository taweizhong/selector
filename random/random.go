package random

import (
	"context"
	"github.com/taweizhong/selector"
	"math/rand"
)

type BalancerBuilder interface {
	Builder()
}
type Balancer struct{}

func (b *Balancer) Pick(ctx context.Context, nodes []selector.Node) (selector.Node, error) {
	if len(nodes) == 0 {
		return nil, selector.ErrNoAvailableNode
	}
	index := rand.Intn(len(nodes))
	return nodes[index], nil
}

func New() selector.Selector {
	return NewRandomSelector().Builder()
}

func NewRandomSelector() selector.Builder {
	return &selector.NodeSelectBuilder{
		BalancerBuilder: &Builder{},
	}
}

type Builder struct{}

func (b *Builder) Build() selector.Balancer {
	return &Balancer{}
}
