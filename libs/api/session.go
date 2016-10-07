package api

type sessionS struct{}

func (s sessionS) Get(login string, pass string) (httpData HttpData, token string, err error) {

	failure := new(failureS)

	credentials := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{
		Login:    login,
		Password: pass,
	}

	success := struct {
		Token string `json:"token"`
	}{}

	success.Token = success.Token

	httpData.Request, httpData.Response, err = httpClient.POST("/session").
	                                           BodyJSON(credentials).
		                                         Request(&success, &failure)
	if err != nil {
		return httpData, "", err
	}

	if ok, err := checkError(failure); ok {
		return httpData, "", err
	}

	return httpData, success.Token, nil

}
