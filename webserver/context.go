package webserver

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Context struct {
	echo.Context
	Helpers Helpers
}

type Helpers struct {
	e      echo.Context
	exreq  Request
	exresp Response
}

type Request struct {
	ectx *echo.Context
}

type Response struct {
	ectx *echo.Context
}

type Param struct {
	name  string
	value string
}

func newContext(ctx echo.Context) *Context {
	return &Context{ctx, Helpers{ctx, Request{&ctx}, Response{&ctx}}}
}

func (c Helpers) Request() *Request {
	return &c.exreq
}

func (c Helpers) Response() *Response {
	return &c.exresp
}

func (r Request) Body(data any) error {
	return (&echo.DefaultBinder{}).BindBody(*r.ectx, data)
}

func (r Request) Query(data any) error {
	return (&echo.DefaultBinder{}).BindQueryParams(*r.ectx, data)
}

func (r Request) Param(name string) *Param {
	return &Param{name: name, value: (*r.ectx).Param(name)}
}

func (r Response) Error(httpStatusCode int, errorCode string, message string, beans ...any) error {
	return (*r.ectx).JSON(httpStatusCode, ResponseError{Code: errorCode, Message: fmt.Sprintf(message, beans...)})
}

func (r Response) String(httpStatusCode int, body string) error {
	return (*r.ectx).String(httpStatusCode, body)
}

func (r Response) JSON(httpStatusCode int, body any) error {
	return (*r.ectx).JSON(httpStatusCode, body)
}

func (p Param) Int() (v int, e error) {
	v, e = strconv.Atoi(p.value)
	if e != nil {
		e = fmt.Errorf(`invalid value of parameter "%s"; expected an integer, received "%s"`, p.name, p.value)
	}
	return v, e
}

func (p Param) String() (v string, e error) {
	return p.value, nil
}
