package rabbitclient

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
)

func TestGetVhosts(t *testing.T) {
	c := newConn("guest", "guest", "0.0.0.0:15672")
	ctx := context.TODO()
	errC := make(chan error, 1)
	outC := make(chan []Vhost, 1)
	GetVhosts(ctx, c, outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case item := <-outC:
		assert.Equal(t, "/", item[0].Name)
	}
}

func TestGetVhost(t *testing.T) {
	c := newConn("guest", "guest", "0.0.0.0:15672")
	ctx := context.TODO()
	errC := make(chan error, 1)
	outC := make(chan Vhost, 1)
	GetVhost(ctx, c, "/", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case vhost := <-outC:
		assert.Equal(t, "/", vhost.Name)
	}
}