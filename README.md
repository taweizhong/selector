## selecter（节点选择器）

### 简介

使用背景：用于LLM服务中相同名称服务、不同资源配置的服务选择。

为多个LLM服务实现负载均衡，可以根据不同的负载均衡器选择最优的节点，用户请求将转发到该节点。

### 使用

#### 随机选择

随机选择节点并返回。

```
package test

import (
	"context"
	"fmt"
	"github.com/taweizhong/selector"
	"github.com/taweizhong/selector/random"
	"testing"
)

func Test_Select(t *testing.T) {
	var nodes []selector.Node
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9090", 100))
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9091", 100))
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9092", 100))
	nodes = append(nodes, selector.NewDefaultNode("http", "127.0.0.1:9093", 100))
	s := random.New()
	s.Appy(nodes)
	node, err := s.Select(context.Background())
	if err != nil {
		return
	}
	fmt.Println(node.Address())
}
```

随机返回并存储与context

```
ctx := selector.BuildPeerContext(context.Background(), &selector.Peer{})
node, err := s.Select(ctx)
	if err != nil {
		return
	}
	fmt.Println(node.Address())
	peer, _ := selector.GetPeerFromContext(ctx)
	fmt.Println(peer.Node.Scheme())
```

Select函数会将选择之后的节点存储到ctx中，方便传递到下一步骤。

#### 加权轮训

根据节点的权重返回节点。

```
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
```

