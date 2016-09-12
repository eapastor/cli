package api

import "time"

type UserS struct {}

type userM struct {
	UUID     string    `json:"uuid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Gravatar string    `json:"gravatar"`
	Balance  float64   `json:"balance"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func (s UserS) Get(token string) (httpData HttpData, user *userM, err error) {

	failure := new(failureS)

	success := new(userM)

	httpData.Request, httpData.Response, err = httpClient.GET("/user").AddHeader("Authorization", token).Request(&success, &failure)
	if err != nil {
		return httpData, nil, err
	}

	if ok, err := checkError(failure); ok {
		return httpData, nil, err
	}

	return httpData, success, nil

}
