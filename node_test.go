package rabbitclient

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
)

func TestConn_GetNodes(t *testing.T) {
	c := NewConn("guest", "guest")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan []Node)
	go c.GetNodes(ctx, "0.0.0.0:15672", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case o := <-outC:
		assert.Equal(t, 1, len(o))
	}
}

func TestConn_GetNode(t *testing.T) {
	c := NewConn("guest", "guest")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan Node)
	go c.GetNode(ctx, "0.0.0.0:15672", "rabbit@localhost", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case o := <-outC:
		assert.NotNil(t, o)
	}
}
