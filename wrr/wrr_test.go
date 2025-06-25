package wrr

import (
	"context"
	"fmt"
	"github.com/taweizhong/selector"
	"testing"
)

func Test_WRR(t *testing.T) {
	var nodesLog map[string]int
	nodesLog = make(map[string]int)
	var nodes []selector.Node
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9090", 6))
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9091", 3))
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9092", 1))
	wrr := New()
	wrr.Appy(nodes)
	for i := 0; i < 100; i++ {
		node, _ := wrr.Select(context.Background())
		nodesLog[node.Address()]++
	}
	for k, v := range nodesLog {
		fmt.Println(k, v)
	}
}
