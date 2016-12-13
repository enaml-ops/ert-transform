package erttransform_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPushAppsManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PushAppsManager Suite")
}
