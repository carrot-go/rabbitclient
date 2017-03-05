package rabbitclient

import (
	"context"
	"net/http"
	"encoding/json"
)

type Vhost struct {
	Name    string
	Tracing bool
}

func GetVhosts(ctx context.Context, c *conn, outC chan<- []Vhost, errC chan<- error) {
	err := c.get(ctx, "vhosts", func(c context.Context, resp *http.Response) error {
		var vhost []Vhost
		err := json.NewDecoder(resp.Body).Decode(&vhost)
		if err != nil {
			return err
		}
		outC <- vhost
		return nil
	})
	if err != nil {
		errC <- err
	}
}

func GetVhost(ctx context.Context, c *conn, vhostName string, outC chan<- Vhost, errC chan<- error) {
	if vhostName == "/" {
		vhostName = "%2f"
	}
	err := c.get(ctx, "vhosts/"+vhostName, func(c context.Context, resp *http.Response) error {
		var vhost Vhost
		err := json.NewDecoder(resp.Body).Decode(&vhost)
		if err != nil {
			return err
		}
		outC <- vhost
		return nil
	})
	if err != nil {
		errC <- err
	}
}
