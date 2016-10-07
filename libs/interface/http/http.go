package http

import (
	LBhttp "github.com/lastbackend/lb/libs/http"
	"net/http"
)

type IHTTP interface {
	Request(successV, failureV interface{}) (eq *http.Request, resp *http.Response, err error)
	// Methods
	POST(pathURL string) *LBhttp.RawReq
	GET(pathURL string) *LBhttp.RawReq
	PUT(pathURL string) *LBhttp.RawReq
	DELETE(pathURL string) *LBhttp.RawReq
	// Header
	AddHeader(key, value string) *LBhttp.RawReq
}
