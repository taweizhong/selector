package wrr

import (
	"context"
	"github.com/taweizhong/selector"
	"sync"
)

type Balancer struct {
	mu            sync.Mutex
	currentWeight map[string]float64
}

func (b *Balancer) Pick(ctx context.Context, nodes []selector.Node) (selector.Node, error) {
	if len(nodes) == 0 {
		return nil, selector.ErrNoAvailableNode
	}
	b.mu.Lock()
	var totalWeight float64
	var currentNode selector.Node
	var nodeWeight float64
	for _, node := range nodes {
		totalWeight += node.Weight()
		curWeight := b.currentWeight[node.Address()]
		curWeight += node.Weight()
		b.currentWeight[node.Address()] = curWeight
		if currentNode == nil || nodeWeight < curWeight {
			currentNode = node
			nodeWeight = curWeight
		}
	}
	b.currentWeight[currentNode.Address()] = nodeWeight - totalWeight
	b.mu.Unlock()
	return currentNode, nil
}

func New() selector.Selector {
	return NewWRRBuilder().Builder()
}

func NewWRRBuilder() selector.Builder {
	return &selector.NodeSelectBuilder{
		BalancerBuilder: &Builder{},
	}
}

type Builder struct{}

func (b *Builder) Build() selector.Balancer {
	return &Balancer{
		currentWeight: make(map[string]float64),
	}
}
