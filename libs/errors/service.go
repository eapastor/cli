package errors

var storageServiceErrors = map[string]string{
	"SERVICE_NOT_FOUND":    "Service with this name not found",
	"BAD_SERVICE_NAME":     "Bad service name",
	"BAD_VCS_URL":          "Bad vcs url",
	"BAD_DOCKER_NAMESPACE": "Bad Docker namespace",
	"BAD_SERVICE_REGION":   "Bad service region",
	"BAD_SERVICE_MEMORY":   "Bad service memory",
}

type service struct{}

var Service service

func (service) BadServiceUUID() string {

	return storageServiceErrors["BAD_SERVICE_NAME"]

}

func (service) BadVCSURL() string {

	return storageServiceErrors["BAD_VCS_URL"]

}

func (service) BadDockerNamespace() string {

	return storageServiceErrors["BAD_DOCKER_NAMESPACE"]

}

func (service) BadRegion() string {

	return storageServiceErrors["BAD_SERVICE_REGION"]

}

func (service) BadMemory() string {

	return storageServiceErrors["BAD_SERVICE_MEMORY"]

}
