package rabbitclient

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
)

func TestConn_GetExchanges(t *testing.T) {
	c := NewConn("guest", "guest", "0.0.0.0:15672")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan []Exchange)
	go c.GetExchanges(ctx, outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case exchanges := <-outC:
		assert.Equal(t, "/", exchanges[0].Vhost)
	}
}

func TestConn_GetVhostExchanges(t *testing.T) {
	c := NewConn("guest", "guest", "0.0.0.0:15672")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan []Exchange)
	go c.GetVhostExchanges(ctx, "/", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case exchanges := <-outC:
		assert.Equal(t, "/", exchanges[0].Vhost)
	}
}

func TestConn_GetExchange(t *testing.T) {
	c := NewConn("guest", "guest", "0.0.0.0:15672")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan Exchange)
	go c.GetExchange(ctx,"/", "amq.rabbitmq.trace", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case exchange := <-outC:
		assert.NotNil(t, exchange)
	}
}
