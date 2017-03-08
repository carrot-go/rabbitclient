package rabbitclient

import (
	"net/http"
	"context"
	"fmt"
	"net/url"
)

type Conn struct {
	hc       *http.Client
	user     string
	password string
}

func NewConn(user, pwd string) *Conn {
	c := &Conn{user: user, password: pwd}
	c.hc = &http.Client{}
	return c
}

func (c *Conn) newRequest(host, method, endpoint string) (*http.Request, error) {
	requestUrl := "http://" + host + "/api/" + endpoint
	u, err := url.Parse(requestUrl)
	if err != nil {
		return nil, err
	}

	u.Opaque = "/api/" + endpoint

	req := &http.Request{
		Method:     method,
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Body:       nil,
		Host:       u.Host,
	}

	req.SetBasicAuth(c.user, c.password)
	req.Header.Set("Accept", "application/json");
	req.Header.Set("User-Agent", "lysu/rabbitclient")

	return req, nil
}

func (c *Conn) get(ctx context.Context, host string, endpoint string, ret_func func(c context.Context, resp *http.Response) error) error {
	return c.do(ctx, host, "GET", endpoint, ret_func)
}

func (c *Conn) do(ctx context.Context, host, method, endpoint string, ret_func func(c context.Context, resp *http.Response) error) error {
	req, err := c.newRequest(host, method, endpoint)
	if err != nil {
		return err
	}
	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf(resp.Status)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	return ret_func(ctx, resp)
}

func (c *Conn) Destroy() {
}