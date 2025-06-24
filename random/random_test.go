package random

import (
	"context"
	"fmt"
	"github.com/taweizhong/selector"
	"testing"
)

func TestRandom(t *testing.T) {
	var nodes []selector.Node
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9090", 100))
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9091", 100))
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9092", 100))
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9093", 100))
	s := New()
	s.Appy(nodes)
	ctx := selector.BuildPeerContext(context.Background(), &selector.Peer{})
	node, err := s.Select(ctx)
	if err != nil {
		return
	}
	fmt.Println(node.Address())
	peer, _ := selector.GetPeerFromContext(ctx)
	fmt.Println(peer.Node.Scheme())
}
