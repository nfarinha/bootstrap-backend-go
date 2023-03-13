package webserver

import "github.com/labstack/echo/v4"

type WebServer struct {
	*echo.Echo
}

type HandlerFunc func(c *Context) error
type InitializerFunc func(r *WebServer) error

func New(initializers *[]InitializerFunc) *WebServer {
	e := echo.New()
	r := &WebServer{e}

	// Hide banner
	e.HideBanner = true

	// wrapper to extend echo Context with additional helpers
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := newContext(c)
			return next(cc)
		}
	})

	// after all internal initializations, invoke any external ones
	for _, f := range *initializers {
		f(r)
	}
	return r
}

func (r *WebServer) DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *WebServer {
	r.Echo.DELETE(path, func(c echo.Context) error { return h(c.(*Context)) }, m...)
	return r
}

func (r *WebServer) HEAD(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *WebServer {
	r.Echo.HEAD(path, func(c echo.Context) error { return h(c.(*Context)) }, m...)
	return r
}

func (r *WebServer) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *WebServer {
	r.Echo.GET(path, func(c echo.Context) error { return h(c.(*Context)) }, m...)
	return r
}

func (r *WebServer) PATCH(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *WebServer {
	r.Echo.PATCH(path, func(c echo.Context) error { return h(c.(*Context)) }, m...)
	return r
}

func (r *WebServer) POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *WebServer {
	r.Echo.POST(path, func(c echo.Context) error { return h(c.(*Context)) }, m...)
	return r
}

func (r *WebServer) PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *WebServer {
	r.Echo.PUT(path, func(c echo.Context) error { return h(c.(*Context)) }, m...)
	return r
}
