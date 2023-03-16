package common

import (
	"net/url"
)

type ClientRequest struct {
	Url     url.URL
	Headers []ClientRequestHeader
	// I need union types :((
	Method string
	Body   string
}

type ClientRequestHeader struct {
	Name  string
	Value string
}
