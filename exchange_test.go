package rabbitclient

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
)

func TestGetExchanges(t *testing.T) {
	c := newConn("guest", "guest", "0.0.0.0:15672")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan []Exchange)
	go GetExchanges(ctx, c, outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case exchanges := <-outC:
		assert.Equal(t, "/", exchanges[0].Vhost)
	}
}

func TestGetVhostExchanges(t *testing.T) {
	c := newConn("guest", "guest", "0.0.0.0:15672")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan []Exchange)
	go GetVhostExchanges(ctx, c, "/", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case exchanges := <-outC:
		assert.Equal(t, "/", exchanges[0].Vhost)
	}
}

func TestGetExchange(t *testing.T) {
	c := newConn("guest", "guest", "0.0.0.0:15672")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan Exchange)
	go GetExchange(ctx, c, "/", "amq.rabbitmq.trace", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case exchange := <-outC:
		assert.NotNil(t, exchange)
	}
}
