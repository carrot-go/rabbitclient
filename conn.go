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
	host     string
}

func NewConn(user, pwd, host string) *Conn {
	c := &Conn{user: user, password: pwd, host: host}
	c.hc = &http.Client{}
	return c
}

func (c *Conn) newRequest(method, endpoint string) (*http.Request, error) {
	requestUrl := "http://" + c.host + "/api/" + endpoint
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

func (c *Conn) get(ctx context.Context, endpoint string, ret_func func(c context.Context, resp *http.Response) error) error {
	return c.do(ctx, "GET", endpoint, ret_func)
}

func (c *Conn) do(ctx context.Context, method, endpoint string, ret_func func(c context.Context, resp *http.Response) error) error {
	req, err := c.newRequest(method, endpoint)
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
