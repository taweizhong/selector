package selector

import "context"

type Balancer interface {
	Pick(ctx context.Context, nodes []Node) (Node, error)
}

type BalancerBuilder interface {
	Build() Balancer
}
