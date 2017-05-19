package testsapi

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	GitHubAPIEndpoint = "https://api.github.com"
)

func GithubAPIGinkgoSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GithubAPIGinkgo Suite")
}
