package api

type serviceS struct{}

type serviceM struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Status    string `json:"status"`
	Provision bool   `json:"provision,omitempty"`
	Source    struct {
		Type       string `json:"type,omitempty"`
		Hub        string `json:"hub,omitempty"` //github,bitbucket,gitlab,dockerhub,registry
		Owner      string `json:"owner,omitempty"`
		Repository string `json:"repo,omitempty"`
		Branch     string `json:"branch,omitempty"`
	} `json:"source"`
}

type servicesM []serviceM

func (s serviceS) List(token string) (HttpData, *[]serviceM, error) {

	var httpData HttpData
	var err error

	failure := new(failureS)

	success := new([]serviceM)

	httpData.Request, httpData.Response, err = httpClient.GET("/services").AddHeader("Authorization", token).Request(&success, &failure)
	if err != nil {
		return httpData, nil, err
	}

	if ok, err := checkError(failure); ok {
		return httpData, nil, err
	}

	return httpData, success, nil

}
