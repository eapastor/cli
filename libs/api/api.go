package api

import (
	hlb "github.com/lastbackend/cli/libs/api/http"
	"net/http"
)

var httpClient *hlb.RawReq

type API struct {
	Session sessionS
}

type HttpData struct {
	Request  *http.Request
	Response *http.Response
}

func New(host string) API {

	httpClient = hlb.New(host)
	return API{
		Session: sessionS{},
	}

}