package rabbitclient

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
)

func TestConn_GetQueues(t *testing.T) {
	c := NewConn("guest", "guest")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan []Queue)
	go c.GetQueues(ctx, "0.0.0.0:15672", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case q := <-outC:
		assert.Equal(t, 1, len(q))
	}
}

func TestConn_GetVhostQueue(t *testing.T) {
	c := NewConn("guest", "guest")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan []Queue)
	go c.GetVhostQueue(ctx, "0.0.0.0:15672", "/", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case q := <-outC:
		assert.Equal(t, 1, len(q))
	}
}

func TestConn_GetQueue(t *testing.T) {
	c := NewConn("guest", "guest")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan Queue)
	go c.GetQueue(ctx, "0.0.0.0:15672","/", "testqueue", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case q := <-outC:
		assert.NotNil(t, q)
	}
}

