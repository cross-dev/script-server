package processor_test

import (
	. "github.com/cross-dev/script-server/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/cross-dev/script-server/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestProcessor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Processor Suite")
}
