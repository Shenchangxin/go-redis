package tcp

import (
	"context"
	"net"
)

type Handler interface {
	Handle(ctx context.Context, coon net.Conn)
	Close() error
}
