package goweb_test

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	"goweb"
)

func TestGET(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.GET("/", handler)
	assert(t, app, "GET", "/", nil, nil, http.StatusOK, "")
}

func TestPUT(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.PUT("/", handler)
	assert(t, app, "PUT", "/", nil, nil, http.StatusOK, "")
}

func TestPATCH(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.PATCH("/", handler)
	assert(t, app, "PATCH", "/", nil, nil, http.StatusOK, "")
}

func TestPOST(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.POST("/", handler)
	assert(t, app, "POST", "/", nil, nil, http.StatusOK, "")
}

func TestDELETE(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.DELETE("/", handler)
	assert(t, app, "DELETE", "/", nil, nil, http.StatusOK, "")
}

func TestOPTIONS(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.OPTIONS("/", handler)
	assert(t, app, "OPTIONS", "/", nil, nil, http.StatusOK, "")
}

func TestHEAD(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.HEAD("/", handler)
	assert(t, app, "HEAD", "/", nil, nil, http.StatusOK, "")
}

func TestRouteNotFound(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.GET("/", handler)
	assert(t, app, "GET", "/foo", nil, nil, http.StatusNotFound, "Page not found")
}

func TestCustomNotFound(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.GET("/", handler)
	app.NotFound(func(c *goweb.Context) *goweb.Response {
		return c.NotFound()
	})
	assert(t, app, "GET", "/foo", nil, nil, http.StatusNotFound, "")
}

func TestEmptyPath(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	assertPanic(t, func() {
		app.GET("", handler)
	})
}

func TestNonSlashLeadingPath(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	assertPanic(t, func() {
		app.GET("foo", handler)
	})
}

func TestParams(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		res := fmt.Sprintf("%s %s", c.Param("name"), c.Param("age"))
		return c.OK().PlainText(res)
	}
	app := goweb.New()
	app.GET("/hello/{name}/{age}", handler)
	assert(t, app, "GET", "/hello/Gopher/5", nil, nil, http.StatusOK, "Gopher 5")
}

func TestParamsNotFound(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		res := fmt.Sprintf("%s %s", c.Param("name_"), c.Param("age"))
		return c.OK().PlainText(res)
	}
	app := goweb.New()
	app.GET("/hello/{name}/{age}", handler)
	assert(t, app, "GET", "/hello/Gopher/5", nil, nil, http.StatusOK, " 5")
}

func TestNoWWW(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.NoWWW()
	app.GET("/", handler)
	transformer := func(r *http.Request) {
		r.Host = "www.example.com"
	}
	assert(t, app, "GET", "/", nil, transformer, http.StatusMovedPermanently, "<a href=\"http://example.com/\">Moved Permanently</a>.\n\n")
}

func TestNoWWWTLS(t *testing.T) {
	handler := func(c *goweb.Context) *goweb.Response {
		return c.OK()
	}
	app := goweb.New()
	app.NoWWW()
	app.GET("/", handler)
	transformer := func(r *http.Request) {
		r.Host = "www.example.com"
		r.TLS = &tls.ConnectionState{}
	}
	assert(t, app, "GET", "/", nil, transformer, http.StatusMovedPermanently, "<a href=\"https://example.com/\">Moved Permanently</a>.\n\n")
}
