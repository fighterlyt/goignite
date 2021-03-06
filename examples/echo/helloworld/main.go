package main

import (
	"context"
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
	"github.com/b2wdigital/goignite/info"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const HelloWorldEndpoint = "app.endpoint.helloworld"

func init() {
	giconfig.Add(HelloWorldEndpoint, "/hello-world", "helloworld endpoint")
}

type Config struct {
	App struct {
		Endpoint struct {
			Helloworld string
		}
	}
}

type Response struct {
	Message string
}

func Get(c echo.Context) (err error) {

	l := gilog.FromContext(context.Background())

	resp := Response{
		Message: "Hello World!!",
	}

	err = giconfig.Unmarshal(&resp)
	if err != nil {
		l.Errorf(err.Error())
	}

	return giecho.JSON(c, http.StatusOK, resp, err)
}

func main() {

	giconfig.Load()

	c := Config{}

	err := giconfig.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	gilogrus.NewLogger()

	info.AppName = "helloworld"

	instance := giecho.Start(ctx)

	instance.Use(middleware.Gzip())
	instance.Use(middleware.CORS())
	instance.Use(middleware.RequestID())

	instance.GET(c.App.Endpoint.Helloworld, Get)

	giecho.Serve(ctx)
}
