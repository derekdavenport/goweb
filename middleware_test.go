package goweb_test

import (
	"net/http"
	"testing"

	"github.com/twharmon/goweb"
)

func TestPassThroughMiddleware(t *testing.T) {
	handler := func(c *goweb.Context) goweb.Responder {
		return c.OK().Text(c.Get("foo").(string))
	}
	app := goweb.New()
	mw := goweb.NewMiddleware()
	mw.Use(func(c *goweb.Context) goweb.Responder {
		c.Set("foo", "bar")
		return nil
	})
	app.GET("/", mw.Apply(handler))
	assert(t, app, "GET", "/", nil, nil, http.StatusOK, "bar")
}

func TestInterruptingMiddleware(t *testing.T) {
	handler := func(c *goweb.Context) goweb.Responder {
		return c.OK()
	}
	app := goweb.New()
	mw := goweb.NewMiddleware()
	mw.Use(func(c *goweb.Context) goweb.Responder {
		return c.BadRequest()
	})
	app.GET("/", mw.Apply(handler))
	assert(t, app, "GET", "/", nil, nil, http.StatusBadRequest, "")
}
