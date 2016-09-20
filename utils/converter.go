package utils

import (
	"regexp"
	"errors"
	"strings"
)

type source struct {
	Resource string
	Hub      string
	Repo     string
	Owner    string
	Vendor   string
	Brunch   string
}

// BToS - Convert bool type to string (yes|no)
func BToS(data bool) string {

	if data {
		return "yes"
	}
	return "no"

}

// https://github.com/lastbackend/vendors.git
// git@github.com:lastbackend/vendors.git

func GitUrlParse(url string) (*source, error) {

	var parsingUrl []string = regexp.MustCompile(`^(?:ssh|git|http(?:s)?)(?:@|:\/\/(?:.+@)?)((\w+)\.\w+)(?:\/|:)(.+)(?:\/)(.+)(?:\..+)$`).FindStringSubmatch(url)

	if len(parsingUrl) < 5 {
		return nil, errors.New("Can't parse url")
	}

	return &source{
		Resource: parsingUrl[0],
		Hub:   parsingUrl[1],
		Vendor:   parsingUrl[2],
		Owner:    parsingUrl[3],
		Repo:     parsingUrl[4],
	}, nil

}

func DockerNamespaceParse(namespace string) (*source, error) {

	var parsingNamespace *source = new(source)
	parsingNamespace.Vendor = "dockerhub"

	splitStr := strings.Split(namespace, "/")
	switch len(splitStr) {
	case 1:
		parsingNamespace.Repo = splitStr[0]
		return parsingNamespace, nil
	case 2:
		parsingNamespace.Owner = splitStr[0]
	case 3:
		parsingNamespace.Hub = splitStr[0]
		parsingNamespace.Owner = splitStr[1]
	default:
		return nil, errors.New("Can't parse url")
	}
	repoAndTag := strings.Split(splitStr[len(splitStr)-1], ":")
	parsingNamespace.Repo = repoAndTag[0]
	if len(repoAndTag) == 2 {
		parsingNamespace.Brunch = repoAndTag[1]
	}

	return parsingNamespace, nil

}