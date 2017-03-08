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

func (c *Conn) GetVhosts(ctx context.Context, host string, outC chan<- []Vhost, errC chan<- error) {
	err := c.get(ctx, host, "vhosts", func(c context.Context, resp *http.Response) error {
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

func (c *Conn) GetVhost(ctx context.Context, host, vhostName string, outC chan<- Vhost, errC chan<- error) {
	if vhostName == "/" {
		vhostName = "%2f"
	}
	err := c.get(ctx, host, "vhosts/"+vhostName, func(c context.Context, resp *http.Response) error {
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
