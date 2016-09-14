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

type serviceCreateS struct {
	Opts optsS`json:"opts,omitempty"`
	Node nodeS `json:"node,omitempty"`
	Source sourceCreateS `json:"source,omitempty"`
}

type optsS struct {
	Name   string `json:"name,omitempty"`
	Memory uint   `json:"memory,omitempty"`
}

type nodeS struct {
	UUID   string `json:"uuid,omitempty"`
	Type   string `json:"type,omitempty"`
	Region string `json:"region,omitempty"`
}

type sourceCreateS struct {
	ID        string `json:"id,omitempty"`
	Type      string `json:"type,omitempty"`
	Hub       string `json:"hub,omitempty"`
	Owner     string `json:"owner,omitempty"`
	Repo      string `json:"repo,omitempty"`
	Branch    string `json:"branch,omitempty"`
	Autobuild bool   `json:"autobuild,omitempty"`
}

func (serviceS) Create(token, name string, memory uint, nodeUUID, nodeType, nodeRegion, sourceID, sourceType,
	sourceHub, sourceOwner, sourceRepo, sourceBranch string, sourceAutobuild bool) (HttpData, *serviceM, error) {

	var httpData HttpData
	var err error

	var sc serviceCreateS = serviceCreateS{
		Opts: optsS{
			Name:   name,
			Memory: memory,
		},
		Node: nodeS{
			UUID:   nodeUUID,
			Type:   nodeType,
			Region: nodeRegion,
		},
		Source: sourceCreateS{
			ID:        sourceID,
			Type:      sourceType,
			Hub:       sourceHub,
			Owner:     sourceOwner,
			Repo:      sourceRepo,
			Branch:    sourceBranch,
			Autobuild: sourceAutobuild,
		},
	}

	failure := new(failureS)

	success := new(serviceM)

	httpData.Request, httpData.Response, err = httpClient.PUT("/service").BodyJSON(sc).AddHeader("Authorization", token).Request(&success, &failure)
	if err != nil {
		return httpData, nil, err
	}

	if ok, err := checkError(failure); ok {
		return httpData, nil, err
	}

	return httpData, success, nil

}

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
