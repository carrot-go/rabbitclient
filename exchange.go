package rabbitclient

import (
	"context"
	"net/http"
	"encoding/json"
)

type Exchange struct {
	Arguments  map[string]interface{} `json:"arguments"`
	AutoDelete bool                   `json:"auto_delete"`
	Durable    bool                   `json:"durable"`
	Internal   bool                   `json:"durable"`
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	Vhost      string                 `json:"vhost"`
}

func (c *Conn) GetExchanges(ctx context.Context, host string, outC chan<- []Exchange, errC chan<- error) {
	err := c.get(ctx, host, "exchanges", func(c context.Context, resp *http.Response) error {
		var exchanges []Exchange
		err := json.NewDecoder(resp.Body).Decode(&exchanges)
		if err != nil {
			return err
		}
		outC <- exchanges
		return nil
	})
	if err != nil {
		errC <- err
	}
}

func (c *Conn) GetVhostExchanges(ctx context.Context, host, vhost string, outC chan<- []Exchange, errC chan<- error) {
	if vhost == "/" {
		vhost = "%2f"
	}
	err := c.get(ctx, host, "exchanges/"+vhost, func(c context.Context, resp *http.Response) error {
		var exchanges []Exchange
		err := json.NewDecoder(resp.Body).Decode(&exchanges)
		if err != nil {
			return err
		}
		outC <- exchanges
		return nil
	})
	if err != nil {
		errC <- err
	}
}

func (c *Conn) GetExchange(ctx context.Context, host, vhost, name string, outC chan<- Exchange, errC chan<- error) {
	if vhost == "/" {
		vhost = "%2f"
	}
	err := c.get(ctx, host, "exchanges/"+vhost+"/"+name, func(c context.Context, resp *http.Response) error {
		var exchange Exchange
		err := json.NewDecoder(resp.Body).Decode(&exchange)
		if err != nil {
			return err
		}
		outC <- exchange
		return nil
	})
	if err != nil {
		errC <- err
	}
}
