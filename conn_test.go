package rabbitclient

import (
	"testing"
	"context"
	"net/http"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestConn_Get(t *testing.T) {
	c := NewConn("guest", "guest")
	ctx := context.TODO()
	err := c.get(ctx, "0.0.0.0:15672", "overview", func(c context.Context, resp *http.Response) error {
		decoder := json.NewDecoder(resp.Body)
		var overview Overview
		err := decoder.Decode(&overview)
		if err != nil {
			return err
		}
		assert.Equal(t, "rabbit@localhost", overview.StatisticsDBNode)
		return nil
	})
	assert.NoError(t, err)
}
