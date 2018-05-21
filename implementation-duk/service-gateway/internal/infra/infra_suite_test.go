package repository_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var T *testing.T

func TestInfra(t *testing.T) {
	T = t
	RegisterFailHandler(Fail)
	RunSpecs(t, "Infra Suite")
}
