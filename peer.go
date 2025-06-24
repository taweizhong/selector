package selector

import "context"

type peerKey struct{}

type Peer struct {
	Node Node
}

func BuildPeerContext(ctx context.Context, p *Peer) context.Context {
	ctx = context.WithValue(ctx, peerKey{}, p)
	return ctx
}
func GetPeerFromContext(ctx context.Context) (*Peer, bool) {
	peer, ok := ctx.Value(peerKey{}).(*Peer)
	return peer, ok
}
