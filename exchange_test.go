package rabbitclient

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
)

func TestConn_GetExchanges(t *testing.T) {
	c := NewConn("guest", "guest")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan []Exchange)
	go c.GetExchanges(ctx, "0.0.0.0:15672",outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case exchanges := <-outC:
		assert.Equal(t, "/", exchanges[0].Vhost)
	}
}

func TestConn_GetVhostExchanges(t *testing.T) {
	c := NewConn("guest", "guest")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan []Exchange)
	go c.GetVhostExchanges(ctx, "0.0.0.0:15672", "/", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case exchanges := <-outC:
		assert.Equal(t, "/", exchanges[0].Vhost)
	}
}

func TestConn_GetExchange(t *testing.T) {
	c := NewConn("guest", "guest")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan Exchange)
	go c.GetExchange(ctx, "0.0.0.0:15672", "/", "amq.rabbitmq.trace", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case exchange := <-outC:
		assert.NotNil(t, exchange)
	}
}
