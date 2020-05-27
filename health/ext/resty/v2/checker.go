package gihealthresty

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Checker struct {
	client  *resty.Client
	options *Options
}

func (c *Checker) Check(ctx context.Context) (err error) {

	request := c.client.R().EnableTrace()

	var response *resty.Response

	response, err = request.Get(strings.Join([]string{c.options.Host, c.options.Endpoint}, ""))

	if response.IsError() {
		return errors.New(strconv.Itoa(response.StatusCode()))
	}

	return err
}

func NewChecker(client *resty.Client, options *Options) *Checker {
	return &Checker{client: client, options: options}
}