package selector

import "context"

type Filter func(ctx context.Context, nodes []Node) []Node
