package processor_test

import (
	. "github.com/cross-dev/script-server/processor"

	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/url"
	"text/template"
)

var _ = Describe("Processor", func() {
	values := make(url.Values)
	values["a"] = []string{"b", "bb"}
	values["c"] = []string{"d"}
	values["ef"] = []string{"gh"}
	Context("happy path", func() {
		values := make(url.Values)
		values["a"] = []string{"b", "bb"}
		values["c"] = []string{"d"}
		values["ef"] = []string{"gh"}
		It("replaces arguments", func() {
			buffer := bytes.NewBuffer(make([]byte, 300))
			buffer.Reset()
			Process(values, "{{a}}-{{c}}_{{ef}}{}", buffer)
			Expect(buffer.String()).To(Equal("b-d_gh{}"))
		})
	})
})
