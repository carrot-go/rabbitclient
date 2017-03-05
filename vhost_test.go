package rabbitclient

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
)

func TestConn_GetVhosts(t *testing.T) {
	c := newConn("guest", "guest", "0.0.0.0:15672")
	ctx := context.TODO()
	errC := make(chan error, 1)
	outC := make(chan []Vhost, 1)
	c.GetVhosts(ctx, outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case item := <-outC:
		assert.Equal(t, "/", item[0].Name)
	}
}

func TestConn_GetVhost(t *testing.T) {
	c := newConn("guest", "guest", "0.0.0.0:15672")
	ctx := context.TODO()
	errC := make(chan error, 1)
	outC := make(chan Vhost, 1)
	c.GetVhost(ctx, "/", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case vhost := <-outC:
		assert.Equal(t, "/", vhost.Name)
	}
}