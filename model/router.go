package model

import "net/http"

type Router struct {
	Method    string
	Path      string
	Name      string
	Handler   http.HandlerFunc
	Protected bool
}
